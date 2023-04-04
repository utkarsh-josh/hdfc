package svcutils

import (
	"net/url"
	"strings"
)

const (
	Name = "name"
)

var AllowedQueryParams = map[string]struct{} {
	Name: struct{}{},
}

// GetStringArrayQueryVariable retrieves string array for the query variable
func GetStringArrayQueryVariable(vars url.Values, variable string) ([]string, error) {
	valStr := vars.Get(variable)
	if valStr == "" {
		return nil, nil
	}

	arr := strings.Split(valStr, ",")
	var params []string
	for _, p := range arr {
		p = strings.TrimSpace(p)
		if p != "" {
			params = append(params, p)
		}
	}
	return params, nil
}
