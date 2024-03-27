package statusio

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/metoro-io/statusphere/common/api"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *StatusioProvider) Name() string {
	return "statusio"
}

type StatusioProvider struct {
	logger     *zap.Logger
	httpClient *http.Client
}

func NewStatusioProvider(logger *zap.Logger, httpClient *http.Client) *StatusioProvider {
	return &StatusioProvider{
		logger:     logger,
		httpClient: httpClient,
	}
}

func (s *StatusioProvider) ScrapeStatusPageHistorical(ctx context.Context, url string) ([]api.Incident, error) {
	return s.scrapeStatusIoPageHistorical(ctx, url)
}

func (s *StatusioProvider) ScrapeStatusPageCurrent(ctx context.Context, url string) ([]api.Incident, error) {
	return s.scrapeStatusIoPageCurrent(ctx, url)
}

// scrapeStatusIoPageCurrent is a helper function that will attempt to scrape the status
// page using the status.io method
// If the status.io method fails, it will return an error
func (s *StatusioProvider) scrapeStatusIoPageCurrent(ctx context.Context, url string) ([]api.Incident, error) {
	// Get the current ongoing incidents
	incidentsOngoing, err := s.getOngoingIncidents(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the ongoing incidents")
	}

	// Get the most recent historical incidentsHistoricalRecent
	incidentsHistoricalRecent, err := s.getHistoricalPageOfIncidents(url, 1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the most recent historical incidentsHistoricalRecent")
	}

	// De-dupe the incidents
	incidents := append(incidentsOngoing, incidentsHistoricalRecent...)
	incidentsDeDuped := make(map[string]api.Incident)
	for _, incident := range incidents {
		incidentsDeDuped[incident.DeepLink] = incident
	}

	incidents = []api.Incident{}
	for _, incident := range incidentsDeDuped {
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

// scrapeStatusIoPageHistorical is a helper function that will attempt to scrape the status page using the status.io method
// If the status.io method fails, it will return an error
func (s *StatusioProvider) scrapeStatusIoPageHistorical(ctx context.Context, url string) ([]api.Incident, error) {
	var incidents []api.Incident

	// Get the last 40 quarters of incidents == 10 years
	i := 40
	for page := 1; page <= i; page++ {
		// Get the html of the status page
		incidentPage, err := s.getHistoricalPageOfIncidents(url, page)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get the historical incidents")
		}
		incidents = append(incidents, incidentPage...)
	}
	return incidents, nil
}

func (s *StatusioProvider) scrapeStatusIoHistoryPage(url string, page int) (string, error) {
	// First we get the status page history
	historyUrl := url + "/history?page=" + strconv.Itoa(page)
	history, err := s.httpClient.Get(historyUrl)
	if err != nil {
		return "", errors.Wrap(err, "failed to make the get request to the history page")
	}

	// Pull out the HTML from the response
	historyHtml, err := io.ReadAll(history.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read the history page response body")
	}

	return string(historyHtml), nil
}

func (s *StatusioProvider) getHistoricalPageOfIncidents(url string, page int) ([]api.Incident, error) {
	historyPageHtml, err := s.scrapeStatusIoHistoryPage(url, page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scrape the status page history")
	}

	// Parse the incidents from the history page
	incidentPage, err := s.parseIncidents(url, historyPageHtml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse the incidents from the history page")
	}
	return incidentPage, nil
}

func (s *StatusioProvider) parseIncidents(url string, html string) ([]api.Incident, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse the history page html")
	}

	// Placeholder for incidents
	var incidents []api.Incident

	// Find the script tag with the JSON data
	doc.Find("div[data-react-class='HistoryIndex']").Each(func(i int, selection *goquery.Selection) {
		dataReactProps, exists := selection.Attr("data-react-props")
		if !exists {
			s.logger.Error("Could not find data-react-props attribute")
			return
		}

		var pageStatus PageStatus

		// Unmarshal the JSON into the PageStatus struct
		err := json.Unmarshal([]byte(dataReactProps), &pageStatus)
		if err != nil {
			s.logger.Error("Error unmarshalling JSON", zap.Error(err))
			return
		}

		// Transform data into Incident and IncidentEvent structs (example)
		for _, month := range pageStatus.Months {
			for _, inc := range month.Incidents {
				link := url + "/incidents/" + inc.Code
				// Parse the timestamp, add logic to handle parsing
				startTime, endTime, err := parseDateString(strconv.Itoa(month.Year), month.Month, inc.Timestamp)
				if err != nil {
					s.logger.Error("Error parsing time", zap.Error(err), zap.String("timestamp", inc.Timestamp), zap.String("deep_link", link))
					return
				}

				// Example transformation, customize as needed
				incident := api.Incident{
					Title:         inc.Name,
					Description:   &inc.Message,
					StartTime:     startTime,
					EndTime:       endTime,
					Impact:        api.Impact(inc.Impact),
					DeepLink:      link,
					StatusPageUrl: url,
				}
				incidents = append(incidents, incident)
			}
		}
	})

	return incidents, nil
}

