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
	Name      string
	Questions [](*Question)
}

//------------------------------------------------------------------------
// Provide an allocator for a category
//------------------------------------------------------------------------

func MakeCategory(name string) *Category {
	return &Category{name, nil}
}

//------------------------------------------------------------------------
// Derived Attributes
//------------------------------------------------------------------------

func (c *Category) Height() int {
	if c == nil {
		return 0
	}
	return 1 + len(c.Questions)
}

//------------------------------------------------------------------------
// AddQuestions
//------------------------------------------------------------------------
// Insert new question(s), such that the slice of questions remains sorted
// by points

func cmp(a, b *Question) int {
	return a.Points - b.Points
}

func (c *Category) AddQuestions(questions ...*Question) {
	if c == nil {
		return
	}
	c.Questions = append(c.Questions, questions...)
	slices.SortFunc(c.Questions, cmp)
}

//------------------------------------------------------------------------
// RemoveQuestion
//------------------------------------------------------------------------
// Removes the given question by pointer

func (c *Category) RemoveQuestion(question *Question) {
	var newQuestions [](*Question) = nil
	for _, v := range c.Questions {
		if v != question {
			newQuestions = append(newQuestions, v)
		}
	}
	c.Questions = newQuestions
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
	if len(c.Questions) == 0 {
		return 0
	} else {
		curr_value := c.Questions[0].Points
		for _, v := range c.Questions {
			if points := v.Points; points > curr_value {
				curr_value = points
			}
		}
		return curr_value
	}
}
