package testcases

import (
	"time"

	"github.com/kberlanga/tcp-vue-go/server/agent/models"
)

type Fields struct {
	Connections []models.Connection
	Stadistics  []models.ShippingStatistics
}

type Args struct {
	ClientID string
}

var FindConnection_tests = []struct {
	Name   string
	Fields Fields
	Args   Args
	Want   int
}{
	{
		Name: "find connection",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "2",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "3",
					Channel:          "3",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
		},
		Args: Args{
			ClientID: "1",
		},
		Want: 0,
	},
	{
		Name: "not find connection",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "2",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "3",
					Channel:          "3",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
			Stadistics: []models.ShippingStatistics{},
		},
		Args: Args{
			ClientID: "4",
		},
		Want: -1,
	},
	{
		Name: "empty connections",
		Fields: Fields{
			Connections: []models.Connection{},
			Stadistics:  []models.ShippingStatistics{},
		},
		Args: Args{
			ClientID: "1",
		},
		Want: -1,
	},
}

var FindStadistic_tests = []struct {
	Name   string
	Fields Fields
	Args   Args
	Want   int
}{
	{
		Name: "find connection",
		Fields: Fields{
			Stadistics: []models.ShippingStatistics{
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "1",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "2",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "2",
					ToClient:     "3",
					Status:       "RECEIVING",
				},
			},
		},
		Args: Args{
			ClientID: "1",
		},
		Want: 0,
	},
	{
		Name: "not find connection",
		Fields: Fields{
			Connections: []models.Connection{},
			Stadistics: []models.ShippingStatistics{
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "1",
					Status:       "SUCCESS",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "2",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "2",
					ToClient:     "3",
					Status:       "SUCCESS",
				},
			},
		},
		Args: Args{
			ClientID: "3",
		},
		Want: -1,
	},
	{
		Name: "empty connections",
		Fields: Fields{
			Connections: []models.Connection{},
			Stadistics:  []models.ShippingStatistics{},
		},
		Args: Args{
			ClientID: "1",
		},
		Want: -1,
	},
}

type FindStadisticChannelArgs struct {
	ChannelID string
}

var FindStadisticChannel_tests = []struct {
	Name   string
	Fields Fields
	Args   FindStadisticChannelArgs
	Want   []int
}{
	{
		Name: "more than 1 stadistic found",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "2",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "3",
					Channel:          "3",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
			Stadistics: []models.ShippingStatistics{
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "1",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "2",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "2",
					ToClient:     "3",
					Status:       "RECEIVING",
				},
			},
		},
		Args: FindStadisticChannelArgs{
			ChannelID: "1",
		},
		Want: []int{0, 1},
	},
	{
		Name: "one stadistic found",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "2",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "3",
					Channel:          "3",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
			Stadistics: []models.ShippingStatistics{
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "1",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "2",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "2",
					ToClient:     "3",
					Status:       "RECEIVING",
				},
			},
		},
		Args: FindStadisticChannelArgs{
			ChannelID: "2",
		},
		Want: []int{2},
	},
	{
		Name: "one stadistic match",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
			Stadistics: []models.ShippingStatistics{
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "1",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "1",
					ToClient:     "2",
					Status:       "RECEIVING",
				},
				{
					ShippingTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ReceiptTime:  time.Now().UTC().Format("2006-01-02 15:04:05Z"),
					ToChannel:    "2",
					ToClient:     "3",
					Status:       "RECEIVING",
				},
			},
		},
		Args: FindStadisticChannelArgs{
			ChannelID: "1",
		},
		Want: []int{0},
	},
	{
		Name: "one stadistic found",
		Fields: Fields{
			Connections: []models.Connection{},
			Stadistics:  []models.ShippingStatistics{},
		},
		Args: FindStadisticChannelArgs{
			ChannelID: "2",
		},
		Want: nil,
	},
}

var RemoveConnection_tests = []struct {
	Name   string
	Fields Fields
	Args   Args
	Want   []models.Connection
}{
	{
		Name: "connection removed",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "2",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "3",
					Channel:          "3",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
		},
		Args: Args{
			ClientID: "1",
		},
		Want: []models.Connection{
			{
				ClientID:         "2",
				Channel:          "1",
				SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
			},
			{
				ClientID:         "3",
				Channel:          "3",
				SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
			},
		},
	},
	{
		Name: "connection not found",
		Fields: Fields{
			Connections: []models.Connection{
				{
					ClientID:         "1",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "2",
					Channel:          "1",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
				{
					ClientID:         "3",
					Channel:          "3",
					SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
				},
			},
		},
		Args: Args{
			ClientID: "4",
		},
		Want: []models.Connection{
			{
				ClientID:         "1",
				Channel:          "1",
				SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
			},
			{
				ClientID:         "2",
				Channel:          "1",
				SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
			},
			{
				ClientID:         "3",
				Channel:          "3",
				SubscriptionTime: time.Now().UTC().Format("2006-01-02 15:04:05Z"),
			},
		},
	},
	{
		Name: "connections empty",
		Fields: Fields{
			Connections: []models.Connection{},
		},
		Args: Args{},
		Want: []models.Connection{},
	},
}
