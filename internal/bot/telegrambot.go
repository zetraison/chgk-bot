package bot

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return telegramBot{
		bot: bot,
	}
}

func (b telegramBot) Send(chatID int64, text string) {
	if _, err := b.bot.Send(tgbotapi.NewMessage(chatID, text)); err != nil {
		log.Printf("Error on send message: %s", err.Error())
	}
}

func (b telegramBot) GetUpdates() interface{} {
	log.Printf("Connect to Bot %s", b.bot.Self.UserName)
	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	log.Printf("Old messages cleared from updates")

	return updates
}