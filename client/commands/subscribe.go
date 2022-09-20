package commands

import (
	"flag"

	"github.com/kberlanga/tcp-vue-go/client/agent"
)

type SubscribeCommand struct {
	fs *flag.FlagSet

	channel string
}

func NewSubscribeCommand() *SubscribeCommand {
	gc := &SubscribeCommand{
		fs: flag.NewFlagSet("receive", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.channel, "channel", "1", "channel number to pass and receive files")

	return gc
}

func (g *SubscribeCommand) Name() string {
	return g.fs.Name()
}

func (g *SubscribeCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *SubscribeCommand) Run() error {
	agent.SubscribeClient(g.channel)
	return nil
}
