package commands

import (
	"flag"

	"github.com/kberlanga/tcp-vue-go/client/agent"
)

type SendCommand struct {
	fs *flag.FlagSet

	channel  string
	filepath string
}

func NewSendCommand() *SendCommand {
	gc := &SendCommand{
		fs: flag.NewFlagSet("send", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.filepath, "filepath", "", "filepath of file to send")
	gc.fs.StringVar(&gc.channel, "channel", "1", "channel number to pass and receive files")

	return gc
}

func (g *SendCommand) Name() string {
	return g.fs.Name()
}

func (g *SendCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *SendCommand) Run() error {
	data := agent.SharedData{
		Channel:  g.channel,
		Filepath: g.filepath,
	}

	return agent.SendData(data)
}
