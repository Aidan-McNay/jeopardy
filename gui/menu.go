//========================================================================
// menu.go
//========================================================================
// Code for handling menu items
//
// Author: Aidan McNay
// Date: June 5th, 2024

package gui

import "fyne.io/fyne/v2"

//------------------------------------------------------------------------
// Define a menu item based on a keyCallback
//------------------------------------------------------------------------

func menuItemFromCallback(name string, callback keyCallback) *fyne.MenuItem {
	menuItem := fyne.NewMenuItem(
		name,
		callback.Callback,
	)
	menuItem.Shortcut = callback.CustomShortcut

	return menuItem
}

//------------------------------------------------------------------------
// Define our individual menu items
//------------------------------------------------------------------------

func newBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := newBoardShortcut(win)
	return menuItemFromCallback("New Board", callback)
}

func loadBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := loadBoardShortcut(win)
	return menuItemFromCallback("Open a Board", callback)
}

func saveBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := saveBoardShortcut(win)
	return menuItemFromCallback("Save a Board", callback)
}

//------------------------------------------------------------------------
// Define our "File" menu based on our menu items
//------------------------------------------------------------------------

func fileMenu(win fyne.Window) *fyne.Menu {
	items := [](*fyne.MenuItem){
		newBoardMenuItem(win),
		loadBoardMenuItem(win),
		saveBoardMenuItem(win),
	}
	return fyne.NewMenu(
		"File",
		items...,
	)
}

//------------------------------------------------------------------------
// Define our main menu
//------------------------------------------------------------------------

func MainMenu(win fyne.Window) *fyne.MainMenu {
	items := [](*fyne.Menu){
		fileMenu(win),
	}
	return fyne.NewMainMenu(items...)
}
