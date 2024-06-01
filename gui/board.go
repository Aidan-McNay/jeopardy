//========================================================================
// board.go
//========================================================================
// An interface for rendering a board in a GUI
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"jeopardy/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// Current Length of the Table
//------------------------------------------------------------------------

func rowCount() int {
	return logic.GetCurrBoard().Height() + 1
}

func colCount() int {
	return logic.GetCurrBoard().Width() + 1
}

//------------------------------------------------------------------------
// Template Cell
//------------------------------------------------------------------------

func templateCell() fyne.CanvasObject {
	return widget.NewLabel("Jeopardy Cell")
}

//------------------------------------------------------------------------
// Update Cells
//------------------------------------------------------------------------

func updateCell(i widget.TableCellID, o fyne.CanvasObject) {
	curr_board := logic.GetCurrBoard()
	if i.Col < curr_board.Width() {
		updateCategoryCell(curr_board.Categories[i.Col], i, o)
	} else if i.Row == 0 {
		o.(*widget.Label).SetText("Add New Category")
	} else {
		o.(*widget.Label).SetText("")
	}
}

//------------------------------------------------------------------------
// Make a New Table
//------------------------------------------------------------------------

func BoardGUI() *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return rowCount(), colCount()
		},
		templateCell,
		updateCell,
	)
}
