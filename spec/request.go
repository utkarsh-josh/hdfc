package spec

// WebsitesRequest is the request schema
type WebsitesRequest struct {
	Websites []string `json:"websites"`
}

// WebsiteStatus is the struct to update website status
type WebsiteStatus struct {
	Name   string
	Status string
}
