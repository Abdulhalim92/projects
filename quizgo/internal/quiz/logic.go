package quiz

import (
	"fmt"
	"quizgo/internal/model"
)

type DataStore interface {
	AddQuizData(quiz model.Quiz)
	CreateQuiz(q string, opts [4]string, ansIndex int) model.Quiz
	GetQuestionsLen() int
	QuestionForID(qID int) string
	AnswerForID(qID int) string
	OptionsForID(qID int) [4]string
}

type QuizLogic struct {
	ds DataStore
}

func (logic QuizLogic) GetQuiz(qID int) string {
	fmt.Printf("Quiz #%d \n", qID+1)
	q := logic.ds.QuestionForID(qID)
	opts := logic.ds.OptionsForID(qID)
	formattedOpts := logic.FormatOpts(opts)

	return fmt.Sprintf("%v\n%v", q, formattedOpts)
}

func (logic QuizLogic) FormatOpts(opts [4]string) string {
	firstRowGap := 20 - len(opts[0])
	secondRowGap := 20 - len(opts[2])
	var firstRowSpace, secondRowSpace string
	for i := 0; i < firstRowGap; i++ {
		firstRowSpace += " "
	}
	for i := 0; i < secondRowGap; i++ {
		secondRowSpace += " "
	}
	return fmt.Sprintf("     A) %v%vB) %v\n\n     C) %v%vD) %v",
		opts[0], firstRowSpace, opts[1],
		opts[2], secondRowSpace, opts[3])
}

func (logic QuizLogic) CheckAnswer(a string, qID int) bool {
	if a == logic.ds.AnswerForID(qID) {
		return true
	}

	return false
}

func (logic QuizLogic) AskQuestion(qID int) {
	var a string
	fmt.Println(logic.GetQuiz(qID))
	fmt.Scan(&a)
	for a != "a" && a != "b" && a != "c" && a != "d" {
		fmt.Println("Type an option letter: a, b, c, d")
		fmt.Scan(&a)
	}

	result := logic.CheckAnswer(a, qID)
	if result {
		fmt.Println("Right!")
	} else {
		fmt.Println("Wrong")
	}
}

func (logic QuizLogic) StartQuiz() {
	for i := 0; i < logic.ds.GetQuestionsLen(); i++ {
		logic.AskQuestion(i)
	}
}
