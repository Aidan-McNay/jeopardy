//========================================================================
// board.go
//========================================================================
// An interface for rendering a board in a GUI
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"image/color"
	"jeopardy/logic"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// addCategoryButton
//------------------------------------------------------------------------
// A button that adds a new category

func addCategoryButton() *fyne.Container {
	button := widget.NewButton("Add Category", func() {
		log.Println("Category Added")
	})
	button.Importance = widget.SuccessImportance
	return container.NewVBox(button, layout.NewSpacer())
}

//------------------------------------------------------------------------
// Make a new widget to represent a board
//------------------------------------------------------------------------

var titleColor color.Color = color.RGBA{41, 111, 246, 255}

func boardWidget() fyne.Widget {
	curr_board := logic.GetCurrBoard()

	var columns []fyne.CanvasObject = nil
	if curr_board != nil {
		for _, v := range curr_board.Categories {
			columns = append(columns, categoryGUI(v))
		}
	}
	columns = append(columns, addCategoryButton())
	gridLayout := container.NewHBox(columns...)

	var boardLayout fyne.CanvasObject
	if curr_board == nil {
		label := widget.NewLabel("No Current Board")
		label.Alignment = fyne.TextAlignCenter
		boardLayout = label
	} else {
		nameText := canvas.NewText(curr_board.Name, titleColor)
		nameText.Alignment = fyne.TextAlignCenter
		nameText.TextStyle = fyne.TextStyle{Bold: true}
		nameText.TextSize = 20
		boardLayout = container.NewVBox(nameText, gridLayout)
	}
	scrollWidget := container.NewScroll(boardLayout)
	return scrollWidget
}

//------------------------------------------------------------------------
// Make a new Board element (as a layout)
//------------------------------------------------------------------------

func BoardGUI() *fyne.Container {
	board := boardWidget()
	return container.NewStack(board)
}

//------------------------------------------------------------------------
// Allow for updating the board Widget
//------------------------------------------------------------------------

func UpdateBoard(board *fyne.Container) {
	board.RemoveAll()
	board.Add(boardWidget())
}
