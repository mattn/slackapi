package main

import (
	"fmt"
	"github.com/pastjean/go-slackbot"
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
		token = "xoxb-3373459148-unuNTStmar3GI9UzZqzgJ2n3"
	}

	s.SetToken(token)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
