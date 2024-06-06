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
	"fyne.io/fyne/v2/theme"
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
	menuItem.Icon = theme.ContentAddIcon()
	return menuItem
}

func loadBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := loadBoardShortcut(win)
	menuItem := menuItemFromCallback("Open a Board", callback)
	menuItem.Icon = theme.FolderOpenIcon()
	return menuItem
}

func saveBoardMenuItem(win fyne.Window) *fyne.MenuItem {
	callback := saveBoardShortcut(win)
	menuItem := menuItemFromCallback("Save a Board", callback)
	menuItem.Icon = theme.DocumentSaveIcon()
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
