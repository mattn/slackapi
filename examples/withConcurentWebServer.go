package main

import (
	"fmt"
	"github.com/pastjean/go-slackbot"
	"log"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go World!")
	})

	errorsChan := make(chan error)

	go concurentHttp(errorsChan)
	go concurentSlackBot(&s, errorsChan)

	// Process Stays up until something crashes
	log.Fatal(<-errorsChan)
}

func concurentHttp(c chan<- error) {
	c <- http.ListenAndServe(":8080", nil)
}

func concurentSlackBot(s *slack.SlackBot, c chan<- error) {
	c <- s.Start()
}
