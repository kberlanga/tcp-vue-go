package agent

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFilename(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one level string",
			args: args{
				filepath: "./test.txt",
			},
			want: "test.txt",
		},
		{
			name: "two level string",
			args: args{
				filepath: "./other/path/file.txt",
			},
			want: "file.txt",
		},
		{
			name: "three level string",
			args: args{
				filepath: "./other/path/three-level/three_level.txt",
			},
			want: "three_level.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFilename(tt.args.filepath)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want []byte
		err  error
	}{
		{
			name: "read file sucess",
			args: args{
				"../test.txt",
			},
			want: []byte("Message from file: Hello, world!"),
		},
		{
			name: "read file error",
			args: args{
				filepath: "./fakepath/test.txt",
			},
			want: nil,
			err:  errors.New("error reading file"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.filepath)
			log.Println(err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSendData(t *testing.T) {
	type args struct {
		data SharedData
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "send data success",
			args: args{
				data: SharedData{
					Channel:  "1",
					Filepath: "../test.txt",
					Data:     []byte("Hello, World!"),
					Filename: "test.txt",
				},
			},
			err: nil,
		},
		{
			name: "send data error reading file",
			args: args{
				data: SharedData{
					Channel:  "1",
					Filepath: "./test.txt",
					Data:     []byte("Hello, World!"),
					Filename: "test.txt",
				},
			},
			err: errors.New("error reading file"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SendData(tt.args.data)
			assert.Equal(t, tt.err, err)
		})
	}
}
