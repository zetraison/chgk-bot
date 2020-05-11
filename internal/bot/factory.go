package bot

import "log"

// Type bot provider type
type Type int

const (
	// Telegram bot provider
	Telegram Type = iota
	// ICQ bot provider
	ICQ
)

// Bot describes available bot functions
type Bot interface {
	// Send sends a message with text to chat with chatID passed as an argument
	Send(chatID int64, text string)
	// Updates returns a channel, which will be filled with events
	Updates() interface{}
}

// GetBot returns bot provider
func GetBot(botType Type, token string) Bot {
	switch botType {
	case Telegram:
		return NewTelegramBot(token)
	case ICQ:
		return NewIcqBot(token)
	default:
		log.Printf("type undefined")
		return nil
	}
}
