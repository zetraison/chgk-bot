package app

const (
	// show help info command (on bot start)
	Start = "start"
	// show help info command
	Help = "help"
	// send random question command
	Question = "question"
	// start round of game
	Round = "round"
	// stop round of game
	Stop = "stop"
	// show results
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
