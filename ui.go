package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func SetupUI(win fyne.Window) {
	keys := []string{}
	valueDisplay := widget.NewMultiLineEntry()
	valueDisplay.Wrapping = fyne.TextWrapWord
	valueDisplay.SetPlaceHolder("🤘 Display Area: Key-Value Content")

	valueDisplay.TextStyle = fyne.TextStyle{Bold: true}
	valueDisplay.Wrapping = fyne.TextWrapWord

	keyList := widget.NewList(
		func() int { return len(keys) },
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			lbl.TextStyle = fyne.TextStyle{Bold: true}
			return lbl
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(keys[i])
		},
	)

	loadLastOpenedDirectory()

	if lastOpenedDir != "" {
		err := OpenDB(lastOpenedDir)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			keys, _ = GetAllKeys()
			refreshKeys(keyList, keys)
		}
	}

	loadBtn := widget.NewButtonWithIcon("Load Badger 🤘", theme.FolderOpenIcon(), func() {

		// Use lastOpenedDir if it exists, else use the current working directory
		initialDir := lastOpenedDir
		if initialDir == "" {
			initialDir = "./" // Fallback to the current directory if no last opened dir
		}

		dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
			if dir == nil {
				return
			}
			err = OpenDB(dir.Path())
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			saveLastOpenedDirectory(dir.Path())
			keys, _ = GetAllKeys()
			refreshKeys(keyList, keys)
		}, win)
	})

	refreshBtn := widget.NewButtonWithIcon("Refresh Keys 🔁", theme.ViewRefreshIcon(), func() {
		keys, _ = GetAllKeys()
		refreshKeys(keyList, keys)
	})

	// Left is the list, right is the value display, middle is resizable
	split := container.NewHSplit(keyList, valueDisplay)
	split.Offset = 0.3

	header := container.NewHBox(loadBtn, refreshBtn)

	ui := container.NewBorder(header, nil, nil, nil, split)

	// Set window content
	win.SetContent(ui)

	// Set window style
	win.Resize(fyne.NewSize(1000, 600))
	win.SetTitle("Badger Workbench")

	keyList.OnSelected = func(id int) {
		if id >= len(keys) {
			return
		}
		val, err := GetValue(keys[id])
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		valueDisplay.SetText(val)
	}
}

var lastOpenedDir string

// Load the last opened directory from a simple text file
func loadLastOpenedDirectory() {
	if _, err := os.Stat("last_opened_dir.txt"); err == nil {
		data, err := os.ReadFile("last_opened_dir.txt")
		if err == nil {
			lastOpenedDir = string(data)
		}
	}
}

// Save the last opened directory to a simple text file
func saveLastOpenedDirectory(dir string) {
	err := os.WriteFile("last_opened_dir.txt", []byte(dir), 0644)
	if err != nil {
		fmt.Println("Error saving last opened directory:", err)
	}
}

func refreshKeys(list *widget.List, keys []string) {
	// Update the List with new keys
	list.Length = func() int { return len(keys) }
	list.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(keys[i])
	}
	list.Refresh() // Refresh the list view
}
