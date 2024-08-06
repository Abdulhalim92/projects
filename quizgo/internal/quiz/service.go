package quiz

func NewGame(l Logic, ds DataStore, attempts int) Game {
	return Game{
		l:        l,
		ds:       ds,
		attempts: attempts,
	}
}

func NewQuizLogic(ds DataStore) QuizLogic {
	return QuizLogic{
		ds: ds,
	}
}

func NewQuizStore() QuizDataStore {
	return QuizDataStore{
		Questions: map[DifficultyLevel][]FormattedQuiz{},
	}
}
