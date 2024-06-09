package main

import (
	"flag"
	"log"
	"tgtest/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(mustToken())
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
