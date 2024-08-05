package main

import (
	"fmt"
	"quizgo/internal/quiz"
)

func main() {
	ds := quiz.NewQuizStore()
	l := quiz.NewQuizLogic(ds)

	quiz1 := ds.CreateQuiz(
		"What is the capital of France?",
		[4]string{"Berlin", "Madrid", "Paris", "Rome"},
		2,
		1)

	quiz2 := ds.CreateQuiz(
		"Which planet is known as the Red Planet?",
		[4]string{"Earth", "Mars", "Jupiter", "Venus"},
		1,
		1)

	quiz3 := ds.CreateQuiz(
		"Who wrote 'Romeo and Juliet'?",
		[4]string{"William Shakespeare", "Mark Twain", "Charles Dickens", "Jane Austen"},
		0,
		1)

	quiz4 := ds.CreateQuiz(
		"What is the smallest prime number?",
		[4]string{"1", "2", "3", "5"},
		1,
		2)
	quiz5 := ds.CreateQuiz(
		"Which element has the chemical symbol 'O'?",
		[4]string{"Gold", "Silver", "Osmium", "Oxygen"},
		3,
		2)
	quiz6 := ds.CreateQuiz(
		"In which year did the Titanic sink?",
		[4]string{"1905", "1912", "1918", "1923"},
		1,
		2)
	quiz7 := ds.CreateQuiz(
		"Who developed the theory of general relativity?",
		[4]string{"Albert Einstein", "Isaac Newton", "Niels Bohr", "Galilio Galilei"},
		0,
		3)
	quiz8 := ds.CreateQuiz(
		"What is the longest river in the world?",
		[4]string{"Amazon", "Yangtze", "Nile", "Mississippi"},
		2,
		3)
	quiz9 := ds.CreateQuiz(
		"Which author wrote 'One Hundred Years of Solitude'?",
		[4]string{"Islabel Allende", "Mario Vargas Llosa", "Jorge Luis Borges", "Gabriel Garcia"},
		3,
		3)

	err := ds.Add(quiz1, quiz2, quiz3, quiz4, quiz5, quiz6, quiz7, quiz8, quiz9)
	if err != nil {
		fmt.Println(err)
		return
	}

	l.StartQuiz()
}
