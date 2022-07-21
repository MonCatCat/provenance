package streaming

import (
	"github.com/provenance-io/provenance/app/streaming/trace"
	"github.com/provenance-io/provenance/internal/streaming"
)

// StreamServiceInitializers contains a map of supported StreamServiceInitializer implementations
var StreamServiceInitializers = map[string]streaming.StreamServiceInitializer{
	"trace": trace.StreamServiceInitializer,
}