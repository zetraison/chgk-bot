package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	helpCommand      = "help"
	openMenuCommand  = "openmenu"
	closeMenuCommand = "closemenu"
	questionCommand  = "question"
	scoreCommand     = "score"

	trueMessage           = "Верно!"
	falseMessage          = "К сожалению, ваш ответ неверный. Правильный ответ: %s"
	roundComplete         = "Время истекло. Правильный ответ: %s"
	yourResult            = "Ваш результат"
	unknownCommandMessage = "Неизвестная команда"
	tryAgainMessage       = "Извините, попробуйте еще раз"
	commandListMessage    = "Список команд"
	menuCloseMessage      = "Меню закрыто"
	endTimeMessage        = "Осталось 10 секунд.."
	commentMessage        = "\n\nКомментарий: %s"

	warningTime = time.Second * 50
	roundTime   = time.Second * 60
)

func showHelp() string {
	return "Список доступных команд:\n\n" +
		"/help - список доступных команд\n" +
		"/question - получить случайный вопрос\n" +
		"/openmenu - открыть меню\n" +
		"/closemenu - закрыть меню\n" +
		"/score - получить ваш результат\n"
}

func main() {
	var question *Question
	var message tgbotapi.MessageConfig
	var roundTimer, warningTimer *time.Timer

	tgApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	if len(tgApiToken) == 0 {
		panic("Telegram API token not set!")
	}

	bot, err := tgbotapi.NewBotAPI(tgApiToken)
	if err != nil {
		panic(err)
	}
	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	var keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", questionCommand)),
			tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", scoreCommand)),
			tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", helpCommand)),
			tgbotapi.NewKeyboardButton(fmt.Sprintf("/%s", closeMenuCommand)),
		),
	)

	db := NewDatabase(chgkGame)
	score := make(map[string]int, 0)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case helpCommand:
				message = tgbotapi.NewMessage(update.Message.Chat.ID, showHelp())
			case openMenuCommand:
				message = tgbotapi.NewMessage(update.Message.Chat.ID, commandListMessage)
				message.ReplyMarkup = keyboard
			case closeMenuCommand:
				message = tgbotapi.NewMessage(update.Message.Chat.ID, menuCloseMessage)
				message.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			case scoreCommand:
				message = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s: %d", yourResult, score[update.Message.Chat.UserName]))
			case questionCommand:
				question, err = db.GetQuestion()
				if err != nil {
					log.Printf("Error on get question: %s", err.Error())
					message = tgbotapi.NewMessage(update.Message.Chat.ID, tryAgainMessage)
					break
				}
				log.Printf(question.String())
				message = tgbotapi.NewMessage(update.Message.Chat.ID, question.Question)

				warningTimer = time.NewTimer(warningTime)
				go func() {
					<-warningTimer.C
					message = tgbotapi.NewMessage(update.Message.Chat.ID, endTimeMessage)
					bot.Send(message)
				}()

				roundTimer = time.NewTimer(roundTime)
				go func() {
					<-roundTimer.C
					text := fmt.Sprintf(roundComplete, question.Answer)
					if question.Comments != "" {
						comment := fmt.Sprintf(commentMessage, question.Comments)
						text += comment
					}
					message = tgbotapi.NewMessage(update.Message.Chat.ID, text)
					bot.Send(message)
				}()
			default:
				message = tgbotapi.NewMessage(update.Message.Chat.ID, unknownCommandMessage)
			}

			bot.Send(message)
			continue
		}

		if question == nil {
			continue
		}

		text := fmt.Sprintf(falseMessage, question.Answer)
		equal := compareString(update.Message.Text, question.Answer)
		if equal {
			text = trueMessage
			score[update.Message.Chat.UserName]++
		}
		if question.Comments != "" {
			comment := fmt.Sprintf(commentMessage, question.Comments)
			text += comment
		}
		message = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(message)

		warningTimer.Stop()
		roundTimer.Stop()
		question = nil
	}
}
