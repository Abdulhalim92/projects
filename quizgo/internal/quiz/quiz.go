package quiz

import (
	"errors"
	"fmt"
	"quizgo/internal/model"
)

type DifficultyLevel = model.DifficultyLevel

type FormattedQuiz struct {
	question string
	answer   string
}

type QuizDataStore struct {
	Questions map[DifficultyLevel][]FormattedQuiz
}

func (qds QuizDataStore) CreateQuiz(q string, opts [4]string, ansIndex int, difLevel DifficultyLevel) model.Quiz {

	return model.Quiz{
		Question:      q,
		AnswerOptions: opts,
		AnswerIndex:   ansIndex,
		Difficulty:    difLevel,
	}
}

func (qds *QuizDataStore) Add(args ...interface{}) error {
	var quizes []model.Quiz
	for _, arg := range args {
		switch arg := arg.(type) {
		case model.Quiz:
			if arg.AnswerIndex > 3 || arg.AnswerIndex < 0 {
				return errors.New("answer index must correspond to the right answer from answer options")
			}
			if arg.Difficulty < 1 || arg.Difficulty > 3 {
				return errors.New("difficulty level must be one of: 1, 2, 3")
			}
			quizes = append(quizes, arg)

		default:
			return errors.New("Add method supports only Quiz type as an argument")
		}
	}

	var fq FormattedQuiz
	var difLevel DifficultyLevel
	for _, v := range quizes {
		fq = qds.GetFormattedQuiz(v)
		difLevel = v.Difficulty
		qds.Questions[difLevel] = append(qds.Questions[difLevel], fq)
	}
	return nil
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

func (qds QuizDataStore) GetFormattedQuiz(quiz model.Quiz) FormattedQuiz {
	answerLetters := [4]string{"a", "b", "c", "d"}
	formattedOpts := qds.FormatOpts(quiz.AnswerOptions)

	q := fmt.Sprintf("%s\n\n%s", quiz.Question, formattedOpts)
	a := answerLetters[quiz.AnswerIndex]
	return FormattedQuiz{
		question: q,
		answer:   a,
	}
}

func (qds QuizDataStore) GetQuestion(difLevel DifficultyLevel, qIndex int) (FormattedQuiz, error) {
	if difLevel < 1 || difLevel > 3 {
		return FormattedQuiz{}, errors.New("difficulty level must be one of: 1, 2, 3")
	}
	return qds.Questions[difLevel][qIndex], nil
}

func (qds QuizDataStore) GetQuestionsByDifficulty(difLevel DifficultyLevel) []FormattedQuiz {
	return qds.Questions[difLevel]
}

func (qds QuizDataStore) GetEasyQuestions() []FormattedQuiz {
	return qds.Questions[1]
}
func (qds QuizDataStore) GetMediumQuestion() []FormattedQuiz {
	return qds.Questions[2]
}
func (qds QuizDataStore) GetHardQuestions() []FormattedQuiz {
	return qds.Questions[3]
}
