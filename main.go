package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kevinmcmahon/go-artic/client"
)

// DefaultArtwork is `A Sunday on La Grande Jatte`
const DefaultArtwork int = 27992

func main() {
	artID := flag.Int(
		"i", int(DefaultArtwork), "Artwork id to fetch",
	)
	clientTimeout := flag.Int64(
		"t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
	)
	saveImage := flag.Bool(
		"s", false, "Save image to current directory",
	)
	outputType := flag.String(
		"o", "text", "Print output in format: text/json",
	)
	verbose := flag.Bool(
		"v", false, "Verbose debug messages",
	)
	flag.Parse()

	articClient := client.New(*verbose)
	articClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	artwork, err := articClient.Fetch(client.ArtworkID(*artID), *saveImage)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(artwork.JSON())
	} else {
		fmt.Println(artwork.PrettyString())
	}
}
