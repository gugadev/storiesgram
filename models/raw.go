package models

// Image represent an image
type Image struct {
	URL string `json:"url"`
}

// Candidate represent a bunch of images
type Candidate struct {
	Candidates []Image `json:"candidates"`
}

// Item represent an items child
type Item struct {
	Images    Candidate `json:"image_versions2"`
	MediaType int       `json:"media_type"`
	PK        int       `json:"pk"`
}

// Raw represent the entire json response
type Raw struct {
	Items []Item `json:"items"`
}
