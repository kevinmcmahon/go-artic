package model

// SearchResponse is the struct representation of an Search Artwork api response
type SearchResponse struct {
	Preference interface{} `json:"preference"`
	Pagination struct {
		Total       int `json:"total"`
		Limit       int `json:"limit"`
		Offset      int `json:"offset"`
		TotalPages  int `json:"total_pages"`
		CurrentPage int `json:"current_page"`
	} `json:"pagination"`
	Data []struct {
		Score     float64 `json:"_score"`
		Thumbnail struct {
			AltText string `json:"alt_text"`
			Width   int    `json:"width"`
			Type    string `json:"type"`
			URL     string `json:"url"`
			Lqip    string `json:"lqip"`
			Height  int    `json:"height"`
		} `json:"thumbnail"`
		APIModel  string `json:"api_model"`
		IsBoosted bool   `json:"is_boosted"`
		APILink   string `json:"api_link"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Timestamp string `json:"timestamp"`
	} `json:"data"`
	Info struct {
		LicenseText  string   `json:"license_text"`
		LicenseLinks []string `json:"license_links"`
		Version      string   `json:"version"`
	} `json:"info"`
}
