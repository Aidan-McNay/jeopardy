//========================================================================
// question.go
//========================================================================
// An interface for rendering a question in a GUI
//
// Author: Aidan McNay
// Date: June 1st, 2024

package gui

import (
	"fmt"
	"strconv"

	"jeopardy/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// deleteQuestion
//------------------------------------------------------------------------
// Creates a dialogue to confirm deletion of a question

func deleteQuestion(question *logic.Question,
	category *logic.Category,
	form *dialog.FormDialog,
	win fyne.Window,
) {
	deleteCallback := func(b bool) {
		if b {
			category.RemoveQuestion(question)
			form.Hide()
			logic.BoardChange()
		}
	}
	dialog.ShowConfirm(
		"Delete Question",
		"Are you sure? This action can't be undone",
		deleteCallback,
		win,
	)
}

//------------------------------------------------------------------------
// editQuestion
//------------------------------------------------------------------------
// Creates a dialogue to edit the question

func editQuestion(win fyne.Window,
	category *logic.Category,
	question *logic.Question,
) {
	newPrompt := widget.NewMultiLineEntry()
	newPrompt.Validator = validation.NewRegexp(`^.+$`, "Prompt must be non-empty")
	newPrompt.Text = question.Prompt

	newAnswer := widget.NewMultiLineEntry()
	newAnswer.Validator = validation.NewRegexp(`^.+$`, "Answer must be non-empty")
	newAnswer.Text = question.Answer

	newPoints := widget.NewEntry()
	newPoints.Validator = isInt
	newPoints.Text = fmt.Sprintf("%v", question.Points)

	deleteButton := widget.NewButtonWithIcon("", theme.CancelIcon(),
		func() {})
	deleteButton.Importance = widget.DangerImportance

	items := []*widget.FormItem{
		widget.NewFormItem("Prompt", newPrompt),
		widget.NewFormItem("Answer", newAnswer),
		widget.NewFormItem("Points", newPoints),
		widget.NewFormItem("Delete Question?", deleteButton),
	}
	onConfirm := func(b bool) {
		if !b {
			return
		}
		question.Prompt = newPrompt.Text
		question.Answer = newAnswer.Text
		question.Points, _ = strconv.Atoi(newPoints.Text)
		logic.BoardChange()
	}

	formTitle := "Edit Question"
	prompt := dialog.NewForm(formTitle, "Save", "Cancel", items,
		onConfirm, win)
	deleteButton.OnTapped = func() {
		deleteQuestion(question, category, prompt, win)
	}

	var height float32 = prompt.MinSize().Height
	var width float32 = 400
	newSize := fyne.NewSize(width, height)
	prompt.Resize(newSize)

	prompt.Show()
}

//------------------------------------------------------------------------
// questionButton
//------------------------------------------------------------------------
// Creates the button to edit a question

func questionButton(win fyne.Window,
	category *logic.Category,
	question *logic.Question,
) fyne.CanvasObject {
	displayText := fmt.Sprintf("%v", question.Points)
	button := widget.NewButton(displayText, func() {
		editQuestion(win, category, question)
	})
	button.Importance = widget.LowImportance
	return button
}
