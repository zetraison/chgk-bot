package main

import (
	"fmt"
)

type Question struct {
	Question     string `xml:"Question"`
	Answer       string `xml:"Answer"`
	PassCriteria string `xml:"PassCriteria"`
	Authors      string `xml:"Authors"`
	Sources      string `xml:"Sources"`
	Comments     string `xml:"Comments"`
	Tournament   string `xml:"tournamentTitle"`
}

type Packet struct {
	Questions []*Question `xml:"question"`
}

func (q *Question) String() string {
	return fmt.Sprintf("Question: %s\n Answer: %s\n PassCriteria: %s\n Author: %s\n Comments: %s\n",
		q.Question, q.Answer, q.PassCriteria, q.Authors, q.Comments)
}
