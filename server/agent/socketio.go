package agent

import (
	"fmt"
	"log"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kberlanga/tcp-vue-go/server/agent/models"
)

type SharedData struct {
	Channel  string
	Filepath string
	Data     []byte
	Filename string
}

func (sd ServerData) setupServer() *socketio.Server {
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

		sd.Connections = append(sd.Connections, models.Connection{
			ClientID:         s.ID(),
			Channel:          channel,
			SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
		})

		sd.Stadistics = append(sd.Stadistics, models.ShippingStatistics{
			ToClient:  s.ID(),
			ToChannel: channel,
			Status:    "PENDING",
		})

		server.BroadcastToRoom("/", "app", "connections", sd.Connections)

		// logs
		log.Printf("connected clients: %+v", server.RoomLen("/", "channels"))
	})

	server.OnEvent("/", "send", func(s socketio.Conn, data SharedData) {
		log.Printf("client is sending file: %s data by channel %s", data.Filename, data.Channel)
		indexes := sd.findStadisticChannel(data.Channel)
		limit := server.RoomLen("/", fmt.Sprint("c", data.Channel))
		for _, index := range indexes[:limit] {
			if sd.Stadistics[index].Status == "SUCCESS" {
				sd.Stadistics = append(sd.Stadistics, models.ShippingStatistics{
					ToClient:     sd.Stadistics[index].ToClient,
					ToChannel:    data.Channel,
					Status:       "RECEIVING",
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				})

				continue
			}

			sd.Stadistics[index].ShippingTime = time.Now().UTC().Format("2006-01-02 15:04:05Z")
			sd.Stadistics[index].Status = "RECEIVING"
		}

		if len(indexes) > 0 {
			server.BroadcastToRoom("/", fmt.Sprint("c", data.Channel), "reply", data)
		}
	})

	server.OnEvent("/", "received", func(s socketio.Conn, data SharedData) {
		i := sd.findStadistic(s.ID())
		sd.Stadistics[i].ReceiptTime = time.Now().UTC().Format("2006-01-02 15:04:05Z")
		sd.Stadistics[i].Status = "SUCCESS"
		server.BroadcastToRoom("/", "app", "stadistics", sd.Stadistics)
	})

	// event to get app connection
	server.OnEvent("/", "app", func(s socketio.Conn, msg string) {
		log.Println(msg)

		// join client to app room
		server.JoinRoom("/", "app", s)
		server.BroadcastToRoom("/", "app", "connections", sd.Connections)
		server.BroadcastToRoom("/", "app", "stadistics", sd.Stadistics)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		i := sd.findConnection(s.ID())
		if i != -1 {
			log.Printf("client %s disconnected", s.ID())
			sd.Connections = sd.removeConnection(s.ID())
			server.BroadcastToRoom("/", "app", "connections", sd.Connections)
		}
	})

	return server
}

func (sd ServerData) findConnection(clientID string) int {
	for i, connection := range sd.Connections {
		if connection.ClientID == clientID {
			return i
		}
	}
	return -1
}

func (sd ServerData) findStadistic(clientID string) int {
	for i, stadistic := range sd.Stadistics {
		if (stadistic.ToClient == clientID) && stadistic.Status != "SUCCESS" {
			return i
		}
	}
	return -1
}

func (sd ServerData) findStadisticChannel(channelID string) []int {
	var found []int
	for i, stadistic := range sd.Stadistics {
		if stadistic.ToChannel == channelID && sd.findConnection(stadistic.ToClient) != -1 {
			found = append(found, i)
		}
	}

	return found
}

func (sd ServerData) removeConnection(clientID string) []models.Connection {
	i := sd.findConnection(clientID)
	if i != -1 {
		return append(sd.Connections[:i], sd.Connections[i+1:]...)
	}

	return sd.Connections
}
