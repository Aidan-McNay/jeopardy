//========================================================================
// board.go
//========================================================================
// A representation of a Jeopardy board
//
// Author: Aidan McNay
// Date: May 30th, 2024

package logic

import "reflect"

//------------------------------------------------------------------------
// Define a Board Type
//------------------------------------------------------------------------

type Board struct {
	Name       string
	Categories [](*Category)
	Players    [](*Player)
}

//------------------------------------------------------------------------
// Provide an allocator for a board
//------------------------------------------------------------------------

func MakeBoard(name string) *Board {
	return &Board{name, nil, nil}
}

//------------------------------------------------------------------------
// Derived Attributes
//------------------------------------------------------------------------

func (b *Board) Width() int {
	if b == nil {
		return 0
	}
	return len(b.Categories)
}

func (b *Board) Height() int {
	if b == nil {
		return 0
	}
	if len(b.Categories) == 0 {
		return 0
	}
	curr_height := b.Categories[0].Height()
	for _, v := range b.Categories {
		if h := v.Height(); h > curr_height {
			curr_height = h
		}
	}
	return curr_height
}

//------------------------------------------------------------------------
// AddCategories
//------------------------------------------------------------------------
// Appends a new category(s)

func (b *Board) AddCategories(categories ...*Category) {
	if b == nil {
		return
	}
	b.Categories = append(b.Categories, categories...)
}

//------------------------------------------------------------------------
// SwapCategories
//------------------------------------------------------------------------
// Swaps the categories at the given indeces

func (b *Board) SwapCategories(idx1, idx2 int) {
	categorySwapper := reflect.Swapper(b.Categories)
	categorySwapper(idx1, idx2)
}

//------------------------------------------------------------------------
// RemoveCategory
//------------------------------------------------------------------------
// Removes the given category by pointer

func (b *Board) RemoveCategory(category *Category) {
	var newCategories [](*Category) = nil
	for _, v := range b.Categories {
		if v != category {
			newCategories = append(newCategories, v)
		}
	}
	b.Categories = newCategories
}

//------------------------------------------------------------------------
// AddPlayers
//------------------------------------------------------------------------
// Appends a new player(s)

func (b *Board) AddPlayers(players ...*Player) {
	if b == nil {
		return
	}
	b.Players = append(b.Players, players...)
}

//------------------------------------------------------------------------
// MaxPoints
//------------------------------------------------------------------------
// Returns the maximum points of any category on the board
//
// Returns 0 if the board contains no categories or if passed a null
// pointer

func (b *Board) MaxPoints() int {
	if b == nil {
		return 0
	}
	if len(b.Categories) == 0 {
		return 0
	} else {
		curr_value := b.Categories[0].MaxPoints()
		for _, v := range b.Categories {
			if points := v.MaxPoints(); points > curr_value {
				curr_value = points
			}
		}
		return curr_value
	}
}
