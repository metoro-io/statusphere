package consumers

import "github.com/metoro-io/metoro/mrs-hudson/scraper/api"

type Consumer interface {
	// Consume consumes the given incidents
	Consume(incidents []api.Incident) error
}
