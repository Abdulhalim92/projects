package quiz

import "fmt"


type DataStore interface {
	GetQuestion(difLevel DifficultyLevel, qIndex int) (FormattedQuiz, error)
	GetQuestionsByDifficulty(difLevel DifficultyLevel) []FormattedQuiz
}

type QuizLogic struct {
	ds DataStore
}

func (ql QuizLogic) AskQuestion(fq FormattedQuiz) bool {
	var userAnswer string
	q := fq.question
	a := fq.answer

	fmt.Println(q)
	fmt.Scan(&userAnswer)

	for userAnswer != "a" && userAnswer != "b" && userAnswer != "c" && userAnswer != "d" {
		fmt.Println("Type an option letter: a, b, c, d")
		fmt.Scan(&userAnswer)
	}

	if userAnswer == a {
		fmt.Println("Right!")
		return true
	} else {
		fmt.Println("Wrong")
		return false
	}
}
