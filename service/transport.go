package service

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/utkarsh-josh/hdfc/spec"
	"github.com/utkarsh-josh/hdfc/svcutils"
)

func DecodeAddWebsitesRequest(r *http.Request) (interface{}, error) {
	request := spec.WebsitesRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}


func DecodeGetWebsitesStatus(r *http.Request) (interface{}, error) {
	request := spec.WebsitesRequest{}
	var err error

	r.ParseForm()
	for k := range r.Form {
		if _, ok := svcutils.AllowedQueryParams[k]; !ok {
			return nil, fmt.Errorf("Invalid QueryParams : %s", k)
		}
	}
	// Query Params
	q := r.URL.Query()

	if q.Get(svcutils.Name) != "" {
		request.Websites, err = svcutils.GetStringArrayQueryVariable(q, svcutils.Name)
		if err != nil {
			return nil, err
		}
	}

	return &request, nil
}