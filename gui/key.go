//========================================================================
// key.go
//========================================================================
// Code for handling key-triggered callbacks
//
// Author: Aidan McNay
// Date: June 5th, 2024

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

//------------------------------------------------------------------------
// Define the type for a key callback
//------------------------------------------------------------------------
// This should identify a key, and a function to call when the key is
// pressed

type keyCallback struct {
	CustomShortcut *desktop.CustomShortcut
	Callback       func()
}

//------------------------------------------------------------------------
// Implement a constructor for a callback
//------------------------------------------------------------------------

func NewCallback(key fyne.KeyName, mod fyne.KeyModifier, callback func()) keyCallback {
	return keyCallback{
		CustomShortcut: &desktop.CustomShortcut{
			KeyName:  key,
			Modifier: mod,
		},
		Callback: callback,
	}
}

//------------------------------------------------------------------------
// Implement member functions of the callback
//------------------------------------------------------------------------

func (c keyCallback) shortcut() *desktop.CustomShortcut {
	return c.CustomShortcut
}

func (c keyCallback) trigger(_ fyne.Shortcut) {
	c.Callback()
}

func (c keyCallback) addToWindow(win fyne.Window) {
	win.Canvas().AddShortcut(c.shortcut(), c.trigger)
}

//------------------------------------------------------------------------
// Define shortcuts we'd like to have for the top-level canvas
//------------------------------------------------------------------------

func newBoardShortcut(win fyne.Window) keyCallback {
	return NewCallback(
		fyne.KeyN,
		fyne.KeyModifierShortcutDefault,
		func() {
			promptNewBoard(win)
		},
	)
}

func loadBoardShortcut(win fyne.Window) keyCallback {
	return NewCallback(
		fyne.KeyO,
		fyne.KeyModifierShortcutDefault,
		func() {
			loadFromFile(win)
		},
	)
}

func saveBoardShortcut(win fyne.Window) keyCallback {
	return NewCallback(
		fyne.KeyS,
		fyne.KeyModifierShortcutDefault,
		func() {
			saveToFile(win)
		},
	)
}

//------------------------------------------------------------------------
// Add the shortcuts to the top-level canvas
//------------------------------------------------------------------------

func AddTopLevelShortcuts(win fyne.Window) {
	newBoardShortcut(win).addToWindow(win)
	loadBoardShortcut(win).addToWindow(win)
	saveBoardShortcut(win).addToWindow(win)
}
