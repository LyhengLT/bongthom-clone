package models
type QuickLinks struct {
	Label string `json:"label"`
	Url string `json:"url"`
	Extra string `json:"extra,omitempty"`
	Img string `json:"img,omitempty"`
}