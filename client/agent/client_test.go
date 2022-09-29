package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func TestGetOpts(t *testing.T) {
	tests := []struct {
		name string
		want *socketio_client.Options
	}{
		{
			name: "get options",
			want: &socketio_client.Options{
				Transport: "websocket",
				Query:     make(map[string]string),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOpts()
			assert.Equal(t, tt.want, got)
		})
	}
}
