package bot

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegramBot struct {
	bot *tgbotapi.BotAPI
}

// NewTelegramBot returns new bot instance
func NewTelegramBot(token string) Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return telegramBot{
		bot: bot,
	}
}

// Send sends a message with text to chat with chatID passed as an argument
func (b telegramBot) Send(chatID int64, text string) {
	if _, err := b.bot.Send(tgbotapi.NewMessage(chatID, text)); err != nil {
		log.Printf("Error on send message: %s", err.Error())
	}
}

// Updates returns a channel, which will be filled with events
func (b telegramBot) Updates() interface{} {
	//b.bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	log.Printf("Connect to Bot %s", b.bot.Self.UserName)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	log.Printf("Old messages cleared from updates")

	return updates
}
