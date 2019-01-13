package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var opts struct {
	filename string
	provider string
}

func main() {
	flag.StringVar(&opts.filename, "f", "", "path to image to process")
	flag.StringVar(&opts.provider, "p", "googlemaps", "choose from: 'google', 'osm'")
	flag.Parse()

	if opts.filename == "" {
		fmt.Printf("Missing filename value!\n")
		return
	}

	file, err := os.Open(opts.filename)
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer file.Close()

	lat, long, err := imageLatLong(file)
	if err != nil {
		log.Fatalf("could not get image lat/long from exif: %s", err)
	}

	if err := revGeocode(opts.provider, lat, long); err != nil {
		log.Fatalf("rev geocode err: %s", err)
	}
}
