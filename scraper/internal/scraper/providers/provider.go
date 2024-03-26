package providers

import (
	"context"
	"github.com/metoro-io/metoro/mrs-hudson/scraper/api"
)

type Provider interface {
	// ScrapeStatusPageHistorical scrapes the status page at the given URL and returns a list of incidents
	// The incidents are historical, meaning they are not just the current incidents, this can be expected to return a large number of incidents
	// And take a long time to run, so we should only run this infrequently, maybe once per week per page
	ScrapeStatusPageHistorical(ctx context.Context, url string) ([]api.Incident, error)

	// ScrapeStatusPageCurrent scrapes the status page at the given URL and returns a list of incidents
	// The incidents are current, meaning they are only the recent incidents, this can be expected to return a small number of incidents
	// And take a short time to run, so we should run this frequently, maybe once per 5 minutes per page
	ScrapeStatusPageCurrent(ctx context.Context, url string) ([]api.Incident, error)
	Name() string
}
