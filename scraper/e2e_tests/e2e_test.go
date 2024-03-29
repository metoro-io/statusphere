package e2e_tests

import (
	"context"
	"github.com/metoro-io/statusphere/scraper/internal/scraper"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/providers"
	"github.com/metoro-io/statusphere/scraper/internal/scraper/providers/atlassian"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

var statusPages = []string{
	"https://status.dropbox.com",
	"https://www.calendlystatus.com",
	"https://status.whatnot.com",
	"https://www.githubstatus.com",
	"https://status.multiversx.com",
	"https://status.1password.com",
	"https://status.edq.com",
	"https://status.snowflake.com",
	"https://status.redhat.com",
	"https://status.payscale.com",
}

func TestE2eDropboxHistorical(t *testing.T) {
	dev, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("Failed to create logger")
	}
	scraper := scraper.NewScraper(dev, http.DefaultClient, []providers.Provider{atlassian.NewAtlassianProvider(dev, http.DefaultClient)})
	incidents, err := scraper.ScrapeStatusPageHistorical(context.Background(), "https://status.dropbox.com")
	if err != nil {
		t.Errorf("Failed to scrape status page: %s", "https://status.dropbox.com")
	}
	dev.Info("Incidents from status page", zap.Any("statusPage", "https://status.dropbox.com"), zap.Any("numIncidents", len(incidents)))
}

func TestE2eCloudflareCurrent(t *testing.T) {
	dev, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("Failed to create logger")
	}
	scraper := scraper.NewScraper(dev, http.DefaultClient, []providers.Provider{atlassian.NewAtlassianProvider(dev, http.DefaultClient)})
	incident, err := scraper.ScrapeStatusPageCurrent(context.Background(), "https://www.cloudflarestatus.com")
	if err != nil {
		t.Errorf("Failed to scrape status page: %s", "https://www.cloudflarestatus.com")
	}
	dev.Info("Incidents from status page", zap.Any("statusPage", "https://www.cloudflarestatus.com"), zap.Any("numIncidents", len(incident)))
}

func TestE2eManyCurrent(t *testing.T) {
	dev, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("Failed to create logger")
	}
	scraper := scraper.NewScraper(dev, http.DefaultClient, []providers.Provider{atlassian.NewAtlassianProvider(dev, http.DefaultClient)})
	for _, statusPage := range statusPages {
		incidents, err := scraper.ScrapeStatusPageCurrent(context.Background(), statusPage)
		if err != nil {
			t.Errorf("Failed to scrape status page: %s", statusPage)
		}
		dev.Info("Incidents from status page", zap.Any("statusPage", statusPage), zap.Any("numIncidents", len(incidents)))
	}
}
