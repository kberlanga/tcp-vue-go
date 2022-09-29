package models

type Connection struct {
	ClientID         string `json:"client"`
	Channel          string `json:"channel"`
	SubscriptionTime string `json:"subscriptionTime"`
}

type ShippingStatistics struct {
	ShippingTime string `json:"shippingTime"`
	ReceiptTime  string `json:"receiptTime"`
	ToChannel    string `json:"toChannel"`
	ToClient     string `json:"toClient"`
	Status       string `json:"status"`
}
