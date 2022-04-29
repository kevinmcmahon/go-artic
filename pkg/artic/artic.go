package artic

import (
	"github.com/kevinmcmahon/go-artic/internal/transform"
	"github.com/kevinmcmahon/go-artic/pkg/data"
	"github.com/kevinmcmahon/go-artic/pkg/iiif"
	"github.com/kevinmcmahon/go-artic/pkg/model"
	"time"
)

const (
	DefaultClientTimeout = 30 * time.Second
)

// ArtworkID is the id of the artwork
type ArtworkID int64
type ImageID string

// Repo is the client for the Art Institute Chicago API
type Repo struct {
	dataClient  *data.Client
	imageClient *iiif.Client
}

// New creates a new client
func New(verbose bool) *Repo {
	return &Repo{
		dataClient:  data.New(verbose),
		imageClient: iiif.New(verbose),
	}
}

// SetTimeout overrides the default ClientTimeout
func (repo *Repo) SetTimeout(d time.Duration) {
	repo.dataClient.SetTimeout(d)
	repo.imageClient.SetTimeout(d)
}

// Load retrieves the artwork as per provided artwork id
func (repo *Repo) Load(id ArtworkID, save bool) (model.Artwork, error) {
	var artwork model.Artwork
	var dataResponse data.ArtworkResponse
	var err error
	artworkID := data.ArtworkID(id)
	dataResponse, err = repo.dataClient.Fetch(artworkID)
	if err != nil {
		return artwork, err
	}

	artwork = transform.MakeArtworkFromResponse(dataResponse)

	if !save {
		return artwork, nil
	}

	imageID := iiif.ImageID(dataResponse.Data.ImageID)
	err = repo.imageClient.SaveToDisk(imageID, ".")
	return artwork, err
}
