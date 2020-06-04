package model

import (
	"encoding/json"
	"fmt"
)

// Artwork is the struct representation of an Artwork response
type Artwork struct {
	ID          int
	URL         string
	Title       string
	Artist      string
	Date        string
	Description string
}

// PrettyString creates a pretty string of the Artwork
func (a Artwork) PrettyString() string {
	p := fmt.Sprintf(
		"Image Id: %d\nTitle: %s\nArtist: %s\nDate: %s\nDescription: %s\n",
		a.ID, a.Title, a.Artist, a.Date, a.Description)
	return p
}

// JSON returns the JSON representation of the artwork
func (a Artwork) JSON() string {
	cJSON, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(cJSON)
}
