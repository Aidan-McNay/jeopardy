package main

import (
	"gui"
	"logic"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Jeopardy Editor")

	toolbar := gui.Toolbar(myWindow)

	label := widget.NewLabel("No Board Selected")
	logic.OnBoardChange(func(board *logic.Board) { label.SetText(board.Name) })

	content := container.NewBorder(toolbar, nil, nil, nil, label)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
