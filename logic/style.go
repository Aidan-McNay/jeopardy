//========================================================================
// style.go
//========================================================================
// Functions for determining the style of the game
//
// Author: Aidan McNay
// Date: June 8th, 2024

package logic

import (
	"image/color"
	"io"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

//------------------------------------------------------------------------
// Define a Style Type
//------------------------------------------------------------------------

type Style struct {
	UseColor  bool
	Color     color.Color
	Image     *fyne.StaticResource
	TextColor color.Color
}

//------------------------------------------------------------------------
// Define an allocator for a base style, defaulting to using color
//------------------------------------------------------------------------

func NewStyle() *Style {
	return &Style{
		UseColor:  true,
		Color:     color.Transparent,
		Image:     nil,
		TextColor: color.Black,
	}
}

//------------------------------------------------------------------------
// Helper Functions
//------------------------------------------------------------------------

func ResourceFromURI(uri fyne.URI) *fyne.StaticResource {
	reader, _ := storage.Reader(uri)
	defer reader.Close()

	data, _ := io.ReadAll(reader)
	return &fyne.StaticResource{
		StaticName:    filepath.Base(uri.String()),
		StaticContent: data,
	}
}

//------------------------------------------------------------------------
// Define a style for the overall game
//------------------------------------------------------------------------

type GameStyle struct {
	CategoryStyle *Style
	QuestionStyle *Style
}

//------------------------------------------------------------------------
// Define an allocator that initializes the style
//------------------------------------------------------------------------

func NewGameStyle() *GameStyle {
	return &GameStyle{
		CategoryStyle: NewStyle(),
		QuestionStyle: NewStyle(),
	}
}
