package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var opts struct {
	folderPath string
	provider   string
}

func main() {
	flag.StringVar(&opts.folderPath, "f", "", "path to folder with images to process")
	flag.StringVar(&opts.provider, "p", "osm", "choose from: 'google', 'osm'")
	flag.Parse()

	fileList, err := ioutil.ReadDir(opts.folderPath)
	if err != nil {
		log.Fatalf("unable to read folder %s: %s", opts.folderPath, err)
	}

	for _, f := range fileList {
		fileName := path.Join(opts.folderPath, f.Name())
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("could not open file: %s", err)
		}

		lat, long, err := imageLatLong(file)
		if err != nil {
			log.Fatalf("could not get image lat/long from exif: %s", err)
		}
		file.Close()

		if err := revGeocode(opts.provider, lat, long); err != nil {
			log.Fatalf("rev geocode err: %s", err)
		}
	}

}
