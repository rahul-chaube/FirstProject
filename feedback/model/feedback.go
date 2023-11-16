package model

type Feedback struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FileData struct {
	Name         string
	Size         string
	ModifiedTime string
}
