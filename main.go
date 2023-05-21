package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {
	dir := flag.String("i", "", "Directory to process")
	outputDir := flag.String("o", "", "Directory to output to")
	width := flag.Int("x", 128, "crop in the width of the image (defaults to 128)")
	height := flag.Int("y", 128, "crop the height of the image (defaults to 128)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 && *dir == "" {
		fmt.Println("Usage: go run crop.go [-i=<dir>] [-o=<outputDir>] [-x=<pixels to crop>] [-y=<pixels to crop>] <image>")
		os.Exit(1)
	}

	if *dir != "" {
		files, err := ioutil.ReadDir(*dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Mkdir(*outputDir, 0777)

		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".png") {
				processImage(filepath.Join(*dir, f.Name()), *outputDir, *width, *height)
			}
		}
	} else {
		processImage(args[0], *outputDir, *width, *height)
	}
}

func processImage(path string, outputDir string, width, height int) {
	src, err := imaging.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	src = imaging.Crop(src, image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	src = imaging.Resize(src, 128, 128, imaging.Lanczos)

	outputPath := strings.TrimSuffix(filepath.Base(path), ".png") + "_128x128.png"
	outputPath = filepath.Join(outputDir, outputPath)
	err = imaging.Save(src, outputPath, imaging.PNGCompressionLevel(png.DefaultCompression))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Image processed: %s\n", outputPath)
}
