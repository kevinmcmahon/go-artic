package artic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	// BaseURL of Art Institute Chicago API
	BaseURL string = "https://aggregator-data.artic.edu/api/v1"
	// BaseURL of Art Institute Chicago Image API
	BaseImageURL string = "https://www.artic.edu/iiif"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout = 30 * time.Second
)

// ArtworkID is the id of the artwork
type ArtworkID int64

// Client is the client for the Artic Institute Chicago API
type Client struct {
	client       *http.Client
	baseURL      string
	baseImageURL string
	verbose      bool
}

// New creates a new ArticClient
func New(verbose bool) *Client {
	return &Client{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL:      BaseURL,
		baseImageURL: BaseImageURL,
		verbose:      verbose,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *Client) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

// Fetch retrieves the artwork as per provided artwork id
func (hc *Client) Fetch(id ArtworkID, save bool) (ArtworkResponse, error) {
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

	if save {
		if err := hc.SaveToDisk(artworkResponse.Data.ImageID, "."); err != nil {
			fmt.Println("Failed to save image!")
			return artworkResponse, err
		}
	}
	return artworkResponse, nil
}

// builds the correct url
func (hc *Client) buildURL(id ArtworkID) string {
	return fmt.Sprintf("%s/artworks/%d", hc.baseURL, id)
}

// builds the correct url
func (hc *Client) buildImageURL(id string) string {
	return fmt.Sprintf("%s/2/%s/full/843,/0/default.jpg", hc.baseImageURL, id)
}

func (hc *Client) SaveToDisk(imageId string, savePath string) error {

	var url = hc.buildImageURL(imageId)

	if hc.verbose {
		fmt.Printf("[DEBUG] image url : %s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s", absSavePath, path.Base(url))

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
