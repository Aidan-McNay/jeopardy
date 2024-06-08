//========================================================================
// style.go
//========================================================================
// A GUI for changing the style of a board
//
// Author: Aidan McNay
// Date: June 8th, 2024

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

//------------------------------------------------------------------------
// Main style GUI
//------------------------------------------------------------------------

func styleGUI(win fyne.Window) {
	openPopup()

	dialog.ShowConfirm("Style Editor", "Placeholder", func(b bool) {
		closePopup()
	}, win)
}
