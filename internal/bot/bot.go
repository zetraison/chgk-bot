package bot

type Bot interface {
	Send(chatID int64, test string)
	GetUpdates() interface{}
}
