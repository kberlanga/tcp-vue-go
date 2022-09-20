package agent

import (
	"log"
	"os"
	"strings"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

type SharedData struct {
	Channel  string
	Filepath string
	Data     []byte
	Filename string
}

func SendData(data SharedData) {

	client, err := socketio_client.NewClient(URI, GetOpts())
	if err != nil {
		log.Printf("newClient error:%v\n", err)
	}

	contentFile, err := readFile(data.Filepath)
	if err != nil {
		log.Printf("readFile error:%v\n", err)
	}

	data.Data = contentFile
	data.Filename = getFilename(data.Filepath)

	client.Emit("send", data)
	log.Printf("client is sending data by channel: %v...\n", data.Channel)

}

func readFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getFilename(filepath string) string {
	splitted := strings.Split(filepath, "/")
	return splitted[len(splitted)-1]
}
