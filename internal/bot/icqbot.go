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

func NewIcqBot(token string) Bot {
	bot, err := botgolang.NewBot(token)
	if err != nil {
		panic(err)
	}
	return icqBot{
		bot: bot,
	}
}

// Send sends message with text to chat with chatID
func (b icqBot) Send(chatID int64, text string) {
	message := b.bot.NewTextMessage(strconv.FormatInt(chatID, 10), text)
	err := message.Send()
	if err != nil {
		log.Println(err.Error())
	}
}

// Updates returns a channel, which will be filled with events
func (b icqBot) Updates() interface{} {
	ctx, _ := context.WithCancel(context.Background())

	updates := b.bot.GetUpdatesChannel(ctx)
	if updates != nil {
		log.Println("Bot started")
	}

	return updates
}
