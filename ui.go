package main

import (
	"fmt"
	"os"
	"strconv"

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
	// Label to display the opened folder path
	folderPathLabel := widget.NewLabel("No folder selected")
	folderPathLabel.Wrapping = fyne.TextWrapWord
	folderPathLabel.TextStyle = fyne.TextStyle{Bold: true}

	if lastOpenedDir != "" {
		err := OpenDB(lastOpenedDir)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			keys, _ = GetAllKeys()
			refreshKeys(keyList, keys)
		}
		folderPathLabel.SetText(lastOpenedDir) // Update the folder path label
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
			folderPathLabel.SetText(dir.Path()) // Update the folder path label
		}, win)
	})

	refreshBtn := widget.NewButtonWithIcon("Refresh Keys 🔁", theme.ViewRefreshIcon(), func() {
		keys, _ = GetAllKeys()
		refreshKeys(keyList, keys)
	})

	keyEntry := widget.NewEntry()
	keyEntry.SetPlaceHolder("Enter Key")

	ttlEntry := widget.NewEntry()
	ttlEntry.SetPlaceHolder("Enter TTL (seconds)") // Used to display and set TTL

	valueEntry := widget.NewMultiLineEntry()
	valueEntry.SetPlaceHolder("Enter Value")
	keyList.OnSelected = func(id int) {
		if id >= 0 && id < len(keys) {
			selectedKey := keys[id]
			val, err := GetValue(selectedKey)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			valueDisplay.SetText(val)
			keyEntry.SetText(selectedKey)
			valueEntry.SetText(val)

			// Retrieve and display TTL
			ttl, err := GetKeyTTL(selectedKey)
			if err != nil {
				ttlEntry.SetText("Error retrieving TTL")
			} else {
				ttlEntry.SetText(fmt.Sprintf("%d", ttl))
			}
		}
	}

	submitBtn := widget.NewButton("Submit", func() {
		key := keyEntry.Text
		value := valueEntry.Text
		ttlText := ttlEntry.Text

		if key == "" {
			dialog.ShowInformation("Invalid Input", "Key cannot be empty.", win)
			return
		}

		var ttl uint64
		if ttlText != "" {
			parsedTTL, err := strconv.ParseUint(ttlText, 10, 64)
			if err != nil {
				dialog.ShowError(fmt.Errorf("Invalid TTL value"), win)
				return
			}
			ttl = parsedTTL
		}

		err := SetValueWithTTL(key, value, ttl)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			// dialog.ShowInformation("Success", "Key-Value pair added/updated successfully.", win)
			keys, _ = GetAllKeys()
			refreshKeys(keyList, keys)
		}
	})

	inputForm := container.NewVBox(
		widget.NewLabel("Add or Update Key-Value"),
		keyEntry,
		ttlEntry, // Add TTL input field
		valueEntry,
		submitBtn,
	)

	// Left is the list, right is the value display, middle is resizable
	split := container.NewHSplit(keyList, valueDisplay)
	split.Offset = 0.3

	// folderPathLabel.Resize(fyne.NewSize(500, folderPathLabel.MinSize().Height)) // Set a minimum width for the label

	folderContainer := container.NewVBox(folderPathLabel)
	header := container.NewBorder(
		nil, // No top border
		nil, // No bottom border
		nil, // No left border
		nil,
		container.NewVBox(container.NewHBox(loadBtn, refreshBtn), folderContainer),
	)

	ui := container.NewBorder(header, inputForm, nil, nil, split)

	// Set window content
	win.SetContent(ui)

	// Set window style
	win.Resize(fyne.NewSize(1000, 600))
	win.SetTitle("Badger Workbench")
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
