package main

import (
	"jeopardy/assets"
	"jeopardy/gui"
	"jeopardy/logic"
	"jeopardy/style"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.NewWithID("github.com.Aidan-McNay.jeopardy")
	myApp.SetIcon(assets.ResourceLogoPng)

	myWindow := myApp.NewWindow("Jeopardy Editor")
	style.InitTheme(myApp)

	toolbar := gui.Toolbar(myWindow)
	boardEditor := gui.BoardGUI(myWindow)

	content := container.NewBorder(toolbar, nil, nil, nil, boardEditor)
	logic.OnBoardChange(func(board *logic.Board) {
		gui.UpdateBoard(boardEditor, myWindow)
		content.Refresh()
	})
	myWindow.SetContent(content)

	gui.AddTopLevelShortcuts(myWindow)
	myWindow.SetMainMenu(gui.MainMenu(myWindow))

	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.SetMaster()
	myWindow.ShowAndRun()
}
