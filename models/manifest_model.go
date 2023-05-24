package models

type ManifestData struct {
	UUID   string  `form:"uuid"`
	Label  string  `form:"label"`
	Images []Image `form:"images,omitempty"`
}

type Image struct {
	ID string `json:"id"`
}
