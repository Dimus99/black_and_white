package main

import (
	"fmt"
	"image"
	"image/color"
	"time"
)

func convertImageToBlackWhite(size image.Point, imageData image.Image, newImg *image.RGBA) {
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := imageData.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			r := float64(originalColor.R) * 0.92126 // magic haha
			g := float64(originalColor.G) * 0.97152
			b := float64(originalColor.B) * 0.90722
			grey := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: grey, G: grey, B: grey, A: originalColor.A,
			}
			newImg.Set(x, y, c)
		}
	}
}

func converterImageFileToBlackWhite(chanel chan [2]string, dbChan chan [4]string, out chan int) {
	for inp := range chanel {
		currentPath := inp[0]
		newImagePath := inp[1]
		start := time.Now()
		imageData := readImg(currentPath)
		size := imageData.Bounds().Size()
		newImg := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
		convertImageToBlackWhite(size, imageData, newImg)
		writeImg(newImagePath, newImg)
		duration := time.Since(start)
		fmt.Println(newImagePath)
		dbChan <- [4]string{currentPath, newImagePath,
			size.String(), duration.String()}
	}
	out <- 1
}
