package service

import (
	"time"
	
	"github.com/go-kit/kit/log"

	"github.com/utkarsh-josh/hdfc/spec"
)

type DL interface {
	AddWebsitesToStatusMap(req *spec.WebsitesRequest) (spec.AddWebsiteResponse, error)
	ListWebsitesStatus() *spec.ListWebsitesResponse
	GetWebsitesStatusFromStatusMap(req *spec.WebsitesRequest) *spec.ListWebsitesResponse
	UpdateWebsitesStatus(statusMap map[string]string)
}

type BL struct{
	logger log.Logger
	dl DL
	duration int
}

func NewService(logger log.Logger, dl DL, duration int) *BL {
	bl := &BL{
		logger: logger,
		dl: dl,
		duration: duration,
	}

	return bl
}

func (bl *BL) AddWebsites(req *spec.WebsitesRequest) (spec.AddWebsiteResponse, error) {
	return bl.dl.AddWebsitesToStatusMap(req)
}

func (bl *BL) GetWebsitesStatus(req *spec.WebsitesRequest) (*spec.ListWebsitesResponse, error) {
	if len(req.Websites) == 0 {
		return bl.dl.ListWebsitesStatus(), nil
	}

	return bl.dl.GetWebsitesStatusFromStatusMap(req), nil
}

func (bl *BL) StatusChecker() {
	ticker := time.NewTicker(time.Duration(bl.duration) * time.Second)

	for {
		select {
		case <-ticker.C:
			bl.logger.Log(
				"method", "StatusChecker",
				"msg", "Period Website Status Check Started",
				"time", time.Now(),
			)
			bl.CheckWebsitesStatus()
		}
	}
}