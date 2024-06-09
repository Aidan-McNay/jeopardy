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

func styleMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := styleShortcut(win)
	menuItem := menuItemFromCallback("Edit Style", callback)
	return menuItem
}

//------------------------------------------------------------------------
// Define our "Board" menu based on our menu items
//------------------------------------------------------------------------

func boardMenu(win fyne.Window) *fyne.Menu {
	items := [](*fyne.MenuItem){
		newBoardMenuItem(win),
		loadBoardMenuItem(win),
		saveBoardMenuItem(win),
		saveAsBoardMenuItem(win),
		fyne.NewMenuItemSeparator(),
		styleMenuItem(win),
	}
	return fyne.NewMenu(
		"Board",
		items...,
	)
}

//------------------------------------------------------------------------
// Define our main menu
//------------------------------------------------------------------------

func MainMenu(win fyne.Window) *fyne.MainMenu {
	items := [](*fyne.Menu){
		boardMenu(win),
	}
	return fyne.NewMainMenu(items...)
}
