package agent

import (
	"bufio"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func SubscribeClient(channel string) {

	client, err := socketio_client.NewClient(URI, GetOpts())
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
	}

	client.On("reply", func(data SharedData) {
		log.Printf("receiving data\nfrom channel: %s\ncontent-file: %s", data.Channel, string(data.Data))
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		ch := channel
		client.Emit("subscription", ch)
		log.Printf("client is subscribed to channel: %v\n", ch)
		reader.ReadLine()
	}

}
