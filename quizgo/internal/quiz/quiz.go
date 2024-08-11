package quiz

import (
	"errors"
	"fmt"
	"quizgo/internal/model"
	"strconv"
)

type ID string

type Dif = model.DifficultyLevel
type Qst = model.Question
type Opts = model.AnswerOptions
type Ans = model.AnswerIndex

type FormattedQuiz struct {
	Question string
	Answer   string
	Options  string
}

type QuizDataStore struct {
	Questions map[model.DifficultyLevel][]FormattedQuiz
}

func (qds QuizDataStore) CreateQuiz(q string, opts [4]string, ansIndex int, difLevel int) model.Quiz {
	return model.Quiz{
		Question:   Qst(q),
		Options:    Opts(opts),
		Index:      Ans(ansIndex),
		Difficulty: Dif(difLevel),
	}
}

func (qds *QuizDataStore) Add(args ...interface{}) ([]ID, error) {
	var quizes []model.Quiz
	var IDs []ID

	for _, arg := range args {
		if arg, ok := arg.(model.Quiz); ok {
			if arg.Index > 3 || arg.Index < 0 {
				return nil, errors.New("answer index must correspond to the right answer from answer options")
			}
			if arg.Difficulty < 1 || arg.Difficulty > 3 {
				return nil, errors.New("difficulty level must be one of: 1, 2, 3")
			}
			quizes = append(quizes, arg)
		} else {
			return nil, errors.New("Add method supports only Quiz type as an argument")
		}
	}

	var fq FormattedQuiz
	var difLevel model.DifficultyLevel
	var id ID
	for _, v := range quizes {
		fq = qds.createFormattedQuiz(v)
		difLevel = v.Difficulty
		index := len(qds.Questions[difLevel])
		qds.Questions[difLevel] = append(qds.Questions[difLevel], fq)

		id = ID(strconv.Itoa(int(difLevel))) + ID(strconv.Itoa(index))
		IDs = append(IDs, id)
	}
	return IDs, nil
}

func (qds QuizDataStore) GetQuestion(id ID) (*FormattedQuiz, error) {
	difLevel, index, err := qds.parseID(id)
	if err != nil {
		return &FormattedQuiz{}, err
	}

	if difLevel < 1 || difLevel > 3 {
		return &FormattedQuiz{}, errors.New("difficulty level must be one of: 1, 2, 3")
	}

	return &qds.Questions[Dif(difLevel)][index], nil
}

func (qds QuizDataStore) EditQuiz(id ID, args ...interface{}) (ID, error) {
	quiz, err := qds.GetQuestion(id)
	if err != nil {
		return "", err
	}

	answerLetters := [4]string{"a", "b", "c", "d"}
	for _, v := range args {
		switch v := v.(type) {

		case model.Question:
			quiz.Question = string(v)
		case string:
			quiz.Question = v

		case model.AnswerIndex:
			quiz.Answer = answerLetters[v]

		case model.AnswerOptions:
			quiz.Options = qds.formatOpts(v)
		case [4]string:
			quiz.Options = qds.formatOpts(Opts(v))

		case model.DifficultyLevel:
			difLevel, index, err := qds.parseID(id)
			if err != nil {
				return "", err
			}

			qds.Questions[v] = append(qds.Questions[v], *quiz)

			qds.Questions[Dif(difLevel)] =
				append(qds.Questions[Dif(difLevel)][:index], qds.Questions[Dif(difLevel)][index+1:]...)

			id = ID(strconv.Itoa(difLevel)) + ID(strconv.Itoa(len(qds.Questions[v])))
		}
	}

	return id, nil
}

func (qds QuizDataStore) GetQuestionsByDifficulty(difLevel model.DifficultyLevel) []FormattedQuiz {
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

func (qds QuizDataStore) formatOpts(opts model.AnswerOptions) string {
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

func (qds QuizDataStore) createFormattedQuiz(quiz model.Quiz) FormattedQuiz {
	answerLetters := [4]string{"a", "b", "c", "d"}
	formattedOpts := qds.formatOpts(quiz.Options)

	q := fmt.Sprintf("%s", quiz.Question)
	a := answerLetters[quiz.Index]
	return FormattedQuiz{
		Question: q,
		Answer:   a,
		Options:  formattedOpts,
	}
}

func (qds QuizDataStore) parseID(id ID) (int, int, error) {
	difLevel, err := strconv.Atoi(string(id[0:1]))
	if err != nil {
		return 0, 0, err
	}

	index, err := strconv.Atoi(string(id[1:]))
	if err != nil {
		return 0, 0, err
	}

	return difLevel, index, nil
}
