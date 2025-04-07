package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("Badger Workbench")

	// Call SetupUI to initialize the UI
	SetupUI(myWindow)

	// Show and run the application
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
