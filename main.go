package main

import (
	"flag"
	"log"
	"firstGoBot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)
	
	// consumer.Start(fetcher, processor)
}

func mustToken() string { 
	token := flag.String(
		"token-bot",
		"",
		"token for acces to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token if not specified")
	}

	return *token
}