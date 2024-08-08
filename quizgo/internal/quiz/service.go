package quiz

func NewGame(l Logic, ds DataStore, heartsAmount int) Game {
	return Game{
		l:      l,
		ds:     ds,
		hearts: heartsAmount,
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
