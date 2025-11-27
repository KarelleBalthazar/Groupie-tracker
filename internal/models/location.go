package models

type LocationsResponse struct {
	Index []LocationItem `json:"index"`
}

type LocationItem struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
