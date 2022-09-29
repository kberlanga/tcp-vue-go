package agent

import (
	"log"
	"net/http"

	"github.com/kberlanga/tcp-vue-go/server/agent/models"
)

type ServerData struct {
	Connections []models.Connection
	Stadistics  []models.ShippingStatistics
}

func InitServer() {

	serverData := ServerData{
		Connections: []models.Connection{},
		Stadistics:  []models.ShippingStatistics{},
	}

	server := serverData.setupServer()

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("../app/dist")))
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
