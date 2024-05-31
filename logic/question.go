//========================================================================
// question.go
//========================================================================
// A representation of a single question in Jeopardy
//
// Author: Aidan McNay
// Date: May 30th, 2024

package logic

//------------------------------------------------------------------------
// Define a Question Type
//------------------------------------------------------------------------

type Question struct {
	Prompt, Answer string
	Points         int
	Answered       bool
}

//------------------------------------------------------------------------
// Provide an allocator for a question
//------------------------------------------------------------------------

func MakeQuestion(prompt, answer string, points int) *Question {
	return &Question{prompt, answer, points, false}
}

//------------------------------------------------------------------------
// Getters and Setters
//------------------------------------------------------------------------

func (q *Question) GetPrompt() string {
	if q == nil {
		return ""
	}
	return q.Prompt
}

func (q *Question) GetAnswer() string {
	if q == nil {
		return ""
	}
	return q.Answer
}

func (q *Question) GetPoints() int {
	if q == nil {
		return 0
	}
	return q.Points
}

func (q *Question) SetAnswered() {
	if q != nil {
		q.Answered = true
	}
}
