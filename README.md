# slackbot

A simple slackbot framework that handles everythings peacefully!

# Usage

see examples folder but here's the simple example

```go
package main

import (
  "fmt"
  "github.com/pastjean/slackbot"
  "log"
  "os"
  )

func main() {
  s := slackbot.NewSlackBot()

  s.OnMessageEvents(func(evt slackbot.MessageEvent) {
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
