//========================================================================
// menu.go
//========================================================================
// Code for handling menu items
//
// Author: Aidan McNay
// Date: June 5th, 2024

package gui

import (
	"fyne.io/fyne/v2"
)

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
	menuItem := menuItemFromCallback("New Board", callback)
	return menuItem
}

func loadBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := loadBoardShortcut(win)
	menuItem := menuItemFromCallback("Open Board", callback)
	return menuItem
}

func saveBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := saveBoardShortcut(win)
	menuItem := menuItemFromCallback("Save", callback)
	return menuItem
}

func saveAsBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := saveAsBoardShortcut(win)
	menuItem := menuItemFromCallback("Save As...", callback)
	return menuItem
}

//------------------------------------------------------------------------
// Define our "File" menu based on our menu items
//------------------------------------------------------------------------

func fileMenu(win fyne.Window) *fyne.Menu {
	items := [](*fyne.MenuItem){
		newBoardMenuItem(win),
		loadBoardMenuItem(win),
		saveBoardMenuItem(win),
		saveAsBoardMenuItem(win),
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
