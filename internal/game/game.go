package game

import (
	"fmt"
	"log"
	"time"

	"chgk-telegram-bot/internal/util"
	"chgk-telegram-bot/pkg/chgkapi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	trueMessage          = "Верно"
	falseMessage         = "К сожалению, ваш ответ неверный"
	roundCompleteMessage = "Время истекло"
	roundAlreadyActive   = "Игра уже активна"
	endTimeMessage       = "Осталось 10 секунд"
	yourResultMessage    = "Ваш результат"

	warningTime = time.Second * 50
	roundTime   = time.Second * 60
)

type Game interface {
	Active() bool
	HandleCommand(chatID int64, username, command string)
	CheckAnswer(chatID int64, username, answer string)
}

type game struct {
	db           chgkapi.Database
	question     chan *chgkapi.Question
	score        map[string]int
	bot          *tgbotapi.BotAPI
	warningTimer *time.Timer
	roundTimer   *time.Timer
}

func NewGame(bot *tgbotapi.BotAPI) Game {
	return &game{
		db:       chgkapi.NewDatabase(chgkapi.ChgkGame),
		question: make(chan *chgkapi.Question, 1),
		score:    make(map[string]int, 0),
		bot:      bot,
	}
}

func (g *game) Active() bool {
	return len(g.question) > 0
}

func (g *game) HandleCommand(chatID int64, username, command string) {
	switch command {
	case Start:
		g.sendDescription(chatID)
	case Help:
		g.sendDescription(chatID)
	case Score:
		g.sendScore(chatID, username)
	case Question:
		if g.Active() {
			g.sendMessage(chatID, roundAlreadyActive)
			return
		}
		g.sendQuestion(chatID)
	case Round:
		if g.Active() {
			return
		}
		g.sendQuestion(chatID)
	default:
		return
	}
}

func (g *game) CheckAnswer(chatID int64, username, answer string) {
	question := <-g.question

	if g.warningTimer != nil {
		g.warningTimer.Stop()
	}
	if g.roundTimer != nil {
		g.roundTimer.Stop()
	}

	equal := util.CompareStrings(answer, question.Answer)

	message := falseMessage
	if equal {
		message = trueMessage
		g.score[username]++
	}

	g.sendMessage(chatID, question.GetAnswer(message))
}

func (g *game) sendScore(chatID int64, username string) {
	score := fmt.Sprintf("%s: %d\n", yourResultMessage, g.score[username])

	g.sendMessage(chatID, score)
}

func (g *game) sendQuestion(chatID int64) {
	question, err := g.db.GetQuestion()
	if err != nil {
		log.Printf("Error on get question: %s", err.Error())
		return
	}
	g.question <- question
	g.sendMessage(chatID, question.Question)

	g.warningTimer = time.NewTimer(warningTime)
	<-g.warningTimer.C
	g.sendMessage(chatID, endTimeMessage)

	g.roundTimer = time.NewTimer(roundTime)
	<-g.roundTimer.C
	<-g.question
	g.sendMessage(chatID, question.GetAnswer(roundCompleteMessage))
}

func (g *game) sendDescription(chatID int64) {
	description := "Правила:\nПосле получения нового вопроса необходимо придумать и отправить ответ в течении 2 минут. На ответ дается 1 попытка"
	description = fmt.Sprintf("%s\n%s", getCommandDescription(), description)

	g.sendMessage(chatID, description)
}

func (g *game) sendMessage(chatID int64, message string) {
	if _, err := g.bot.Send(tgbotapi.NewMessage(chatID, message)); err != nil {
		log.Printf("Error on send message: %s", err.Error())
	}
}
