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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// Update a Cell
//------------------------------------------------------------------------
// This assumes that we are in the correct column

func updateCategoryCell(category *logic.Category, i widget.TableCellID, o fyne.CanvasObject) {
	if i.Row == 0 {
		o.(*widget.Label).SetText(category.Name)
	} else if q_idx := i.Row - 1; q_idx < len(category.Questions) {
		string_of_points := fmt.Sprintf("%v", category.Questions[q_idx].Points)
		o.(*widget.Label).SetText(string_of_points)
	}
}
