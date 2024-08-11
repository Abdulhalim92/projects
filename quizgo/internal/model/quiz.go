package model

type Question string
type AnswerOptions [4]string
type AnswerIndex int
type DifficultyLevel int

type Quiz struct {
	Question   Question
	Options    AnswerOptions
	Index      AnswerIndex
	Difficulty DifficultyLevel
}
