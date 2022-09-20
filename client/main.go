package main

import (
	"log"
	"os"

	"github.com/kberlanga/tcp-vue-go/client/commands"
)

func main() {

	if err := commands.Root(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
