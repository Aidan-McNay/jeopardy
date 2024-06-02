package main

import (
	"jeopardy/gui"
	"jeopardy/logic"
	"jeopardy/style"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.NewWithID("github.com.Aidan-McNay.jeopardy")
	myWindow := myApp.NewWindow("Jeopardy Editor")
	style.InitTheme(myApp)

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
