package app

const (
	Start    = "start"
	Help     = "help"
	Question = "question"
	Round    = "round"
	Stop     = "stop"
	Score    = "score"
)

func getCommandDescription() string {
	return "Список команд:\n\n" +
		"/help - показать список доступных команд\n" +
		"/question - получить случайный вопрос\n" +
		"/round - начать раунд из 10 вопросов\n" +
		"/stop - прервать раунд из 10 вопросов\n" +
		"/score - получить ваш результат\n"
}
