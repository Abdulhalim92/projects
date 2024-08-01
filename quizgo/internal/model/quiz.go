package model

type Quiz struct {
	QuestionID    int
	Question      string
	AnswerOptions [4]string
	AnswerIndex   int
}
