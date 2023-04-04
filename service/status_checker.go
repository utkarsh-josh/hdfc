package service

import (
	"net/http"
	"net/url"
	"time"
)

func (bl *BL) CheckWebsitesStatus() {
	m := make(map[string]string)

	resp := bl.dl.ListWebsitesStatus()
	for website := range resp.StatusMap {
		statusCode, err := bl.sendOverHttp(http.MethodGet, website)
		if err != nil || statusCode != http.StatusOK {
			m[website] = "DOWN"
			continue
		}
		m[website] = "UP"
	}

	bl.dl.UpdateWebsitesStatus(m)

	bl.logger.Log(
		"method", "CheckWebsitesStatus",
		"msg", "Period Website Status Check Completed",
		"time", time.Now(),
	)
}

func (bl *BL) sendOverHttp(method, website string) (int, error) {
	u, err := url.Parse(website)
	if err != nil {
		bl.logger.Log(
			"method", "sendOverHttp",
			"msg", "Invalid URL",
			"url", website,
			"err", err.Error(),
			"time", time.Now(),
		)
		return http.StatusBadGateway, err
	}

	url := url.URL{Scheme: "http", Host: u.Host, Path: u.Path}
	client := &http.Client{}
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		bl.logger.Log(
			"method", "sendOverHttp",
			"msg", "Failed to create http request",
			"url", url.String(),
			"err", err.Error(),
			"time", time.Now(),
		)
		return http.StatusBadGateway, err
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
		return http.StatusBadGateway, err
	}

	return resp.StatusCode, nil
}