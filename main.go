package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	// Create an app
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	// Create a window for the app
	myWindow := myApp.NewWindow("Username and Text")

	content := buildWindow(myWindow, myApp)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetFixedSize(true)
	// Set the content of the window to be our box
	myWindow.SetContent(content)

	fmt.Println(dir)
	// Show the window and run the application
	myWindow.ShowAndRun()
}
