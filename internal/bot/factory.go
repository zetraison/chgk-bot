package bot

import (
	"log"
)

type Type int

const (
	Telegram Type = iota
	ICQ
)

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
