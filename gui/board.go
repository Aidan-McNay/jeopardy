//========================================================================
// board.go
//========================================================================
// An interface for rendering a board in a GUI
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"errors"
	"fmt"
	"image/color"
	"jeopardy/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// categoryExists
//------------------------------------------------------------------------
// Checks whether a category name already exists

func categoryExists(name string) error {
	board := logic.GetCurrBoard()
	if board == nil {
		return nil
	}
	for _, v := range board.Categories {
		if name == v.Name {
			errorText := fmt.Sprintf("%v already exists", name)
			return errors.New(errorText)
		}
	}
	return nil
}

//------------------------------------------------------------------------
// addCategory
//------------------------------------------------------------------------
// Creates a dialogue to add a new category

func addCategory(win fyne.Window) {
	newName := widget.NewEntry()
	newName.Validator = validation.NewAllStrings(
		validation.NewRegexp(`^.+$`, "Category must have a non-empty name"),
		categoryExists,
	)
	items := []*widget.FormItem{
		widget.NewFormItem("Category Name", newName),
	}
	onConfirm := func(b bool) {
		if !b {
			return
		}
		board := logic.GetCurrBoard()
		board.Categories = append(board.Categories, logic.MakeCategory(newName.Text))
		logic.BoardChange()
	}
	dialog.ShowForm("New Category", "Add Category", "Cancel", items,
		onConfirm, win)
}

//------------------------------------------------------------------------
// addCategoryButton
//------------------------------------------------------------------------
// A button that adds a new category

func addCategoryButton(win fyne.Window) *fyne.Container {
	button := widget.NewButton("Add Category", func() {
		addCategory(win)
	})
	button.Importance = widget.SuccessImportance
	return container.NewVBox(button, layout.NewSpacer())
}

//------------------------------------------------------------------------
// Make a new widget to represent a board
//------------------------------------------------------------------------

var titleColor color.Color = color.RGBA{41, 111, 246, 255}

func boardWidget(win fyne.Window) fyne.Widget {
	curr_board := logic.GetCurrBoard()

	var columns []fyne.CanvasObject = nil
	if curr_board != nil {
		for _, v := range curr_board.Categories {
			columns = append(columns, categoryGUI(win, v))
		}
	}
	columns = append(columns, addCategoryButton(win))
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

func BoardGUI(win fyne.Window) *fyne.Container {
	board := boardWidget(win)
	return container.NewStack(board)
}

//------------------------------------------------------------------------
// Allow for updating the board Widget
//------------------------------------------------------------------------

func UpdateBoard(board *fyne.Container, win fyne.Window) {
	board.RemoveAll()
	board.Add(boardWidget(win))
}
