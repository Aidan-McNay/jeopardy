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
	"jeopardy/assets"
	"jeopardy/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// addSwapper
//------------------------------------------------------------------------
// Adds a Swapper, to swap the two adjacent categories

func addSwapper(idx1, idx2 int) fyne.CanvasObject {
	swapIcon := theme.NewThemedResource(assets.ResourceSwapPng)
	swapButton := widget.NewButtonWithIcon("", swapIcon, func() {
		curr_board := logic.GetCurrBoard()
		curr_board.SwapCategories(idx1, idx2)
		logic.BoardChange()
	})
	return container.NewVBox(swapButton, layout.NewSpacer())
}

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

	prompt := dialog.NewForm("New Category", "Add Category", "Cancel", items,
		onConfirm, win)

	var height float32 = prompt.MinSize().Height
	var width float32 = 400
	newSize := fyne.NewSize(width, height)
	prompt.Resize(newSize)

	prompt.Show()
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

func boardWidget(win fyne.Window) fyne.Widget {
	curr_board := logic.GetCurrBoard()

	var columns []fyne.CanvasObject = nil
	if curr_board != nil {
		for idx, v := range curr_board.Categories {
			if idx != 0 {
				columns = append(columns, addSwapper(idx, idx-1))
			}
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
		nameText := canvas.NewText(curr_board.Name, theme.PrimaryColor())
		nameText.Alignment = fyne.TextAlignCenter
		nameText.TextStyle = fyne.TextStyle{Bold: true}
		nameText.TextSize = 20

		nameBorder := canvas.NewRectangle(theme.BackgroundColor())
		nameBorder.StrokeWidth = 2
		nameBorder.StrokeColor = theme.PrimaryColor()

		name := container.NewStack(nameBorder, nameText)
		boardLayout = container.NewVBox(name, gridLayout)
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
