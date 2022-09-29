package commands

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendCommand_Name(t *testing.T) {
	type fields struct {
		fs       *flag.FlagSet
		channel  string
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get name",
			fields: fields{
				fs:       flag.NewFlagSet("send", flag.ContinueOnError),
				channel:  "1",
				filepath: "./test.txt",
			},
			want: "send",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &SendCommand{
				fs:       tt.fields.fs,
				channel:  tt.fields.channel,
				filepath: tt.fields.filepath,
			}
			got := g.Name()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSendCommand_Init(t *testing.T) {
	type fields struct {
		fs       *flag.FlagSet
		channel  string
		filepath string
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
			name: "init success",
			fields: fields{
				fs:       flag.NewFlagSet("send", flag.ContinueOnError),
				channel:  "1",
				filepath: "./test.txt",
			},
			args: args{
				args: []string{"1", "./test.txt"},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &SendCommand{
				fs:       tt.fields.fs,
				channel:  tt.fields.channel,
				filepath: tt.fields.filepath,
			}
			err := g.Init(tt.args.args)
			assert.Equal(t, tt.err, err)
		})
	}
}
