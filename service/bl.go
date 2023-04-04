package service

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/spec"
)

// DL is the interface to call dl layer functions
type DL interface {
	AddWebsitesToStatusMap(req *spec.WebsitesRequest) (spec.AddWebsiteResponse, error)
	ListWebsitesStatus() *spec.ListWebsitesResponse
	GetWebsitesStatusFromStatusMap(req *spec.WebsitesRequest) *spec.ListWebsitesResponse
	UpdateWebsitesStatus(website, status string)
}

// BL is the business logic layer struct
type BL struct {
	logger   log.Logger
	dl       DL
	duration int
}

// NewService returns bl layer handler
func NewService(logger log.Logger, dl DL, duration int) *BL {
	bl := &BL{
		logger:   logger,
		dl:       dl,
		duration: duration,
	}

	return bl
}

// AddWebsites adds websites to the memory map for status check
func (bl *BL) AddWebsites(req *spec.WebsitesRequest) (spec.AddWebsiteResponse, error) {
	return bl.dl.AddWebsitesToStatusMap(req)
}

// GetWebsitesStatus fetches status of the websites status from the memory map
func (bl *BL) GetWebsitesStatus(req *spec.WebsitesRequest) (*spec.ListWebsitesResponse, error) {
	if len(req.Websites) == 0 {
		return bl.dl.ListWebsitesStatus(), nil
	}

	return bl.dl.GetWebsitesStatusFromStatusMap(req), nil
}

// StatusChecker is the event loop for handling periodic status check
func (bl *BL) StatusChecker() {
	ticker := time.NewTicker(time.Duration(bl.duration) * time.Second)
	var statusChan chan spec.WebsiteStatus
	for {
		select {
		case <-ticker.C:
			bl.logger.Log(
				"method", "StatusChecker",
				"msg", "Period Website Status Check Triggered",
				"time", time.Now(),
			)
			resp := bl.dl.ListWebsitesStatus()
			statusChan = make(chan spec.WebsiteStatus, len(resp.StatusMap))
			bl.CheckWebsitesStatus(resp.StatusMap, statusChan)

		case status := <-statusChan:
			bl.dl.UpdateWebsitesStatus(status.Name, status.Status)
		}
	}
}
