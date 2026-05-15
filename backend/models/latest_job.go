package models
type LatestJob struct {
	Title string `json:"title"`
	Company string `json:"company"`
	IsNew bool `json:"isNew"`
}