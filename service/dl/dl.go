package dl

import (
	"fmt"

	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/spec"
	"github.com/utkarsh-josh/hdfc/svcconst"
)

// DL is the database logic layer struct
type DL struct {
	logger     log.Logger
	websiteMap map[string]string
}

// NewDL returns dl layer handler
func NewDL(logger log.Logger) *DL {
	dl := &DL{
		logger:     logger,
		websiteMap: make(map[string]string),
	}
	return dl
}

// AddWebsitesToStatusMap adds websites to the memory map for status check
func (dl *DL) AddWebsitesToStatusMap(req *spec.WebsitesRequest) (spec.AddWebsiteResponse, error) {
	if len(req.Websites) == 0 {
		return false, fmt.Errorf("Empty website list provided")
	}

	m := make(map[string]string)
	for _, website := range req.Websites {
		if status, ok := dl.websiteMap[website]; ok {
			m[website] = status
			continue
		}
		m[website] = svcconst.StatusNotYetChecked
	}

	dl.websiteMap = m
	return spec.AddWebsiteResponse(true), nil
}

// ListWebsitesStatus list down status of all the websites status from the memory map
func (dl *DL) ListWebsitesStatus() *spec.ListWebsitesResponse {
	resp := &spec.ListWebsitesResponse{StatusMap: dl.websiteMap}
	return resp
}

// GetWebsitesStatusFromStatusMap fetches status of given websites from the memory map
func (dl *DL) GetWebsitesStatusFromStatusMap(req *spec.WebsitesRequest) *spec.ListWebsitesResponse {
	m := make(map[string]string)
	for _, website := range req.Websites {
		if val, ok := dl.websiteMap[website]; ok {
			m[website] = val
		}
	}

	resp := &spec.ListWebsitesResponse{StatusMap: m}
	return resp
}

// UpdateWebsitesStatus updates the status of existing websites in the memory map
func (dl *DL) UpdateWebsitesStatus(website, status string) {
	if _, ok := dl.websiteMap[website]; ok {
		dl.websiteMap[website] = status
	}
}
