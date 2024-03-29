package main

import (
	"github.com/zetraison/chgk-bot/pkg/database"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	botgolang "github.com/mail-ru-im/bot-golang"

	"github.com/zetraison/chgk-bot/internal/app"
	"github.com/zetraison/chgk-bot/internal/bot"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file")
	}

	token := os.Getenv("ICQ_BOT_TOKEN")
	if token == "" {
		panic("ICQ_BOT_TOKEN env not set!")
	}

	icqBot := bot.GetBot(bot.ICQ, token)
	chgkDB := database.NewDatabase(database.ChgkGame)
	sessions := make(map[string]app.Game)

	updates := icqBot.Updates().(<-chan botgolang.Event)

	// main messages loop
	for update := range updates {
		if update.Payload.Message() == nil {
			continue
		}

		var game app.Game
		userID := update.Payload.From.ID
		if v, ok := sessions[userID]; ok {
			game = v
		} else {
			game = app.NewGame(icqBot, chgkDB)
			sessions[userID] = game
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
