package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	"github.com/zetraison/chgk-bot/internal/app"
	"github.com/zetraison/chgk-bot/internal/bot"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		panic("TELEGRAM_BOT_TOKEN env not set!")
	}

	telegramBot := bot.GetBot(bot.Telegram, token)
	game := app.NewGame(telegramBot)

	updates := telegramBot.Updates().(tgbotapi.UpdatesChannel)

	// main messages loop
	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		username := update.Message.Chat.UserName
		text := update.Message.Text

		log.Printf("[%d][%s] %s", chatID, username, text)

		if update.Message.IsCommand() {
			command := update.Message.Command()
			go game.HandleCommand(chatID, username, command)
			continue
		}

		if !game.Active() {
			continue
		}

		go game.CheckAnswer(chatID, username, text)
	}
}
