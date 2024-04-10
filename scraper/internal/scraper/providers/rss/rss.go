package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/metoro-io/statusphere/common/api"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func (s *RssProvider) Name() string {
	return "RSS"
}

type RssProvider struct {
	logger     *zap.Logger
	httpClient *http.Client
}

func NewRssProvider(logger *zap.Logger, httpClient *http.Client) *RssProvider {
	return &RssProvider{
		logger:     logger,
		httpClient: httpClient,
	}
}

// There is no historical page differentiation for rss pages so we skip
func (s *RssProvider) ScrapeStatusPageHistorical(ctx context.Context, url string) ([]api.Incident, error) {
	_, isRssPage, err := s.isRssPage(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to determine if the page is an rss page")
	}
	if !isRssPage {
		return nil, errors.New("page is not a rss page")
	}
	return []api.Incident{}, nil
}

func (s *RssProvider) ScrapeStatusPageCurrent(ctx context.Context, url string) ([]api.Incident, error) {
	return s.scrapeRssPage(ctx, url)
}

// scrapeRssPage is a helper function that will attempt to scrape the status
// page using the rss method
// If the ress method fails, it will return an error
func (s *RssProvider) scrapeRssPage(ctx context.Context, url string) ([]api.Incident, error) {
	rssPage, isRssPage, err := s.isRssPage(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to determine if the page is an rss page")
	}
	if !isRssPage {
		return nil, errors.New("page is not a rss page")
	}

	// Get the incidents from the rss page
	return s.getIncidentsFromRssPage(rssPage, url)
}

// We determine if a page is an rss page by checking if there is a /history page and
// that history page contains the data-react-class='HistoryIndex' attribute
func (s *RssProvider) isRssPage(url string) (string, bool, error) {
	// Get the history page
	historyUrls := getUrls(url)
	for _, historyUrl := range historyUrls {
		history, err := s.httpClient.Get(historyUrl)
		if err != nil {
			return "", false, errors.Wrap(err, "failed to make the get request to the history page")
		}
		if history.StatusCode != http.StatusOK {
			continue
		}

		// Is the body well formed xml?
		if isXMLContent(history.Body) {
			return historyUrl, true, nil
		}

	}
	return "", false, nil
}

func (s *RssProvider) getIncidentsFromRssPage(url string, statusPageUrl string) ([]api.Incident, error) {
	var incidents []api.Incident

	// Fetch the RSS or Atom feed
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the feed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the feed: %w", err)
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the feed: %w", err)
	}

	for _, item := range feed.Items {
		var parsedTime time.Time

		// Check and parse the item's published or updated time
		if item.PublishedParsed != nil {
			parsedTime = *item.PublishedParsed
		}

		// Strip HTML tags from title and description - assuming a function stripHTML exists
		title := stripHTML(item.Title)
		description := stripHTML(item.Content)
		if description == "" {
			description = stripHTML(item.Description)
		}
		deepLink := item.Link

		incidents = append(incidents, api.Incident{
			Title:       title,
			Description: &description,
			StartTime:   parsedTime,
			EndTime:     &parsedTime,
			DeepLink:    deepLink,
			// Not all RSS feeds have an impact field, so we default to none
			Impact:        api.ImpactNone,
			StatusPageUrl: statusPageUrl,
		})
	}

	if len(incidents) == 0 {
		return nil, errors.New("no incidents found")
	}

	return incidents, nil
}

func getUrls(url string) []string {
	return []string{
		url + "/history.atom",
		url + "/feed.atom",
		url + "/de.atom",
		url + "/_rss",
		url + "/rss/all.rss",
		url + "/en-us/status/feed/",
	}
}

// isXMLContent tries to parse the response body as XML to check if it is valid XML.
func isXMLContent(body io.Reader) bool {
	decoder := xml.NewDecoder(body)
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		directive, ok := token.(xml.Directive)
		if ok {
			if strings.Contains(strings.ToLower(string(directive)), "html") {
				break
			}
		}

		if token == nil {
			break
		}
		// If we can decode one token, it's likely an XML.
		// This is a simplistic check and might need to be more sophisticated
		// depending on the context.

		return true
	}
	return false
}

// stripHTML uses a regular expression to remove HTML tags from a string.
func stripHTML(input string) string {
	// Compile the regular expression to match HTML tags.
	// The expression "<[^>]*>" matches anything that starts with "<" and ends with ">",
	// containing any characters except ">" in between.
	re, err := regexp.Compile("<[^>]*>")
	if err != nil {
		panic("Invalid regular expression")
	}

	// Replace all HTML tags with an empty string.
	return re.ReplaceAllString(input, "")
}
