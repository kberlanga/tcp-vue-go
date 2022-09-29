package commands

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartCommand_Name(t *testing.T) {
	type fields struct {
		fs *flag.FlagSet
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get name",
			fields: fields{
				fs: flag.NewFlagSet("start", flag.ContinueOnError),
			},
			want: "start",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &StartCommand{
				fs: tt.fields.fs,
			}
			got := g.Name()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStartCommand_Init(t *testing.T) {
	type fields struct {
		fs *flag.FlagSet
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
				fs: flag.NewFlagSet("start", flag.ContinueOnError),
			},
			args: args{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &StartCommand{
				fs: tt.fields.fs,
			}
			err := g.Init(tt.args.args)
			assert.Equal(t, tt.err, err)
		})
	}
}
