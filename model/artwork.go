package model

import (
	"encoding/json"
	"fmt"
)

// ArtworkResult is an individual artwork returned by the artworks api
type ArtworkResult struct {
	Data struct {
		ID        int         `json:"id"`
		APIModel  string      `json:"api_model"`
		APILink   string      `json:"api_link"`
		IsBoosted bool        `json:"is_boosted"`
		Title     string      `json:"title"`
		AltTitles interface{} `json:"alt_titles"`
		Thumbnail struct {
			URL     string `json:"url"`
			Type    string `json:"type"`
			Lqip    string `json:"lqip"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			AltText string `json:"alt_text"`
		} `json:"thumbnail"`
		MainReferenceNumber         string      `json:"main_reference_number"`
		Pageviews                   int         `json:"pageviews"`
		PageviewsRecent             int         `json:"pageviews_recent"`
		HasNotBeenViewedMuch        bool        `json:"has_not_been_viewed_much"`
		BoostRank                   int         `json:"boost_rank"`
		DateStart                   int         `json:"date_start"`
		DateEnd                     int         `json:"date_end"`
		DateDisplay                 string      `json:"date_display"`
		DateQualifierTitle          string      `json:"date_qualifier_title"`
		DateQualifierID             interface{} `json:"date_qualifier_id"`
		ArtistDisplay               string      `json:"artist_display"`
		PlaceOfOrigin               string      `json:"place_of_origin"`
		Description                 string      `json:"description"`
		Dimensions                  string      `json:"dimensions"`
		MediumDisplay               string      `json:"medium_display"`
		Inscriptions                string      `json:"inscriptions"`
		CreditLine                  string      `json:"credit_line"`
		PublicationHistory          string      `json:"publication_history"`
		ExhibitionHistory           string      `json:"exhibition_history"`
		ProvenanceText              string      `json:"provenance_text"`
		PublishingVerificationLevel string      `json:"publishing_verification_level"`
		InternalDepartmentID        int         `json:"internal_department_id"`
		CollectionStatus            string      `json:"collection_status"`
		FiscalYear                  int         `json:"fiscal_year"`
		FiscalYearDeaccession       interface{} `json:"fiscal_year_deaccession"`
		IsPublicDomain              bool        `json:"is_public_domain"`
		IsZoomable                  bool        `json:"is_zoomable"`
		MaxZoomWindowSize           int         `json:"max_zoom_window_size"`
		CopyrightNotice             interface{} `json:"copyright_notice"`
		HasMultimediaResources      bool        `json:"has_multimedia_resources"`
		HasEducationalResources     bool        `json:"has_educational_resources"`
		Colorfulness                float64     `json:"colorfulness"`
		Color                       struct {
			H          int     `json:"h"`
			L          int     `json:"l"`
			S          int     `json:"s"`
			Percentage float64 `json:"percentage"`
			Population int     `json:"population"`
		} `json:"color"`
		Latitude                   float64       `json:"latitude"`
		Longitude                  float64       `json:"longitude"`
		Latlon                     string        `json:"latlon"`
		IsOnView                   bool          `json:"is_on_view"`
		GalleryTitle               string        `json:"gallery_title"`
		GalleryID                  int           `json:"gallery_id"`
		ArtworkTypeTitle           interface{}   `json:"artwork_type_title"`
		ArtworkTypeID              interface{}   `json:"artwork_type_id"`
		DepartmentTitle            string        `json:"department_title"`
		DepartmentID               string        `json:"department_id"`
		ArtistID                   int           `json:"artist_id"`
		ArtistTitle                string        `json:"artist_title"`
		AltArtistIds               []interface{} `json:"alt_artist_ids"`
		ArtistIds                  []int         `json:"artist_ids"`
		ArtistTitles               []string      `json:"artist_titles"`
		CategoryIds                []string      `json:"category_ids"`
		CategoryTitles             []string      `json:"category_titles"`
		ArtworkCatalogueIds        []interface{} `json:"artwork_catalogue_ids"`
		TermTitles                 []string      `json:"term_titles"`
		StyleID                    string        `json:"style_id"`
		StyleTitle                 string        `json:"style_title"`
		AltStyleIds                []string      `json:"alt_style_ids"`
		StyleIds                   []string      `json:"style_ids"`
		StyleTitles                []string      `json:"style_titles"`
		ClassificationID           string        `json:"classification_id"`
		ClassificationTitle        string        `json:"classification_title"`
		AltClassificationIds       []string      `json:"alt_classification_ids"`
		ClassificationIds          []string      `json:"classification_ids"`
		ClassificationTitles       []string      `json:"classification_titles"`
		SubjectID                  string        `json:"subject_id"`
		AltSubjectIds              []string      `json:"alt_subject_ids"`
		SubjectIds                 []string      `json:"subject_ids"`
		SubjectTitles              []string      `json:"subject_titles"`
		MaterialID                 string        `json:"material_id"`
		AltMaterialIds             []interface{} `json:"alt_material_ids"`
		MaterialIds                []string      `json:"material_ids"`
		MaterialTitles             []string      `json:"material_titles"`
		TechniqueID                string        `json:"technique_id"`
		AltTechniqueIds            []string      `json:"alt_technique_ids"`
		TechniqueIds               []string      `json:"technique_ids"`
		TechniqueTitles            []string      `json:"technique_titles"`
		ThemeTitles                []string      `json:"theme_titles"`
		ImageID                    string        `json:"image_id"`
		AltImageIds                []interface{} `json:"alt_image_ids"`
		DocumentIds                []string      `json:"document_ids"`
		SoundIds                   []string      `json:"sound_ids"`
		VideoIds                   []interface{} `json:"video_ids"`
		TextIds                    []interface{} `json:"text_ids"`
		SectionIds                 []interface{} `json:"section_ids"`
		SectionTitles              []interface{} `json:"section_titles"`
		SiteIds                    []int         `json:"site_ids"`
		SuggestAutocompleteBoosted string        `json:"suggest_autocomplete_boosted"`
		SuggestAutocompleteAll     []struct {
			Input    []string `json:"input"`
			Contexts struct {
				Groupings []string `json:"groupings"`
			} `json:"contexts"`
			Weight int `json:"weight,omitempty"`
		} `json:"suggest_autocomplete_all"`
		LastUpdatedSource string `json:"last_updated_source"`
		LastUpdated       string `json:"last_updated"`
		Timestamp         string `json:"timestamp"`
	} `json:"data"`
	Info struct {
	} `json:"info"`
	Config struct {
	} `json:"config"`
}

// Artwork creates an Artwork from an ArtworkResponse
func (ar ArtworkResult) Artwork() Artwork {
	return Artwork{
		ID:          ar.Data.ID,
		URL:         ar.Data.APILink,
		Title:       ar.Data.Title,
		Description: ar.Data.Description,
		Date:        ar.Data.DateDisplay,
		Artist:      ar.Data.ArtistDisplay,
	}
}

// Artwork is the struct representation of an Artwork response
type Artwork struct {
	ID          int
	URL         string
	Title       string
	Artist      string
	Date        string
	Description string
}

// PrettyString cretes a pretty string of the Comic
func (a Artwork) PrettyString() string {
	p := fmt.Sprintf(
		"Title: %s\nArtist: %s\nDate: %s\nDescription: %s\nImage Id: %d\n",
		a.Title, a.Artist, a.Date, a.Description, a.ID)
	return p
}

// JSON returns the JSON representation of the comic
func (a Artwork) JSON() string {
	cJSON, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(cJSON)
}
