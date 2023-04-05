package service

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/utkarsh-josh/hdfc/spec"
	"github.com/utkarsh-josh/hdfc/svcconst"
)

// CheckWebsitesStatus checks status of all the websites in the memory map
func (bl *BL) CheckWebsitesStatus(statusMap map[string]string, ch chan spec.WebsiteStatus) {
	var wg sync.WaitGroup
	for website := range statusMap {
		wg.Add(1)
		go bl.sendOverHTTP(http.MethodGet, website, ch, &wg)
	}
	wg.Wait()
	bl.logger.Log(
		"method", "CheckWebsitesStatus",
		"msg", "Period Website Status Check Completed",
		"time", time.Now(),
	)
}

// sendOverHttp send http request to check the status of the website
func (bl *BL) sendOverHTTP(method, website string, ch chan spec.WebsiteStatus, wg *sync.WaitGroup) {
	defer wg.Done()

	u, err := url.Parse(website)
	if err != nil {
		bl.logger.Log(
			"method", "sendOverHttp",
			"msg", "Invalid URL",
			"url", website,
			"err", err.Error(),
			"time", time.Now(),
		)
		ch <- spec.WebsiteStatus{Name: website, Status: svcconst.StatusDown}
		return
	}

	url := url.URL{Scheme: "http", Host: u.Host, Path: u.Path}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		bl.logger.Log(
			"method", "sendOverHttp",
			"msg", "Failed to create http request",
			"url", url.String(),
			"err", err.Error(),
			"time", time.Now(),
		)
		ch <- spec.WebsiteStatus{Name: website, Status: svcconst.StatusDown}
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		bl.logger.Log(
			"method", "sendOverHttp",
			"msg", "Failed to send http request",
			"url", url.String(),
			"err", err.Error(),
			"time", time.Now(),
		)
		ch <- spec.WebsiteStatus{Name: website, Status: svcconst.StatusDown}
		return
	}

	if resp.StatusCode != http.StatusOK {
		ch <- spec.WebsiteStatus{Name: website, Status: svcconst.StatusDown}
		return
	}
	ch <- spec.WebsiteStatus{Name: website, Status: svcconst.StatusUp}
}
