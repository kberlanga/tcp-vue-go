package agent

import (
	"testing"

	"github.com/kberlanga/tcp-vue-go/server/agent/testcases"
	"github.com/stretchr/testify/assert"
)

func TestServerData_findConnection(t *testing.T) {
	for _, tt := range testcases.FindConnection_tests {
		t.Run(tt.Name, func(t *testing.T) {
			sd := ServerData{
				Connections: tt.Fields.Connections,
				Stadistics:  tt.Fields.Stadistics,
			}
			got := sd.findConnection(tt.Args.ClientID)
			assert.Equal(t, tt.Want, got)
		})
	}
}

func TestServerData_findStadistic(t *testing.T) {
	for _, tt := range testcases.FindStadistic_tests {
		t.Run(tt.Name, func(t *testing.T) {
			sd := ServerData{
				Connections: tt.Fields.Connections,
				Stadistics:  tt.Fields.Stadistics,
			}
			got := sd.findStadistic(tt.Args.ClientID)
			assert.Equal(t, tt.Want, got)
		})
	}
}

func TestServerData_findStadisticChannel(t *testing.T) {
	for _, tt := range testcases.FindStadisticChannel_tests {
		t.Run(tt.Name, func(t *testing.T) {
			sd := ServerData{
				Connections: tt.Fields.Connections,
				Stadistics:  tt.Fields.Stadistics,
			}
			got := sd.findStadisticChannel(tt.Args.ChannelID)
			assert.Equal(t, tt.Want, got)
		})
	}
}

func TestServerData_removeConnection(t *testing.T) {
	for _, tt := range testcases.RemoveConnection_tests {
		t.Run(tt.Name, func(t *testing.T) {
			sd := ServerData{
				Connections: tt.Fields.Connections,
				Stadistics:  tt.Fields.Stadistics,
			}
			got := sd.removeConnection(tt.Args.ClientID)
			assert.Equal(t, tt.Want, got)
		})
	}
}
