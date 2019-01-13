package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"googlemaps.github.io/maps"
)

var GoogleMapsKey = os.Getenv("GOOGLE_MAPS_API_KEY")

var opts struct {
	filename string
}

func main() {
	flag.StringVar(&opts.filename, "f", "", "full path to file to check exif data")
	flag.Parse()

	if opts.filename == "" {
		fmt.Printf("Missing filename value!")
		return
	}

	c, err := maps.NewClient(maps.WithAPIKey(GoogleMapsKey))
	if err != nil {
		log.Fatalf("could not initialize maps client: %s", err)
	}

	locationForImage(c, opts.filename)
}

func locationForImage(client *maps.Client, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	lat, long, err := x.LatLong()
	if err != nil {
		log.Fatalf("could not get lat/long: %s", err)
		return
	}
	fmt.Printf("%1.8f, %1.8f\n", lat, long)

	results, err := client.ReverseGeocode(context.TODO(), &maps.GeocodingRequest{
		LatLng: &maps.LatLng{lat, long},
	})
	if err != nil {
		log.Fatalf("error on reverse geocode: %s", err)
		return
	}

	for _, r := range results {
		fmt.Printf("Address: %s\n", r.FormattedAddress)
	}
}
