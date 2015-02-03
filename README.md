# slag

[Slack API](https://api.slack.com/) client implemented in go, it should handle everything simply and
peacefully.

# Usage

see examples folder but here's the simple example

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pastjean/slackbot"
	"github.com/pastjean/slackbot/rtm"
)

func main() {
	s := slackbot.NewSlackBot()

	s.OnMessageEvents(func(evt rtm.MessageEvent) {
		fmt.Printf("%v\n", evt)
	})

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN environment variable should be set")
	}

	s.SetToken(token)
	log.Fatal(s.Start())
}
```

# Not implemented

- [Messages subtypes](https://api.slack.com/events/message)
- Lots of standard api calls
