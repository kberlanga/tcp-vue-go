package agent

import (
	"errors"
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

func SendData(data SharedData) error {

	client, err := socketio_client.NewClient(URI, GetOpts())
	if err != nil {
		return err
	}

	contentFile, err := readFile(data.Filepath)
	if err != nil {
		return err
	}

	data.Data = contentFile
	data.Filename = getFilename(data.Filepath)

	client.Emit("send", data)
	log.Printf("client is sending data by channel: %v...\n", data.Channel)

	return nil
}

func readFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error reading file")
	}

	return data, nil
}

func getFilename(filepath string) string {
	splitted := strings.Split(filepath, "/")
	return splitted[len(splitted)-1]
}
