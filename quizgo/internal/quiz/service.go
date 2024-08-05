package quiz

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
