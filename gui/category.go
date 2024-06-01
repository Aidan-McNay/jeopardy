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
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// addQuestionButton
//------------------------------------------------------------------------
// A button that adds a question to the given category

func addQuestionButton() fyne.CanvasObject {
	button := widget.NewButton("Add Question", func() {
		log.Println("Question Added")
	})
	button.Importance = widget.HighImportance
	return button
}

//------------------------------------------------------------------------
// Make a new Category element
//------------------------------------------------------------------------

func categoryGUI(category *logic.Category) fyne.CanvasObject {
	var rows []fyne.CanvasObject = nil

	rows = append(rows, widget.NewLabel(category.Name))
	for _, v := range category.Questions {
		displayText := fmt.Sprintf("%v", v.Points)
		rows = append(rows, widget.NewLabel(displayText))
	}
	rows = append(rows, addQuestionButton())
	rows = append(rows, layout.NewSpacer())
	return container.NewVBox(rows...)
}
