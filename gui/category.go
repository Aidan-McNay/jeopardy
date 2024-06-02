//========================================================================
// category.go
//========================================================================
// An interface for rendering a category in a GUI
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"errors"
	"fmt"
	"jeopardy/logic"
	"strconv"

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
// otherCategoryExists
//------------------------------------------------------------------------
// Checks whether a category name already exists, if it's not our original
// name

func otherCategoryExists(origName string) func(name string) error {
	return func(name string) error {
		board := logic.GetCurrBoard()
		if board == nil {
			return nil
		}
		for _, v := range board.Categories {
			if (name == v.Name) && (name != origName) {
				errorText := fmt.Sprintf("%v already exists", name)
				return errors.New(errorText)
			}
		}
		return nil
	}
}

//------------------------------------------------------------------------
// editCategory
//------------------------------------------------------------------------
// Creates a dialogue to edit the category

func editCategory(win fyne.Window, category *logic.Category) {
	newName := widget.NewEntry()
	newName.SetText(category.Name)
	newName.Validator = validation.NewAllStrings(
		validation.NewRegexp(`^.+$`, "Category must have a non-empty name"),
		otherCategoryExists(category.Name),
	)

	items := []*widget.FormItem{
		widget.NewFormItem("Category Name", newName),
	}
	onConfirm := func(b bool) {
		if !b {
			return
		}
		category.Name = newName.Text
		logic.BoardChange()
	}
	prompt := dialog.NewForm("Edit Category", "Save", "Cancel", items,
		onConfirm, win)

	var height float32 = prompt.MinSize().Height
	var width float32 = 400
	newSize := fyne.NewSize(width, height)
	prompt.Resize(newSize)

	prompt.Show()
}

//------------------------------------------------------------------------
// categoryButton
//------------------------------------------------------------------------
// Creates the button to edit a category

func categoryButton(win fyne.Window, category *logic.Category) fyne.CanvasObject {
	name := widget.NewButton(category.Name, func() {
		editCategory(win, category)
	})
	name.Importance = widget.LowImportance

	categoryBorder := canvas.NewRectangle(theme.BackgroundColor())
	categoryBorder.StrokeWidth = 2
	categoryBorder.StrokeColor = theme.PrimaryColor()
	return container.NewStack(categoryBorder, name)
}

//------------------------------------------------------------------------
// isInt
//------------------------------------------------------------------------
// Determines whether a string can be represented as an integer

func isInt(s string) error {
	_, err := strconv.Atoi(s)
	if err != nil {

		return errors.New(s + " is not a valid number")
	}
	return nil
}

//------------------------------------------------------------------------
// addQuestion
//------------------------------------------------------------------------
// Creates a dialogue to add a new question

func addQuestion(win fyne.Window, category *logic.Category) {
	newPrompt := widget.NewMultiLineEntry()
	newPrompt.Validator = validation.NewRegexp(`^.+$`, "Prompt must be non-empty")

	newAnswer := widget.NewMultiLineEntry()
	newAnswer.Validator = validation.NewRegexp(`^.+$`, "Answer must be non-empty")

	newPoints := widget.NewEntry()
	newPoints.Validator = isInt

	items := []*widget.FormItem{
		widget.NewFormItem("Prompt", newPrompt),
		widget.NewFormItem("Answer", newAnswer),
		widget.NewFormItem("Points", newPoints),
	}
	onConfirm := func(b bool) {
		if !b {
			return
		}
		prompt := newPrompt.Text
		answer := newAnswer.Text
		points, _ := strconv.Atoi(newPoints.Text)
		newQuestion := logic.MakeQuestion(prompt, answer, points)
		category.AddQuestions(newQuestion)
		logic.BoardChange()
	}

	formTitle := fmt.Sprintf("New Question for %v", category.Name)
	prompt := dialog.NewForm(formTitle, "Add Question", "Cancel", items,
		onConfirm, win)

	var height float32 = prompt.MinSize().Height
	var width float32 = 400
	newSize := fyne.NewSize(width, height)
	prompt.Resize(newSize)

	prompt.Show()
}

//------------------------------------------------------------------------
// addQuestionButton
//------------------------------------------------------------------------
// A button that adds a question to the given category

func addQuestionButton(win fyne.Window, category *logic.Category) fyne.CanvasObject {
	button := widget.NewButton("Add Question", func() {
		addQuestion(win, category)
	})
	button.Importance = widget.HighImportance
	return button
}

//------------------------------------------------------------------------
// Make a new Category element
//------------------------------------------------------------------------

func categoryGUI(win fyne.Window, category *logic.Category) fyne.CanvasObject {
	var rows []fyne.CanvasObject = nil

	rows = append(rows, categoryButton(win, category))
	for _, v := range category.Questions {
		rows = append(rows, questionButton(win, v))
	}
	rows = append(rows, addQuestionButton(win, category))
	rows = append(rows, layout.NewSpacer())
	return container.NewVBox(rows...)
}
