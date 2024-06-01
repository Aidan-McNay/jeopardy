//========================================================================
// toolbar.go
//========================================================================
// A toolbar, for providing icon-based actions for the user
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"log"
	"logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// New Board Creation
//------------------------------------------------------------------------

func promptNewBoard(win fyne.Window) {
	newName := widget.NewEntry()
	items := []*widget.FormItem{
		widget.NewFormItem("Board Name", newName),
	}
	onConfirm := func(b bool) {
		if !b {
			return
		}
		logic.NewBoard(newName.Text)
	}
	dialog.ShowForm("New Board", "Create New Board", "Cancel", items,
		onConfirm, win)
}

//------------------------------------------------------------------------
// File Manipulation
//------------------------------------------------------------------------

func loadFromFile(win fyne.Window) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if reader == nil {
			log.Println("Cancelled")
			return
		}

		logic.LoadCurrBoard(reader)
	}, win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".jpdy"}))
	fd.Show()
}

func saveToFile(win fyne.Window) {
	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if writer == nil {
			log.Println("Cancelled")
			return
		}

		logic.SaveCurrBoard(writer)
	}, win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".jpdy"}))
	fd.Show()
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
			log.Println("Settings clicked")
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Help clicked")
		}),
	)
}
