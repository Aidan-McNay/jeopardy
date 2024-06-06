//========================================================================
// toolbar.go
//========================================================================
// A toolbar, for providing icon-based actions for the user
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"image/color"
	"jeopardy/logic"
	"jeopardy/style"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// Keep track of whether we can create a popup
//------------------------------------------------------------------------
// Pop-ups are done by the user sequentially, so we don't need a mutex

var canCreatePopup bool = true

func canOpenPopup() bool { return canCreatePopup }
func openPopup()         { canCreatePopup = false }
func closePopup()        { canCreatePopup = true }

//------------------------------------------------------------------------
// New Board Creation
//------------------------------------------------------------------------

func promptNewBoard(win fyne.Window) {
	if !canOpenPopup() {
		return
	}
	openPopup()

	newName := widget.NewEntry()
	newName.Validator = validation.NewRegexp(`^.+$`, "Board must have a non-empty name")

	items := []*widget.FormItem{
		widget.NewFormItem("Board Name", newName),
	}
	onConfirm := func(b bool) {
		closePopup()
		if !b {
			return
		}
		logic.NewBoard(newName.Text)
	}
	prompt := dialog.NewForm("New Board", "Create New Board", "Cancel", items,
		onConfirm, win)

	var height float32 = prompt.MinSize().Height
	var width float32 = 400
	newSize := fyne.NewSize(width, height)
	prompt.Resize(newSize)

	prompt.Show()
}

//------------------------------------------------------------------------
// File Manipulation
//------------------------------------------------------------------------

func loadFromFile(win fyne.Window) {
	if !canOpenPopup() {
		return
	}
	openPopup()

	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		closePopup()
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if reader == nil {
			// Cancelled
			return
		}

		logic.LoadCurrBoard(reader)
	}, win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".jpdy"}))
	fd.Show()
}

func saveToFile(win fyne.Window) {
	if !canOpenPopup() {
		return
	}
	openPopup()

	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		closePopup()
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if writer == nil {
			// Cancelled
			return
		}

		logic.SaveCurrBoard(writer)
	}, win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".jpdy"}))
	fd.Show()
}

//------------------------------------------------------------------------
// Help Window
//------------------------------------------------------------------------

func showHelp() {
	titleText := canvas.NewText("Jeopardy", color.White)
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleText.TextSize = 30

	authorText := widget.NewLabel("Author: Aidan McNay")
	authorText.Alignment = fyne.TextAlignCenter

	descriptionText := widget.NewLabel(
		"A platform for creating, storing, and running Jeopardy-like games. " +
			"Users can customize questions and categories, as well as the " +
			"overall theme of the game",
	)
	descriptionText.Alignment = fyne.TextAlignCenter
	descriptionText.Wrapping = fyne.TextWrapWord

	sourceURL, _ := url.Parse("https://github.com/Aidan-McNay/jeopardy")
	sourceLink := widget.NewHyperlink("Reference/Source code",
		sourceURL,
	)
	sourceLink.Alignment = fyne.TextAlignCenter

	disclaimerText := widget.NewLabel(
		"Disclaimer: This project is not affiliated with JEOPARDY!™",
	)
	disclaimerText.Alignment = fyne.TextAlignCenter
	disclaimerText.TextStyle = fyne.TextStyle{Italic: true}

	content := container.NewVBox(
		titleText,
		authorText,
		descriptionText,
		sourceLink,
		disclaimerText,
	)
	content.Resize(fyne.NewSize(200, 100))

	w := fyne.CurrentApp().NewWindow("Information")
	w.SetContent(content)
	w.Show()
}

//------------------------------------------------------------------------
// Main Toolbar
//------------------------------------------------------------------------

func Toolbar(win fyne.Window) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			promptNewBoard(win)
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			loadFromFile(win)
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			saveToFile(win)
		}),
		widget.NewToolbarAction(theme.ColorPaletteIcon(), func() {
			log.Println("Style clicked")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			log.Println("Play clicked")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			style.ColorDialog(win)
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			showHelp()
		}),
	)
}
