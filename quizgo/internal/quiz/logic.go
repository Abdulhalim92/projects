package quiz

import (
	"fmt"
)

type DataStore interface {
	GetQuestion(difLevel DifficultyLevel, qIndex int) (FormattedQuiz, error)

	GetQuestionsByDifficulty(difLevel DifficultyLevel) []FormattedQuiz
}

type QuizLogic struct {
	ds DataStore
}

func (ql QuizLogic) AskQuestion(fq FormattedQuiz) {
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
	} else {
		fmt.Println("Wrong")
	}
}

func (ql QuizLogic) StartQuiz() {
	var difLevel DifficultyLevel = 1
	var quizNum int = 1
	difficulties := []string{"Easy", "Medium", "Hard"}
	for ; difLevel <= 3; difLevel++ {
		for _, v := range ql.ds.GetQuestionsByDifficulty(difLevel) {

			fmt.Printf("\nQuiz #%d          Difficulty: %s\n", quizNum, difficulties[difLevel-1])
			ql.AskQuestion(v)
			quizNum++
		}

	}
}
