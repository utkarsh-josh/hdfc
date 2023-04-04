package spec

// AddWebsiteResponse is the response schema for Add Websites
type AddWebsiteResponse bool

// ListWebsitesResponse is the response schema for List Websites Status
type ListWebsitesResponse struct {
	StatusMap map[string]string `json:"statusMap"`
}
