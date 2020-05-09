package database

import (
	"fmt"
)

const (
	correctAnswer = "Правильный ответ"
	comment       = "Комментарий"
)

// Question struct
// Used for parse xml representation of question
type Question struct {
	Question     string `xml:"Question"`
	Answer       string `xml:"Answer"`
	PassCriteria string `xml:"PassCriteria"`
	Authors      string `xml:"Authors"`
	Sources      string `xml:"Sources"`
	Comments     string `xml:"Comments"`
	Tournament   string `xml:"tournamentTitle"`
}

// GetAnswer returns question string representation
func (q *Question) String() string {
	return fmt.Sprintf("Question: %s\n Answer: %s\n PassCriteria: %s\n Author: %s\n Comments: %s\n",
		q.Question, q.Answer, q.PassCriteria, q.Authors, q.Comments)
}

// GetAnswer returns question answer
func (q *Question) GetAnswer(message string) string {
	text := fmt.Sprintf("%s. %s: %s", message, correctAnswer, q.Answer)

	if q.Comments != "" {
		text += fmt.Sprintf("\n\n%s: %s", comment, q.Comments)
	}

	return text
}
