package src

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type SharedData struct {
	Channel  string
	Filepath string
	Data     []byte
	Filename string
}

func InitServer() {

	server := setupServer()

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("../app/dist")))
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func setupServer() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		return nil
	})

	server.OnEvent("/", "subscription", func(s socketio.Conn, channel string) {
		log.Printf("channel %s is connected", channel)

		// join client to room c-[channel]
		server.JoinRoom("/", fmt.Sprint("c", channel), s)

		// join client to room channels
		server.JoinRoom("/", "channels", s)
		log.Printf("connected clients: %+v", server.RoomLen("/", "channels"))
		server.BroadcastToRoom("/", "app", "connections", server.RoomLen("/", "channels"))
	})

	server.OnEvent("/", "send", func(s socketio.Conn, data SharedData) {
		log.Printf("client is sending file: %s data by channel %s", data.Filename, data.Channel)
		server.BroadcastToRoom("/", fmt.Sprint("c", data.Channel), "reply", data)
	})

	// event to get app connection
	server.OnEvent("/", "app", func(s socketio.Conn, msg string) {
		log.Println(msg)

		// join client to app room
		server.JoinRoom("/", "app", s)
		server.BroadcastToRoom("/", "app", "connections", server.RoomLen("/", "channels"))
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("client disconnected")
		server.BroadcastToRoom("/", "app", "connections", server.RoomLen("/", "channels"))
	})

	return server
}
