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
	"jeopardy/assets"
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

var currURI fyne.URI = nil
var haveABoard = false

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
		currURI = nil
		haveABoard = true
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

		currURI = reader.URI()
		haveABoard = true
		logic.LoadCurrBoard(reader)
	}, win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".jpdy"}))
	fd.Show()
}

func saveToFile(win fyne.Window, forceSaveAs bool) {
	if !haveABoard {
		// No board to save
		return
	}
	if (currURI != nil) && !forceSaveAs {
		writer, _ := storage.Writer(currURI)
		logic.SaveCurrBoard(writer)
		return
	}
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
		currURI = writer.URI()
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
	newBoardAction := widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		promptNewBoard(win)
	})
	loadBoardAction := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		loadFromFile(win)
	})
	saveBoardAction := widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		saveToFile(win, false)
	})
	styleBoardAction := widget.NewToolbarAction(theme.ColorPaletteIcon(), func() {
		log.Println("Style clicked")
	})
	runBoardAction := widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
		log.Println("Play clicked")
	})
	otherThemeAction := widget.NewToolbarAction(theme.SettingsIcon(), func() {})
	settingsAction := widget.NewToolbarAction(theme.SettingsIcon(), func() {
		style.ColorDialog(win)
	})
	helpAction := widget.NewToolbarAction(theme.HelpIcon(), func() {
		showHelp()
	})

	refreshIcons := func() {
		newBoardAction.SetIcon(theme.ContentAddIcon())
		loadBoardAction.SetIcon(theme.FolderOpenIcon())
		saveBoardAction.SetIcon(theme.DocumentSaveIcon())
		styleBoardAction.SetIcon(theme.ColorPaletteIcon())
		runBoardAction.SetIcon(theme.MediaPlayIcon())
		settingsAction.SetIcon(theme.SettingsIcon())
		helpAction.SetIcon(theme.HelpIcon())

		switch style.GetVariant() {
		case theme.VariantDark:
			otherThemeAction.SetIcon(theme.NewThemedResource(assets.ResourceSunPng))
		default:
			otherThemeAction.SetIcon(theme.NewThemedResource(assets.ResourceMoonPng))
		}
	}
	refreshIcons()

	otherThemeAction.OnActivated = func() {
		switch style.GetVariant() {
		case theme.VariantDark:
			style.SetVariant(theme.VariantLight)
		default:
			style.SetVariant(theme.VariantDark)
		}
		refreshIcons()
		logic.BoardChange()
		style.StoreColorPreferences(fyne.CurrentApp())
	}

	return widget.NewToolbar(
		newBoardAction,
		loadBoardAction,
		saveBoardAction,
		styleBoardAction,
		widget.NewToolbarSpacer(),
		runBoardAction,
		widget.NewToolbarSeparator(),
		otherThemeAction,
		settingsAction,
		helpAction,
	)
}
