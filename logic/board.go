//========================================================================
// board.go
//========================================================================
// A representation of a Jeopardy board
//
// Author: Aidan McNay
// Date: May 30th, 2024

package logic

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
