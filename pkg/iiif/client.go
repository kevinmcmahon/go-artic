package iiif

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type ImageID string

// Client is the client for the Art Institute Chicago Image API
type Client struct {
	client       *http.Client
	baseImageURL string
	verbose      bool
}

const (
	// BaseImageURL of Art Institute Chicago Image API
	BaseImageURL string = "https://www.artic.edu/iiif"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout = 30 * time.Second
)

// New creates a new image client
func New(verbose bool) *Client {
	return &Client{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseImageURL: BaseImageURL,
		verbose:      verbose,
	}
}

// SetTimeout overrides the default ClientTimeout
func (c *Client) SetTimeout(d time.Duration) {
	c.client.Timeout = d
}

// builds the correct url
func (c *Client) buildUrl(id ImageID) string {
	return fmt.Sprintf("%s/2/%s/full/843,/0/default.jpg", c.baseImageURL, id)
}

func (c *Client) SaveToDisk(imageId ImageID, savePath string) error {

	var url = c.buildUrl(imageId)

	if c.verbose {
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
