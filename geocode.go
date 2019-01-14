package main

import (
	"fmt"
	"os"

	geo "github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
	"github.com/codingsince1985/geo-golang/mapquest/nominatim"
	"github.com/rwcarlsen/goexif/exif"
)

var (
	googleAPIKey   = os.Getenv("GOOGLE_API_KEY")
	mapquestAPIKey = os.Getenv("MAPQUEST_API_KEY")
)

func imageLatLong(f *os.File) (float64, float64, error) {
	x, err := exif.Decode(f)
	if err != nil {
		return 0, 0, err
	}
	return x.LatLong()
}

func revGeocode(provider string, lat, long float64) error {
	var geocoder geo.Geocoder
	switch provider {
	case "osm":
		geocoder = nominatim.Geocoder(mapquestAPIKey)
	case "google":
		geocoder = google.Geocoder(googleAPIKey)
	default:
		return fmt.Errorf("unknown provider '%s'", provider)
	}

	address, err := geocoder.ReverseGeocode(lat, long)
	if err != nil {
		return err
	}
	fmt.Printf("addr: %s\n", address)
	return nil
}
