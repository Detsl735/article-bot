package telegram

import "github.com/Detsl735/article-bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
