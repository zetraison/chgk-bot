package app

const (
	// Start shows help info (on bot start)
	Start = "start"
	// Help shows help info
	Help = "help"
	// Question sends random question
	Question = "question"
	// Round starts round of game
	Round = "round"
	// Stop stops round of game
	Stop = "stop"
	// Score shows results
	Score = "score"
)

func getCommandDescription() string {
	return "Список команд:\n\n" +
		"/help - показать список доступных команд\n" +
		"/question - получить случайный вопрос\n" +
		"/round - начать раунд из 10 вопросов\n" +
		"/stop - прервать раунд из 10 вопросов\n" +
		"/score - получить ваш результат\n"
}
