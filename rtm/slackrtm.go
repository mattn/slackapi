package rtm

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func Dial(url string) (*websocket.Conn, error) {
	var dialer *websocket.Dialer
	var dialHeader http.Header

	conn, _, err := dialer.Dial(url, dialHeader)

	return conn, err
}
