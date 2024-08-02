package quiz

import (
	"fmt"
	"quizgo/internal/model"
)

type DataStore interface {
	AddQuizData(quiz model.Quiz)
	CreateQuiz(q string, opts [4]string, ansIndex int) model.Quiz
	GetQuestionsLen() int
	AnswerForID(qID int) string
	GetFormattedQuiz(qID int) string
}

type QuizLogic struct {
	ds DataStore
}

func (ql QuizLogic) CheckAnswer(a string, qID int) bool {
	if a == ql.ds.AnswerForID(qID) {
		return true
	}

	return false
}

func (ql QuizLogic) AskQuestion(qID int) {
	var a string
	fmt.Println(ql.ds.GetFormattedQuiz(qID))
	fmt.Scan(&a)
	for a != "a" && a != "b" && a != "c" && a != "d" {
		fmt.Println("Type an option letter: a, b, c, d")
		fmt.Scan(&a)
	}

	result := ql.CheckAnswer(a, qID)
	if result {
		fmt.Println("Right!")
	} else {
		fmt.Println("Wrong")
	}
}

func (ql QuizLogic) StartQuiz() {
	for i := 0; i < ql.ds.GetQuestionsLen(); i++ {
		ql.AskQuestion(i)
	}
}
