package quiz

import (
	"fmt"
	"strings"
)

type DataStore interface {
	GetQuestion(id ID) (*FormattedQuiz, error)
	GetQuestionsByDifficulty(difLevel Dif) []FormattedQuiz
}

type QuizLogic struct {
	ds DataStore
}

func (ql QuizLogic) AskQuestion(fq FormattedQuiz) bool {
	var userAnswer string
	q := fq.Question
	opts := fq.Options
	a := fq.Answer

	fmt.Printf("%s\n\n%s\n\n", q, opts)
	fmt.Print(">>> ")
	fmt.Scan(&userAnswer)

	userAnswer = strings.ToLower(userAnswer)
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
