package telegram

import "tgtest/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

func New(client *telegram.Client)
