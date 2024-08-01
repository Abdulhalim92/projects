package main

import (
	"quizgo/internal/quiz"
)

func main() {
	ds := quiz.NewQuizStore()
	l := quiz.NewQuizLogic(ds)

	quiz1 := ds.CreateQuiz(
		"What is the name of the highest mountain?",
		[4]string{"Kilimanjaro", "Makalu", "Everest", "Ismoil Somoni Peak"},
		2)

	quiz2 := ds.CreateQuiz(
		"Which animal is largest on Earth?",
		[4]string{"Blue whale", "Elephant", "Polar bear", "Crocodile"},
		0)

	quiz3 := ds.CreateQuiz(
		"Who is the wealthiest person in the world?",
		[4]string{"Jeff Bezos", "Elon Musk", "Bill Gates", "Mark Zuckerberg"},
		1)

	ds.AddQuizData(quiz1)
	ds.AddQuizData(quiz2)
	ds.AddQuizData(quiz3)

	l.StartQuiz()
}
