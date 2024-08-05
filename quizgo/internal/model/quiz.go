package model

type DifficultyLevel int

type Quiz struct {
	Question      string
	AnswerOptions [4]string
	AnswerIndex   int
	Difficulty    DifficultyLevel
}
