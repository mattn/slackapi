package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pastjean/slackbot/api"
	"github.com/pastjean/slackbot/rtm"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN environment variable should be set")
	}

	resp, err := api.RtmStart(token)
	if err != nil {
		log.Fatal(err)
	}

	slackrtm := rtm.NewSlackRTM(resp)

	// You can directly access the EventController underneat
	slackrtm.OnMessageEvents(func(evt rtm.MessageEvent) {
		fmt.Printf("%v\n", evt)
	})

	log.Fatal(slackrtm.Start())
}
