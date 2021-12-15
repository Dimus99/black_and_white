package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func readImg(path string) image.Image {
	existingImageFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer existingImageFile.Close()
	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		log.Fatal(err)
	}
	return imageData
}

func writeImg(newImagePath string, img *image.RGBA) {
	file, err := os.Create(newImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
