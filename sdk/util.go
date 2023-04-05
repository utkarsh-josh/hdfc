package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (s *serviceSDK) sendHTTPRequest(method, url, reqData string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(reqData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid request or Internal server error")
	}

	return body, nil
}
