package main

import (
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"chgk-bot/internal/app"
)

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	log.Printf("Connect to Bot %s", bot.Self.UserName)
	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
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

func main() {
	token := os.Getenv("TELEGRAM_API_TOKEN")
	if len(token) == 0 {
		panic("TELEGRAM_API_TOKEN env not set!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	game := app.NewGame(bot)

	updates := getUpdates(bot)
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
