package main

import (
	"jeopardy/gui"
	"jeopardy/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Jeopardy Editor")

	toolbar := gui.Toolbar(myWindow)
	grid := gui.BoardGUI(myWindow)

	content := container.NewBorder(toolbar, nil, nil, nil, grid)
	logic.OnBoardChange(func(board *logic.Board) {
		gui.UpdateBoard(grid, myWindow)
		content.Refresh()
	})
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}
