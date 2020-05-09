package bot

import (
	"context"
	"log"
	"strconv"

	botgolang "github.com/mail-ru-im/bot-golang"
)

type icqBot struct {
	bot *botgolang.Bot
}

// NewIcqBot returns new bot instance
func NewIcqBot(token string) Bot {
	bot, err := botgolang.NewBot(token)
	if err != nil {
		panic(err)
	}
	return icqBot{
		bot: bot,
	}
}

// Send sends a message with text to chat with chatID passed as an argument
func (b icqBot) Send(chatID int64, text string) {
	message := b.bot.NewTextMessage(strconv.FormatInt(chatID, 10), text)
	err := message.Send()
	if err != nil {
		log.Println(err.Error())
	}
}

// Updates returns a channel, which will be filled with events
func (b icqBot) Updates() interface{} {
	updates := b.bot.GetUpdatesChannel(context.TODO())
	if updates != nil {
		log.Println("Bot started")
	}

	return updates
}
