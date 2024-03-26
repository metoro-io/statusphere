package urlgetter

import "time"

type URLGetter interface {
	// GetUrlsToScrape returns a list of URLs to scrape.
	// This can be called at any point so the URLGetter should be able to return the URLs quickly
	// And should only return URLs that should actually be scraped
	GetUrlsToScrape() ([]string, error)

	// GetHistoricalUrlsToScrape returns a list of URLs to scrape that are historical
	// This can be called at any point so the URLGetter should be able to return the URLs quickly
	// And should only return URLs that should actually be historical scraped
	GetHistoricalUrlsToScrape() ([]string, error)

	// UpdateLastScrapedTime updates the last scraped time for the given URL
	UpdateLastScrapedTime(url string, time time.Time) error

	// UpdateLastScrapedTimeHistorical updates the last scraped time for the given URL for historical scraping
	UpdateLastScrapedTimeHistorical(url string, time time.Time) error
}
