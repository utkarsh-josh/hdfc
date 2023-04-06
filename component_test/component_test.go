package component_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/utkarsh-josh/hdfc/spec"
	"github.com/utkarsh-josh/hdfc/svcconst"
)

func Test_AddWebsites(t *testing.T) {
	tests := []struct {
		name    string
		req     *spec.WebsitesRequest
		wantErr bool
	}{
		{
			name: "Success",
			req: &spec.WebsitesRequest{
				Websites: []string{"www.google.com", "www.facebook.com"},
			},
			wantErr: false,
		},
		{
			name: "Success: Storing status of existing websites",
			req: &spec.WebsitesRequest{
				Websites: []string{
					"www.google.com",
					"www.facebook.com",
					"www.twitter.com",
					"www.google.com/utkarsh",
					"www.utkarsh812.com",
					"@$%^",
				},
			},
			wantErr: false,
		},
		{
			name:    "Error: Empty list of websites provided",
			req:     &spec.WebsitesRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testObj.API.AddWebsites(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWebsites error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ListWebsitesStatus(t *testing.T) {
	tests := []struct {
		name             string
		wantErr          bool
		expectedResponse *spec.ListWebsitesResponse
		sleepDurationSec int
	}{
		{
			name:    "Success: Status not yet checked",
			wantErr: false,
			expectedResponse: &spec.ListWebsitesResponse{
				StatusMap: map[string]string{
					"www.google.com":         svcconst.StatusNotYetChecked,
					"www.facebook.com":       svcconst.StatusNotYetChecked,
					"www.twitter.com":        svcconst.StatusNotYetChecked,
					"www.google.com/utkarsh": svcconst.StatusNotYetChecked,
					"www.utkarsh812.com":     svcconst.StatusNotYetChecked,
					"@$%^":                   svcconst.StatusNotYetChecked,
				},
			},
			sleepDurationSec: 0,
		},
		{
			name:    "Success",
			wantErr: false,
			expectedResponse: &spec.ListWebsitesResponse{
				StatusMap: map[string]string{
					"www.google.com":         svcconst.StatusUp,
					"www.facebook.com":       svcconst.StatusUp,
					"www.twitter.com":        svcconst.StatusUp,
					"www.google.com/utkarsh": svcconst.StatusDown,
					"www.utkarsh812.com":     svcconst.StatusDown,
					"@$%^":                   svcconst.StatusDown,
				},
			},
			sleepDurationSec: svcconst.TestEnvWaitTimeForStatusCheck,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sleepDurationSec > 0 {
				testObj.Logger.Log(
					"method", "Test_ListWebsitesStatus",
					"msg", "Waiting for status checker to check websites status",
					"time", time.Now(),
				)
			}
			time.Sleep(time.Duration(tt.sleepDurationSec) * time.Second)
			resp, err := testObj.API.ListWebsitesStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListWebsitesStatus error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(resp, tt.expectedResponse) {
				t.Errorf("ListWebsitesStatus - \ngot info %+v \nwant info %+v \ndiff: %s", resp, tt.expectedResponse, cmp.Diff(resp, tt.expectedResponse))
				return
			}
		})
	}
}

func Test_GetWebsitesStatus(t *testing.T) {
	tests := []struct {
		name             string
		wantErr          bool
		queryParams      map[string]string
		expectedResponse *spec.ListWebsitesResponse
	}{
		{
			name:    "Success",
			wantErr: false,
			expectedResponse: &spec.ListWebsitesResponse{
				StatusMap: map[string]string{
					"www.google.com":   svcconst.StatusUp,
					"www.facebook.com": svcconst.StatusUp,
				},
			},
			queryParams: map[string]string{
				"name": "www.google.com, www.facebook.com",
			},
		},
		{
			name:    "Error: Invalid query params",
			wantErr: true,
			queryParams: map[string]string{
				"id": "www.google.com, www.facebook.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testObj.API.GetWebsitesStatus(tt.queryParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWebsitesStatus error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(resp, tt.expectedResponse) {
				t.Errorf("GetWebsitesStatus - \ngot info %+v \nwant info %+v \ndiff: %s", resp, tt.expectedResponse, cmp.Diff(resp, tt.expectedResponse))
				return
			}
		})
	}
}
