package main

import (
	"flag"
	"log"
)

func main() {
	// t := mustToken()
	// token = flags.Get(token)

	// tgClient = telegram.New(token)

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