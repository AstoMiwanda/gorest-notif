package service

import (
	"github.com/pushm0v/gorest-notif/client"
	"github.com/pushm0v/gorest-notif/model"
)

type NotifService interface {
	SendEmail(m *model.Message) (err error)
}

type notifService struct {
	emailClient client.Client
}

func NewNotifService(emailClient client.Client) NotifService {
	return &notifService{
		emailClient: emailClient,
	}
}

func (n *notifService) SendEmail(m *model.Message) (err error) {
	return n.emailClient.Send(m)
}
