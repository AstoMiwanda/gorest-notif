package model

type MessageType string

const (
	Email MessageType = "Email"
)

type Message struct {
	Subject         string `json:"subject"`
	From            string `json:"from"`
	FromName        string `json:"from_name"`
	Destination     string `json:"destination"`
	DestinationName string `json:"destination_name"`
	Body            string `json:"body"`
	Type            MessageType
}
