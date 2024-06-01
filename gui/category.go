//========================================================================
// category.go
//========================================================================
// An interface for rendering a category in a GUI
//
// Author: Aidan McNay
// Date: May 31st, 2024

package gui

import (
	"fmt"
	"jeopardy/logic"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// isInt
//------------------------------------------------------------------------
// Determines whether a string can be represented as an integer

func isInt(s string) error {
	_, err := strconv.Atoi(s)
	return err
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
	dialog.ShowForm(formTitle, "Add Question", "Cancel", items,
		onConfirm, win)
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

	rows = append(rows, widget.NewLabel(category.Name))
	for _, v := range category.Questions {
		displayText := fmt.Sprintf("%v", v.Points)
		rows = append(rows, widget.NewLabel(displayText))
	}
	rows = append(rows, addQuestionButton(win, category))
	rows = append(rows, layout.NewSpacer())
	return container.NewVBox(rows...)
}
