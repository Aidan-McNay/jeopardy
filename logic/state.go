//========================================================================
// state.go
//========================================================================
// A wrapper for managing the current board state
//
// Author: Aidan McNay
// Date: May 31st, 2024

package logic

import (
	"file"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

//------------------------------------------------------------------------
// Current Board
//------------------------------------------------------------------------

var currBoard *Board

//------------------------------------------------------------------------
// Callbacks to use when the board is changed
//------------------------------------------------------------------------

var callbacks []func(board *Board)

func OnBoardChange(callback func(board *Board)) {
	callbacks = append(callbacks, callback)
}

func BoardChange() {
	for _, c := range callbacks {
		c(currBoard)
	}
}

//------------------------------------------------------------------------
// Getters and Setters
//------------------------------------------------------------------------

func GetCurrBoard() *Board {
	return currBoard
}

func SetCurrBoard(new_board *Board) {
	currBoard = new_board
	BoardChange()
}

//------------------------------------------------------------------------
// Changes a URIWriteCloser to use the correct extension
//------------------------------------------------------------------------

func getCorrectExtension(fileWriter fyne.URIWriteCloser) fyne.URIWriteCloser {
	uri := fileWriter.URI()
	if uri.Extension() == ".jpdy" {
		return fileWriter
	}

	// Need to re-get the correct URI, closing the current writer
	fileWriter.Close()
	storage.Delete(uri)

	uri_string := uri.String() + ".jpdy"
	uri, err := storage.ParseURI(uri_string)
	if err != nil {
		log.Fatal(err)
	}

	f, err := storage.Writer(uri)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

//------------------------------------------------------------------------
// Loading and Saving the Board
//------------------------------------------------------------------------

func LoadCurrBoard(fileReader fyne.URIReadCloser) {
	file.Load(fileReader, currBoard)
	BoardChange()
}

func SaveCurrBoard(fileWriter fyne.URIWriteCloser) {
	extensionWriter := getCorrectExtension(fileWriter)
	file.Save(extensionWriter, currBoard)
}

//------------------------------------------------------------------------
// Starting a New Board
//------------------------------------------------------------------------

func NewBoard(name string) {
	SetCurrBoard(MakeBoard(name))
}
