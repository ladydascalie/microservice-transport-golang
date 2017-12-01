package microservicetransport

import (
	"io"
	"net/url"

	"github.com/LUSHDigital/microservice-transport-golang/config"
)

// Header - Models a http header.
type Header struct {
	Key   string
	Value string
}

// Request - Models a request to a service.
type Request struct {
	Body     io.ReadCloser // Body to pass in the request.
	Method   string        // HTTP method/verb for the request.
	Query    url.Values    // Query string values.
	Resource string        // Endpoint/resource on the requested service.
	Protocol string        // Transfer protocol to access the service with.
	Headers  []Header      // Headers to pass with the request.
}

// getProtocol - Get the transfer protocol to use for the service
func (r *Request) getProtocol() string {
	switch r.Protocol {
	case config.ProtocolHTTP, config.ProtocolHTTPS:
		return r.Protocol
	default:
		return config.ProtocolHTTP
	}
}
