package agent

import (
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

const (
	URI = "http://localhost:3000/socket.io/"
)

func GetOpts() *socketio_client.Options {
	return &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
}
