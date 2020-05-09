package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	botgolang "github.com/mail-ru-im/bot-golang"

	"github.com/zetraison/chgk-bot/internal/app"
	"github.com/zetraison/chgk-bot/internal/bot"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("BOT_TOKEN env not set!")
	}

	icqBot := bot.GetBot(bot.ICQ, token)
	game := app.NewGame(icqBot)

	updates := icqBot.Updates().(<-chan botgolang.Event)

	// main messages loop
	for update := range updates {
		if update.Payload.Message() == nil {
			continue
		}

		chatID, err := strconv.ParseInt(update.Payload.Chat.ID, 10, 64)
		if err != nil {
			log.Printf("Can not parse chatID")
			continue
		}
		username := update.Payload.From.FirstName
		text := update.Payload.Text

		log.Printf("[%d][%s] %s", chatID, username, text)

		if strings.HasPrefix(text, "/") {
			command := strings.TrimPrefix(text, "/")
			go game.HandleCommand(chatID, username, command)
			continue
		}

		if !game.Active() {
			continue
		}

		go game.CheckAnswer(chatID, username, text)
	}
}
