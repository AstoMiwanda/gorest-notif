package client

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"

	"github.com/pushm0v/gorest-notif/model"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridClient struct {
	mailClient *mail.SGMailV3
}

func NewSendgridClient() Client {
	m := mail.NewV3Mail()

	return &sendgridClient{
		mailClient: m,
	}
}

func (c *sendgridClient) constructMessage(m *model.Message) {
	to := mail.NewEmail(m.DestinationName, m.Destination)
	from := mail.NewEmail("No Reply", "no-reply@kitabisa.com")
	content := mail.NewContent("text/html", fmt.Sprintf("<p>%s</p>", m.Body))

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = m.Subject

	c.mailClient.SetFrom(from)
	c.mailClient.AddContent(content)
	c.mailClient.AddPersonalizations(personalization)
}

func (c *sendgridClient) Send(m *model.Message) error {
	c.constructMessage(m)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(c.mailClient)
	_, err := sendgrid.API(request)

	return err
}