func (s *StatusioProvider) getOngoingIncidents(url string) ([]api.Incident, error) {
	pageHtml, err := s.getOngoingIncidentsPageHtml(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the ongoing incidents page html")
	}

	// Parse the incidents from the ongoing incidents page
	incidents, err := s.parseCurrentIncidents(url, pageHtml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse the incidents from the ongoing incidents page")
	}
	return incidents, nil
}

func (s *StatusioProvider) getOngoingIncidentsPageHtml(url string) (string, error) {
	history, err := s.httpClient.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "failed to make the get request to the history page")
	}

	// Pull out the HTML from the response
	historyHtml, err := io.ReadAll(history.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read the history page response body")
	}

	return string(historyHtml), nil
}

func (s *StatusioProvider) parseCurrentIncidents(url string, html string) ([]api.Incident, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse the history page html")
	}

	var incidents []api.Incident

	// Find and iterate over each unresolved incident
	doc.Find(".unresolved-incident").Each(func(i int, selection *goquery.Selection) {
		var incident api.Incident

		// Extract the incident's title
		incident.Title = selection.Find(".actual-title").Text()
		deepLink := selection.Find(".incident-title a").First().AttrOr("href", "")
		deepLink = url + deepLink
		incident.DeepLink = deepLink
		incident.StatusPageUrl = url
		var minTime *time.Time = nil

		selection.Find(".update").Each(func(i int, sel *goquery.Selection) {
			event := api.IncidentEvent{}
			// Extract the update's timestamp
			timestamp := sel.Find("small").Find("span").First().AttrOr("data-datetime-unix", "")
			timeInt, err := strconv.ParseInt(timestamp, 10, 64)
			if err != nil {
				s.logger.Error("Error parsing timestamp", zap.Error(err))
				return
			}
			event.Time = time.UnixMilli(timeInt)
			if minTime == nil || event.Time.Before(*minTime) {
				minTime = &event.Time
			}

			// Title
			event.Title = sel.Find("strong").Text()

			// Extract the update's message
			event.Description = sel.Find("span").First().Text()

			incident.Events = append(incident.Events, event)
		})
		incident.StartTime = *minTime
		// Append the extracted incident to the slice
		incidents = append(incidents, incident)
	})
	return incidents, nil
}

// We need to parse strings in this format
// "Mar <var data-var='date'>13</var>, <var data-var='time'>06:55</var> - <var data-var='time'>16:02</var> UTC"
// "Feb <var data-var='date'>25</var>, <var data-var='time'>23:44</var> - Feb <var data-var='date'>26</var>, <var data-var='time'>20:27</var> UTC"
func parseDateString(year string, month string, dateString string) (time.Time, *time.Time, error) {
	// Define layout patterns
	const layoutSingle = "2006 January 2, 15:04 MST"
	const layoutStart = "2006 January 2, 15:04 MST"
	const layoutEnd = "2006 January 2, 15:04 MST"

	// Precompile regex for extracting date and time components, ignoring HTML tags
	re := regexp.MustCompile(`<var data-var='[^']+'>([^<]+)</var>`)
	matches := re.FindAllStringSubmatch(dateString, -1)

	if len(matches) < 2 {
		return time.Time{}, nil, errors.New("unable to parse date string: insufficient data found")
	}
	// Extract date and time components
	date := matches[0][1]      // date is always the first match
	startTime := matches[1][1] // start time is the second match
	endTimeComponent := ""     // End time might not be available

	if len(matches) > 2 {
		endTimeComponent = matches[2][1]
	}

	// Construct date strings for parsing
	startDateString := fmt.Sprintf("%s %s %s, %s UTC", year, month, date, startTime)
	endDateString := fmt.Sprintf("%s %s %s, %s UTC", year, month, date, endTimeComponent)

	if endTimeComponent == "" { // Handle single day format
		startTime, err := time.Parse(layoutSingle, startDateString)
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("error parsing start time: %w", err)
		}
		return startTime, nil, nil
	} else { // Handle multi-day or same day with end time format
		startTime, err := time.Parse(layoutStart, startDateString)
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("error parsing start time: %w", err)
		}
		// Check if date changes for end time
		if len(matches) == 4 { // If there's a separate end date
			endDate := matches[2][1]
			endDateString = fmt.Sprintf("%s %s %s, %s UTC", year, month, endDate, matches[3][1])
		}

		endTime, err := time.Parse(layoutEnd, endDateString)
		if err != nil {
			return time.Time{}, nil, fmt.Errorf("error parsing end time: %w", err)
		}

		return startTime, &endTime, nil
	}
}

// Additional structs to capture the overall structure of the JSON
type PageStatus struct {
	Components []Component `json:"components"`
	Months     []Month     `json:"months"`
}

type Component struct {
	Name string `json:"name"`
}

type Month struct {
	Incidents []IncidentRaw `json:"incidents"`
	Year      int           `json:"year"`
	Month     string        `json:"name"`
}

type IncidentRaw struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Code      string `json:"code"`
	Impact    string `json:"impact"`
}
