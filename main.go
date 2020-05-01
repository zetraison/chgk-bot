package main

import "log"

func main() {
	db := NewDatabase(chgkGame)

	// get one random question
	question, err := db.GetQuestion()
	if err != nil {
		panic(err)
	}
	log.Printf(question.Question)

	// get questions packet
	packet, err := db.GetQuestionPacket(10)
	if err != nil {
		panic(err)
	}
	for _, question := range packet.Questions {
		log.Printf(question.Question)
	}
}
