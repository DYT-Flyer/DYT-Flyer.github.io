package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// List of functions:
// loadImage()
// loadImgs() -> loadImgs()

// Load an image from a file
func loadImage(file string) *canvas.Image {
	img := canvas.NewImageFromFile(file)
	return img
}

func loadImgs() {
	//fileindex + userindex
	var files []string
	var folder string
	fmt.Println(simpleComplex)
	if simpleComplex == 0 {
		files = simpleFiles[simple[fileIndex+userIndex]-1]
		folder = simpleFolder
	} else {
		files = complexFiles[complex[fileIndex+userIndex]-1]
		folder = complexFolder
	}

	for i, p := range permutation {
		img := loadImage(folder + files[p])
		img.FillMode = canvas.ImageFillContain
		img.SetMinSize(fyne.NewSize(200, 200))
		imgs[i] = img
	}
}
