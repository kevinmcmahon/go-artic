package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kevinmcmahon/go-artic/model"
)

const (
	// BaseURL of xkcd
	BaseURL string = "https://aggregator-data.artic.edu/api/v1"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 30 * time.Second
	// DefaultArtwork is `A Sunday on La Grande Jatte`
	DefaultArtwork int = 27922
)

// ArtworkID is the id of the artwork
type ArtworkID int

// ArticClient is the client for the Artic Institute Chicago API
type ArticClient struct {
	client  *http.Client
	baseURL string
	verbose bool
}

// NewArticClient creates a new ArticClient
func NewArticClient() *ArticClient {
	return &ArticClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *ArticClient) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

// SetVerbose to true for debug messages
func (hc *ArticClient) SetVerbose(v bool) {
	hc.verbose = v
}

// Fetch retrieves the artwork as per provided artwork id
func (hc *ArticClient) Fetch(id ArtworkID, save bool) (model.Artwork, error) {
	var url = hc.buildURL(id)
	if hc.verbose {
		fmt.Printf("[DEBUG] url : %s\n", url)
	}
	resp, err := hc.client.Get(url)
	if err != nil {
		return model.Artwork{}, err
	}
	defer resp.Body.Close()

	var artworkResult model.ArtworkResult
	if err := json.NewDecoder(resp.Body).Decode(&artworkResult); err != nil {
		return model.Artwork{}, err
	}

	if save {
		// if err := hc.SaveToDisk(artworkResult.Img, "."); err != nil {
		// 	fmt.Println("Failed to save image!")
		// }
	}
	return artworkResult.Artwork(), nil
}

// builds the correct url
func (hc *ArticClient) buildURL(id ArtworkID) string {
	return fmt.Sprintf("%s/artworks/%d", hc.baseURL, id)
}
