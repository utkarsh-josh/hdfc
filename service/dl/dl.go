package dl

import (
	"fmt"

	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/spec"
)

type DL struct {
	logger log.Logger
	websiteMap map[string]string
}

func NewDL(logger log.Logger) *DL {
	dl := &DL{
		logger: logger,
		websiteMap: make(map[string]string),
	}
	return dl
}

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
		m[website] = "StatusNotYetChecked"
	}

	dl.websiteMap = m
	return spec.AddWebsiteResponse(true), nil
}

func (dl *DL) ListWebsitesStatus() *spec.ListWebsitesResponse {
	resp := &spec.ListWebsitesResponse{StatusMap: dl.websiteMap}
	return resp
}

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

func (dl *DL) UpdateWebsitesStatus(statusMap map[string]string) {
	dl.websiteMap = statusMap
}