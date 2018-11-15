package remotehttp

import (
	"time"
)

// WebConfig contains basic http session configuration parameters.
type WebConfig struct {
	LogLevel            int           // verbosity of output
	ConTimeout          time.Duration // Initial Connection timeout.
	SSLHandshakeTimeout time.Duration // SSL handshake timeout.
	RxTimeout           time.Duration // Receiving transfer timeout.
}

// Headers contains a number of headers for the http session.
type Headers []struct {
	Label string // Header label
	Value string // Header value
}
