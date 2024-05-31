//========================================================================
// player.go
//========================================================================
// A representation of a Jeopardy player
//
// Author: Aidan McNay
// Date: May 30th, 2024

package logic

import "fmt"

//------------------------------------------------------------------------
// Define a Player Type
//------------------------------------------------------------------------

type Player struct {
	name  string
	score int
}

//------------------------------------------------------------------------
// Getters and Setters
//------------------------------------------------------------------------

func (p *Player) GetName() string {
	if p == nil {
		return ""
	}
	return p.name
}

func (p *Player) SetName(name string) {
	if p == nil {
		return
	}
	p.name = name
}

func (p *Player) ResetScore() {
	if p == nil {
		return
	}
	p.score = 0
}

func (p *Player) IncrScore(v int) {
	if p == nil {
		return
	}
	p.score += v
}

func (p *Player) GetScore() int {
	if p == nil {
		return 0
	}
	return p.score
}

//------------------------------------------------------------------------
// AsString
//------------------------------------------------------------------------
// A string representation of a player, for debugging

func (p *Player) AsString() string {
	if p == nil {
		return "Null Player"
	}
	return fmt.Sprintf("%v: %v", p.name, p.score)
}
