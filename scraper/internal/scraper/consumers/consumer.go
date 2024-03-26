package consumers

import (
	"github.com/metoro-io/statusphere/common/api"
)

type Consumer interface {
	// Consume consumes the given incidents
	Consume(incidents []api.Incident) error
}
