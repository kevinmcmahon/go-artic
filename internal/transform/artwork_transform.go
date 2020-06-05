package transform

import (
	"github.com/kevinmcmahon/go-artic/internal/model"
	"github.com/kevinmcmahon/go-artic/pkg/artic"
)

// MakeArtworkFromResponse creates an Artwork
func MakeArtworkFromResponse(ar artic.ArtworkResponse) model.Artwork {
	return model.Artwork{
		ID:          ar.Data.ID,
		URL:         ar.Data.APILink,
		Title:       ar.Data.Title,
		Description: ar.Data.Description,
		Date:        ar.Data.DateDisplay,
		Artist:      ar.Data.ArtistDisplay,
	}
}
