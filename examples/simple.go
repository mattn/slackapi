package main

import (
	"fmt"
	slack "github.com/pastjean/slackbot"
	"log"
	"os"
)

func main() {
	s := slack.NewSlackBot()

	s.OnMessageEvents(func(evt slack.MessageEvent) {
		fmt.Printf("%v\n", evt)
	})

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN environment variable should be set")
	}

	s.SetToken(token)
	log.Fatal(s.Start())
}
