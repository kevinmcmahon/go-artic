package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ArtworkID int64

// Client is the client for the Art Institute Chicago Image API
type Client struct {
	client  *http.Client
	baseURL string
	verbose bool
}

const (
	// BaseURL of Art Institute Chicago API
	BaseURL string = "https://aggregator-data.artic.edu/api/v1"

	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout = 30 * time.Second
)

// New creates a new image client
func New(verbose bool) *Client {
	return &Client{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
		verbose: verbose,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *Client) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

// builds the correct url
func (hc *Client) buildURL(id ArtworkID) string {
	return fmt.Sprintf("%s/artworks/%d", hc.baseURL, id)
}

func (hc *Client) Fetch(id ArtworkID) (ArtworkResponse, error) {
	var artworkResponse ArtworkResponse

	var url = hc.buildURL(id)
	if hc.verbose {
		fmt.Printf("[DEBUG] url : %s\n", url)
	}
	resp, err := hc.client.Get(url)
	if err != nil {
		return artworkResponse, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&artworkResponse); err != nil {
		return artworkResponse, err
	}
	return artworkResponse, nil
}
