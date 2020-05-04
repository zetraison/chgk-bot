package game

const (
	Start    = "start"
	Help     = "help"
	Question = "question"
	Round    = "round"
	Exit     = "exit"
	Score    = "score"
)

func getCommandDescription() string {
	return "Список команд:\n\n" +
		"/help - показать список доступных команд\n" +
		"/question - получить случайный вопрос\n" +
		"/round - начать раунд из 10 вопросов\n" +
		"/exit - прервать раунд\n" +
		"/score - получить ваш результат\n"
}
