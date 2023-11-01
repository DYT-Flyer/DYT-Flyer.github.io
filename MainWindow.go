package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// List of functions:
// buildWindow -> processAndStore() -> showImageWindow() -> ... ->
// processAndStore -> userIDExists() -> appendToFile()

// username entry, submit button, instruction label
func buildWindow(myWindow fyne.Window, myApp fyne.App) *fyne.Container {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter Student ID")

	submitButton := widget.NewButton("Submit", func() {
		userID = usernameEntry.Text
		success := processAndStore()
		myWindow.Hide()
		if success {
			showImageWindow(myApp)
		} else {
			customDialog := dialog.NewCustom("Error", "OK", widget.NewLabel("Name already exists"), myWindow)
			customDialog.Show()
		}
		myWindow.Show()
	})

	label := widget.NewLabel("Enter your name above, then click submit. You will be shown a series of images " +
		"and asked to rank them by the provided criteria by clicking on the photos in order from best to worst. " +
		"Once you have ranked all the photos, click next and then finally submit. Thank you.")

	label.Wrapping = fyne.TextWrapWord
	// Add Instructions
	content := container.NewVBox(usernameEntry, submitButton, label)

	return content
}

// submit button action
// references userIDExists and appendToFile in Json.go
// userIDExists returns rows which I will use as unique identifier
func processAndStore() bool {
	usernameExists, err := userIDExists()
	fileIndex = (rows % 5) * 10
	fmt.Println(rows)

	if err != nil {
		log.Fatalf("Error: %v", err)
		return false
	}

	if usernameExists {
		// Decided to allow duplicate usernames
		return true
	}

	err = appendToFile()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return false
	}

	return true
}
