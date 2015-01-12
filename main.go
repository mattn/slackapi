package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

var s SlackRtm

func main() {
	s := NewSlackRtm()
	s.OnUnknownEvent(func(evt Event, evtstring []byte) {
		fmt.Println("%s", evtstring)
	})

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		token = "xoxb-3373459148-unuNTStmar3GI9UzZqzgJ2n3"
	}

	conn, err := startSlackRtm(token)
	if err != nil {
		log.Fatal(err)
	}

	readLoop(conn)
}

func readLoop(c *websocket.Conn) {
	for {
		_, r, err := c.NextReader()

		if err != nil {
			log.Fatal(err)
			return
		}

		err = parseWSEvent(r)
		if err != nil {
			log.Print(err)
		}

	}
}
