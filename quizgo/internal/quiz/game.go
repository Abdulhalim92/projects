package quiz

import (
	"fmt"
)

type Logic interface {
	AskQuestion(fq FormattedQuiz) bool
}

type Game struct {
	l        Logic
	ds       DataStore
	attempts int
}

func (g Game) DrawHearts(hAmount int) {
	rows := [5]string{
		"  ###   ###  ",
		" ##### ##### ",
		"   #######   ",
		"     ###     ",
	}

	for _, currRow := range rows {
		for j := 0; j < hAmount; j++ {
			fmt.Print(currRow)
		}
		fmt.Println()
	}

}

func (g Game) LogGameInfo(quizN int, difLevel DifficultyLevel, hAmount int) {

	difficulties := []string{"Easy", "Medium", "Hard"}
	fmt.Printf("\nQuiz #%d  |  Difficulty: %s\n", quizN, difficulties[difLevel-1])
	g.DrawHearts(hAmount)

}

func (g Game) StartGame() {
	var difLevel DifficultyLevel = 1
	var quizN int = 1
	var result bool
	hearts := g.attempts
	for ; difLevel <= 3; difLevel++ {
		for _, v := range g.ds.GetQuestionsByDifficulty(difLevel) {
			if hearts == 0 {
				fmt.Println("Game over")
				return
			}
			g.LogGameInfo(quizN, difLevel, hearts)
			result = g.l.AskQuestion(v)
			if !result {
				hearts--
			}
			quizN++
		}
	}
	g.DrawHearts(hearts)
	fmt.Println("You won!")
}
