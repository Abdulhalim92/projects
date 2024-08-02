package quiz

import (
	"fmt"
	"quizgo/internal/model"
)

var lastID int

func init() {
	lastID = -1
}

type QuizDataStore struct {
	questions     map[int]string
	answerOptions map[int][4]string
	answers       map[int]string
}

func (qds QuizDataStore) QuestionForID(qID int) string {
	q := qds.questions[qID]
	return q
}

func (qds QuizDataStore) AnswerForID(qID int) string {
	answerLetter := qds.answers[qID]

	return answerLetter
}

func (qds QuizDataStore) OptionsForID(qID int) [4]string {
	opts := qds.answerOptions[qID]
	return opts
}

func (qds QuizDataStore) CreateQuiz(q string, opts [4]string, ansIndex int) model.Quiz {
	lastID++
	return model.Quiz{
		QuestionID:    lastID,
		Question:      q,
		AnswerOptions: opts,
		AnswerIndex:   ansIndex,
	}
}

func (qds QuizDataStore) AddQuizData(quiz model.Quiz) {
	answerLetters := [4]string{"a", "b", "c", "d"}
	qds.questions[quiz.QuestionID] = quiz.Question
	qds.answerOptions[quiz.QuestionID] = quiz.AnswerOptions
	qds.answers[quiz.QuestionID] = answerLetters[quiz.AnswerIndex]
}

func (qds QuizDataStore) GetQuestionsLen() int {
	return len(qds.questions)
}

func (qds QuizDataStore) FormatOpts(opts [4]string) string {
	firstRowGap := 20 - len(opts[0])
	secondRowGap := 20 - len(opts[2])
	var firstRowSpace, secondRowSpace string
	for i := 0; i < firstRowGap; i++ {
		firstRowSpace += " "
	}
	for i := 0; i < secondRowGap; i++ {
		secondRowSpace += " "
	}
	return fmt.Sprintf("     A) %v%vB) %v\n\n     C) %v%vD) %v",
		opts[0], firstRowSpace, opts[1],
		opts[2], secondRowSpace, opts[3])
}

func (qds QuizDataStore) GetFormattedQuiz(qID int) string {
	fmt.Printf("Quiz #%d \n", qID+1)
	q := qds.QuestionForID(qID)
	opts := qds.OptionsForID(qID)
	formattedOpts := qds.FormatOpts(opts)

	return fmt.Sprintf("%v\n%v", q, formattedOpts)
}
