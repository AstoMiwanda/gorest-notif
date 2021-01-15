package client

import "github.com/pushm0v/gorest-notif/model"

type Client interface {
	Send(m *model.Message) error
}
