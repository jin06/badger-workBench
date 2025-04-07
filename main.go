package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("BadgerWorkBench")

	SetupUI(myWindow)

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
