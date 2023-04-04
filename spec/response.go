package spec

type AddWebsiteResponse bool

type ListWebsitesResponse struct {
	StatusMap map[string]string `json:"statusMap"`
}
