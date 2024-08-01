package quiz

func NewQuizLogic(ds DataStore) QuizLogic {
	return QuizLogic{
		ds: ds,
	}
}

func NewQuizStore() QuizDataStore {
	return QuizDataStore{
		questions: map[int]string{
			// 0: "What is the name of the highest mountain?",
			// 1: "Which animal is the largest on Earth?",
			// 2: "Who is the welthiest person in the world?",
		},

		answers: map[int]string{
			// 0: "Everest",
			// 1: "Blue whale",
			// 2: "Elon Musk",
		},

		answerOptions: map[int][4]string{
			// 0: {"Kilimanjaro", "Makalu", "Everest", "Ismoil Somoni Peak"},
			// 1: {"Blue whale", "Elephant", "Crocodile", "Polar bear"},
			// 2: {"Jeff Bezos", "Elon Musk", "Bill Gates", "Mark Zuckerberg"},
		},
	}
}
