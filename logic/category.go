//========================================================================
// category.go
//========================================================================
// A representation of a single category in Jeopardy
//
// Author: Aidan McNay
// Date: May 30th, 2024

package logic

import "slices"

//------------------------------------------------------------------------
// Define a Category Type
//------------------------------------------------------------------------

type Category struct {
	name      string
	questions [](*Question)
}

//------------------------------------------------------------------------
// Provide an allocator for a category
//------------------------------------------------------------------------

func MakeCategory(name string) *Category {
	return &Category{name, nil}
}

//------------------------------------------------------------------------
// AddQuestions
//------------------------------------------------------------------------
// Insert new question(s), such that the slice of questions remains sorted
// by points
//
// Inspired by https://stackoverflow.com/a/55460931/23068975

func cmp(a, b *Question) int {
	return a.points - b.points
}

func (c *Category) AddQuestion(questions ...*Question) {
	if c == nil {
		return
	}
	c.questions = append(c.questions, questions...)
	slices.SortFunc(c.questions, cmp)
}

//------------------------------------------------------------------------
// MaxPoints
//------------------------------------------------------------------------
// Returns the maximum points of any question in the category
//
// Returns 0 if the category contains no questions

func (c *Category) MaxPoints() int {
	if c == nil {
		return 0
	}
	if len(c.questions) == 0 {
		return 0
	} else {
		curr_value := c.questions[0].points
		for _, v := range c.questions {
			if points := v.points; points > curr_value {
				curr_value = points
			}
		}
		return curr_value
	}
}
