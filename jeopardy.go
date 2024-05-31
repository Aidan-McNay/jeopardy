package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func toolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			log.Println("Add clicked")
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			log.Println("Open clicked")
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			log.Println("Save clicked")
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

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Jeopardy Editor")

	content := container.NewBorder(toolbar(), nil, nil, nil, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
