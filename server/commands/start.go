package commands

import (
	"flag"

	agent "github.com/kberlanga/tcp-vue-go/server/agent"
)

type StartCommand struct {
	fs *flag.FlagSet
}

func NewStartCommand() *StartCommand {
	gc := &StartCommand{
		fs: flag.NewFlagSet("start", flag.ContinueOnError),
	}

	return gc
}

func (g *StartCommand) Name() string {
	return g.fs.Name()
}

func (g *StartCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *StartCommand) Run() error {
	agent.InitServer()
	return nil
}
