package sdk

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/log"
	"github.com/utkarsh-josh/hdfc/spec"
)

// Service is the interface to call API
type Service interface {
	AddWebsites(req *spec.WebsitesRequest) error
	ListWebsitesStatus() (*spec.ListWebsitesResponse, error)
	GetWebsitesStatus(queryParams map[string]string) (*spec.ListWebsitesResponse, error)
}

type serviceSDK struct {
	logger log.Logger
	host   string
}

// NewServiceSDK returns sdk layer handler
func NewServiceSDK(logger log.Logger, host string) Service {
	return &serviceSDK{
		logger: logger,
		host:   host,
	}
}

func (s *serviceSDK) AddWebsites(req *spec.WebsitesRequest) error {
	jsonData, _ := json.Marshal(req)

	u := url.URL{Scheme: "http", Host: s.host, Path: spec.WebsiteURL}
	_, err := s.sendHTTPRequest(http.MethodPost, u.String(), string(jsonData))
	return err
}

func (s *serviceSDK) ListWebsitesStatus() (*spec.ListWebsitesResponse, error) {
	u := url.URL{Scheme: "http", Host: s.host, Path: spec.WebsiteURL}
	body, err := s.sendHTTPRequest(http.MethodGet, u.String(), "")
	if err != nil {
		return nil, err
	}

	resp := &spec.ListWebsitesResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *serviceSDK) GetWebsitesStatus(queryParams map[string]string) (*spec.ListWebsitesResponse, error) {
	u := url.URL{Scheme: "http", Host: s.host, Path: spec.WebsiteURL}
	query := u.Query()
	for queryKey, queryVal := range queryParams {
		query.Set(queryKey, queryVal)
	}
	u.RawQuery = query.Encode()

	body, err := s.sendHTTPRequest(http.MethodGet, u.String(), "")
	if err != nil {
		return nil, err
	}

	resp := &spec.ListWebsitesResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
