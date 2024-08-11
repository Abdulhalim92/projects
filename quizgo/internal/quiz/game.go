package quiz

import (
	"fmt"
	"time"
)

type Logic interface {
	AskQuestion(fq FormattedQuiz) bool
}

type Game struct {
	l      Logic
	ds     DataStore
	hearts int
}

func (g Game) drawHearts(hAmount int) {
	rows := [6]string{
		`  ###   ###  `,
		` ##### ##### `,
		`  #########  `,
		`    #####    `,
		`     ###     `,
		`      #      `,
	}

	for _, currRow := range rows {
		for j := 0; j < hAmount; j++ {
			fmt.Print(currRow)
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func (g Game) LogInfo(quizN, hAmount int, difLevel Dif) {
	difficulties := []string{"Easy", "Medium", "Hard"}
	fmt.Printf("\nQuiz #%d  |  Difficulty: %s\n\n", quizN, difficulties[difLevel-1])
	g.drawHearts(hAmount)

}

func (g Game) Start() {
	start := time.Now()
	var difLevel Dif = 1
	var quizN int = 1
	var result bool
	hearts := g.hearts

	for ; difLevel <= 3; difLevel++ {
		for _, v := range g.ds.GetQuestionsByDifficulty(difLevel) {
			if hearts == 0 {
				fmt.Println("Game over")
				return
			}

			g.LogInfo(quizN, hearts, difLevel)
			result = g.l.AskQuestion(v)
			if !result {
				hearts--
			}
			quizN++
		}
	}

	timePassed := time.Now().Sub(start).Truncate(time.Second)

	fmt.Print("\nYou won!\n")
	fmt.Printf("Hearts ramained: %d  |  Complited in: %v\n", hearts, timePassed)
}
