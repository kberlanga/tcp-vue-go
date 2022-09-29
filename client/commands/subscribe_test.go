package commands

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribeCommand_Name(t *testing.T) {
	type fields struct {
		fs      *flag.FlagSet
		channel string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get name",
			fields: fields{
				fs: flag.NewFlagSet("receive", flag.ContinueOnError),
			},
			want: "receive",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &SubscribeCommand{
				fs:      tt.fields.fs,
				channel: tt.fields.channel,
			}
			if got := g.Name(); got != tt.want {
				t.Errorf("SubscribeCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribeCommand_Init(t *testing.T) {
	type fields struct {
		fs      *flag.FlagSet
		channel string
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "init",
			fields: fields{
				fs: flag.NewFlagSet("receive", flag.ContinueOnError),
			},
			args: args{
				args: []string{"1"},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &SubscribeCommand{
				fs:      tt.fields.fs,
				channel: tt.fields.channel,
			}
			err := g.Init(tt.args.args)
			assert.Equal(t, tt.err, err)
		})
	}
}
