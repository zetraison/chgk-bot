package main

import (
	"chgk-bot/internal/app"
	"chgk-bot/internal/bot"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	//"strings"
	//"time"

	"github.com/mail-ru-im/bot-golang"
	//"chgk-bot/internal/app"
)

func main() {
	token := os.Getenv("ICQ_BOT_TOKEN")
	if len(token) == 0 {
		panic("ICQ_BOT_TOKEN env not set!")
	}

	bot := bot.GetBot(bot.ICQ, token)
	game := app.NewGame(bot)

	updates := bot.GetUpdates().(<-chan botgolang.Event)

	// main messages loop
	for update := range updates {
		if update.Payload.Message() == nil {
			continue
		}

		chatID, err := strconv.ParseInt(update.Payload.Chat.ID, 10, 64)
		if err != nil {
			fmt.Printf("Can not parse chatID")
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
