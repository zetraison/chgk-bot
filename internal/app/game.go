package app

import (
	"fmt"
	"log"
	"time"

	"github.com/zetraison/chgk-bot/internal/bot"
	"github.com/zetraison/chgk-bot/internal/util"
	"github.com/zetraison/chgk-bot/pkg/database"
)

const (
	trueMessage            = "Верно"
	falseMessage           = "К сожалению, ваш ответ неверный"
	roundCompleteMessage   = "Время истекло"
	roundAlreadyActive     = "Игра уже активна"
	endTimeMessage         = "Осталось 10 секунд"
	yourResultMessage      = "Ваш результат"
	commandNotExistMessage = "Комманды не существует"
	rulesMessage           = "Правила:\n\n" +
		"- В режиме случайного вопроса (/question), после получения вопроса необходимо придумать и отправить ответ " +
		"в течении 2 минут. На ответ дается 1 попытка.\n\n" +
		"- В режиме игры (/round), стартует цикл из 10 вопросов, в котором необходимо в течении 2х минут придумать и " +
		"отправить ответ на каждый вопрос. Интервал между вопросами 1 минута"

	thinkingTime    = time.Second * 50
	finishTime      = time.Second * 10
	questionInRound = 3
)

// Game describes available game functions
type Game interface {
	Active() bool
	HandleCommand(chatID int64, username, command string)
	CheckAnswer(chatID int64, username, answer string)
}

type game struct {
	db            database.Database
	bot           bot.Bot
	question      chan *database.Question
	score         map[string]int
	thinkingTimer *time.Timer
	finishTimer   *time.Timer
}

// NewGame returns new game instance
func NewGame(_bot bot.Bot, db database.Database) Game {
	return &game{
		db:       db,
		question: make(chan *database.Question, 1),
		score:    make(map[string]int),
		bot:      _bot,
	}
}

// Active returns active game status.
// Game active where questions pool channel is not empty
func (g *game) Active() bool {
	return len(g.question) > 0
}

// HandleCommand handle command from user with username and send result to chat by chatID
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
			g.sendMessage(chatID, roundAlreadyActive)
			return
		}
		g.startRound(chatID)
	case Stop:
		g.stop()
		g.sendScore(chatID, username)
	default:
		g.sendMessage(chatID, commandNotExistMessage)
	}
}

// CheckAnswer checks user answer with expected question answer and send check result to chat by chatID
func (g *game) CheckAnswer(chatID int64, username, answer string) {
	question := <-g.question

	if g.thinkingTimer != nil {
		g.thinkingTimer.Stop()
	}
	if g.finishTimer != nil {
		g.finishTimer.Stop()
	}

	equal := util.CompareStrings(answer, question.Answer)

	message := falseMessage
	if equal {
		message = trueMessage
		g.score[username]++
	}

	g.sendMessage(chatID, question.GetAnswer(message))
}

// Stop removes all questions from pool channel
func (g *game) stop() {
	for len(g.question) != 0 {
		<-g.question
	}
}

// sendQuestion send one question to chat by chatID
func (g *game) sendQuestion(chatID int64) {
	question, err := g.db.GetQuestion()
	if err != nil {
		log.Printf("Error on get question: %s", err.Error())
		return
	}
	g.sendMessage(chatID, question.Question)
	g.question <- question

	g.thinkingTimer = time.NewTimer(thinkingTime)
	<-g.thinkingTimer.C
	g.sendMessage(chatID, endTimeMessage)

	g.finishTimer = time.NewTimer(finishTime)
	<-g.finishTimer.C
	<-g.question
	g.sendMessage(chatID, question.GetAnswer(roundCompleteMessage))
}

// startRound runs game of 10 questions and send its to chat by chatID
func (g *game) startRound(chatID int64) {
	packet, err := g.db.GetQuestionPacket(questionInRound)
	if err != nil {
		log.Printf("Error on get questions packet: %s", err.Error())
		return
	}
	for _, question := range packet.Questions {
		g.sendMessage(chatID, question.Question)
		g.question <- question

		g.thinkingTimer = time.NewTimer(thinkingTime)
		<-g.thinkingTimer.C
		g.sendMessage(chatID, endTimeMessage)

		g.finishTimer = time.NewTimer(finishTime)
		<-g.finishTimer.C
		<-g.question
		g.sendMessage(chatID, question.GetAnswer(roundCompleteMessage))
	}
}

// sendScore sends score of username to chat by chatID
func (g *game) sendScore(chatID int64, username string) {
	score := fmt.Sprintf("%s: %d\n", yourResultMessage, g.score[username])
	g.sendMessage(chatID, score)
}

// sendDescription sends short descriptions of game rules and help information
func (g *game) sendDescription(chatID int64) {
	description := fmt.Sprintf("%s\n%s", getCommandDescription(), rulesMessage)
	g.sendMessage(chatID, description)
}

// sendMessage sends message to chat by chatID
func (g *game) sendMessage(chatID int64, text string) {
	g.bot.Send(chatID, text)
}
