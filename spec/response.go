package spec

type AddWebsiteResponse bool

/*type GetWebsiteResponse struct{
	Name string `json:"name"`
	Status string `json:"status"`
}*/

type ListWebsitesResponse struct {
	StatusMap map[string]string `json:"statusMap"`
}