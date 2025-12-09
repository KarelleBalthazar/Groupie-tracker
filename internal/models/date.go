package models

type DatesResponse struct {
	Index []DateItem `json:"index"`
}

type DateItem struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
