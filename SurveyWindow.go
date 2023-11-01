// Package comment

package main

import (
	"math/rand"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// List of functions:
// createButton()
// getResults()
// shuffle()
// buildImageWindow() -> shuffle() -> loadImgs()
// loadImageWindow()
// refreshImageWindow()
// showImageWindow()

func createButton(img *canvas.Image, sharedInt *int) *widget.Button {
	button := widget.NewButton("", nil)
	button.OnTapped = func() {
		if *sharedInt == -1 {
			*sharedInt = 0
			if img.Hidden {
				img.Show()
				button.SetText("")
			}
			return
		}

		if img.Visible() {
			*sharedInt++
			button.SetText(strconv.Itoa(*sharedInt))
			img.Hide()
		}
	}
	return button
}

type fixedSizeLayout struct {
	fixedWidth, fixedHeight float32
}

func (f *fixedSizeLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, o := range objects {
		o.Resize(fyne.NewSize(f.fixedWidth, f.fixedHeight))
		o.Move(fyne.NewPos(0, 0))
	}
}

func (f *fixedSizeLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(f.fixedWidth, f.fixedHeight)
}

func getResults(buttons []*widget.Button) {
	for i, button := range buttons {
		userChoices[i], _ = strconv.Atoi(button.Text)
	}
}

// Fisher-Yates Shuffle Algorithm to permute the given slice
func shuffle() {
	for i := len(permutation) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                                           // Generate a random index between 0 and i (inclusive)
		permutation[i], permutation[j] = permutation[j], permutation[i] // Swap elements at i and j
	}
}

func buildImageWindow(myApp fyne.App, imageWindow fyne.Window) (*widget.Label, *fyne.Container, *widget.Button, *widget.Button, *widget.Label) {

	hbox := container.NewHBox()
	shuffle()
	loadImgs()
	sharedInt := 0
	var buttons []*widget.Button

	for _, img := range imgs {
		button := createButton(img, &sharedInt)
		buttons = append(buttons, button)
		box := container.NewPadded(button, img)
		container := container.New(&fixedSizeLayout{150, 150}, box)
		hbox.Add(container)
	}

	var prompt *widget.Label
	if simpleComplex == 0 {
		prompt = widget.NewLabel("Prompt: " + simplePrompts[simple[fileIndex+userIndex]-1])
		prompt.Wrapping = fyne.TextWrapWord
	} else {
		prompt = widget.NewLabel("Prompt: " + complexPrompts[complex[fileIndex+userIndex]-1])
		prompt.Wrapping = fyne.TextWrapWord
	}

	clearBtn := widget.NewButton("Clear", func() {
		// Set all button texts in the HBox to "Start"
		sharedInt = -1
		for _, button := range buttons {
			sharedInt = -1
			button.OnTapped()
		}
	})

	nextBtn := widget.NewButton("Next", func() {
		if sharedInt == 8 {
			userIndex++
			getResults(buttons)
			saveResults()
			if userIndex == 10 && simpleComplex == 1 {
				if instruction == 3 {
					customDialog := dialog.NewCustom("Finished", "Exit", widget.NewLabel("Thank you for your time."), imageWindow)
					customDialog.Show()
					imageWindow.Close()
				} else {
					instruction++
					userIndex = 0
					simpleComplex = 0
				}
			}

			if userIndex == 10 && simpleComplex == 0 {
				userIndex = 0
				simpleComplex = 1
			}
			refreshImageWindow(myApp, imageWindow)
		}
	})

	label := widget.NewLabel(instructions[instruction])
	label.Wrapping = fyne.TextWrapWord

	return prompt, hbox, clearBtn, nextBtn, label
}

func loadImageWindow(myApp fyne.App, imageWindow fyne.Window) fyne.CanvasObject {
	prompt, hbox, clearBtn, nextBtn, labelInstructions := buildImageWindow(myApp, imageWindow)
	// Add the button to the bottom right
	vbox := container.NewVBox(prompt, hbox, clearBtn, nextBtn, labelInstructions)
	return vbox
}

func refreshImageWindow(myApp fyne.App, imageWindow fyne.Window) {
	imageWindow.SetContent(loadImageWindow(myApp, imageWindow))
	imageWindow.Resize(fyne.NewSize(800, 1600))
	imageWindow.Show()
}

func showImageWindow(myApp fyne.App) {
	imageWindow := myApp.NewWindow("Image Window")
	imageWindow.SetContent(loadImageWindow(myApp, imageWindow))
	imageWindow.Resize(fyne.NewSize(800, 1600))
	imageWindow.Show()
}
