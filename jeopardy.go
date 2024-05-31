package main

import (
	"log"

	"logic"
	"storage"

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
	board := logic.MakeBoard("Test Board")

	question1 := logic.MakeQuestion("prompt1", "answer1", 200)
	question2 := logic.MakeQuestion("prompt2", "answer2", 100)
	question3 := logic.MakeQuestion("prompt3", "answer3", 300)
	question4 := logic.MakeQuestion("prompt4", "answer4", 400)

	category1 := logic.MakeCategory("Category1")
	category2 := logic.MakeCategory("Category2")

	category1.AddQuestions(question1, question2)
	category2.AddQuestions(question3, question4)

	board.AddCategories(category1, category2)

	storage.Save("board.txt", board)

	//------------------

	myApp := app.New()
	myWindow := myApp.NewWindow("Jeopardy Editor")

	content := container.NewBorder(toolbar(), nil, nil, nil, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
