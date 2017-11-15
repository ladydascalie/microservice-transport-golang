package microservicetransport

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LUSHDigital/microservice-transport-golang/config"
	"github.com/LUSHDigital/microservice-transport-golang/domain"
)

// Service - Responsible for communication with a service.
type Service struct {
	Branch         string        // VCS branch the service is built from.
	CurrentRequest *http.Request // Current HTTP request being actioned.
	Environment    string        // CI environment the service operates in.
	Namespace      string        // Namespace of the service.
	Name           string        // Name of the service.
	Version        int           // Major API version of the service.
}

// Call - Do the current service request.
func (s *Service) Call() (*http.Response, error) {
	return HTTPClient.Do(s.CurrentRequest)
}

// Dial - Create a request to a service resource.
func (s *Service) Dial(request *Request) error {
	var err error

	// Make any alterations based upon the namespace.
	switch s.Namespace {
	case "aggregators":
		s.Name = strings.Join([]string{config.AggregatorDomainPrefix, s.Name}, "-")
	}

	// Determine the service namespace to use based on the service version.
	serviceNamespace := s.Name
	if s.Version != 0 {
		serviceNamespace = fmt.Sprintf("%s-%d", serviceNamespace, s.Version)
	}

	// Get the name of the service.
	dnsName := domain.BuildServiceDNSName(s.Name, s.Branch, s.Environment, serviceNamespace)

	// Build the resource URL.
	resourceUrl := fmt.Sprintf("%s://%s/%s", request.getProtocol(), dnsName, request.Resource)

	// Append the query string if we have any.
	if len(request.Query) > 0 {
		resourceUrl = fmt.Sprintf("%s?%s", resourceUrl, request.Query.Encode())
	}

	// Create the request.
	s.CurrentRequest, err = http.NewRequest(request.Method, resourceUrl, request.Body)
	return err
}

// Dial - Get the name of the service
func (s *Service) GetName() string {
	return s.Name
}