package service

import (
	"fmt"
	"net/http"

	"github.com/ekharisma/sendgrid-web-service/internals/static"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type IEmailClient interface {
	Generate(email *Email) *mail.SGMailV3
	Send(mail *mail.SGMailV3) (bool, error)
}

type EmailClient struct {
	Config *static.Config
}

type Email struct {
	SenderName   string `json:"senderName"`
	SenderMail   string `json:"senderMail"`
	Subject      string `json:"subject"`
	ReceiverName string `json:"receiverName"`
	ReceiverMail string `json:"receiverMail"`
	Content      string `json:"content"`
}

func NewEmailClient(config *static.Config) IEmailClient {
	return &EmailClient{Config: config}
}

func (s *EmailClient) Generate(email *Email) *mail.SGMailV3 {
	from := mail.NewEmail(email.SenderName, email.SenderMail)
	subject := email.Subject
	content := email.Content
	to := mail.NewEmail(email.ReceiverName, email.ReceiverMail)
	return mail.NewSingleEmail(from, subject, to, content, content)
}

func (s *EmailClient) Send(mail *mail.SGMailV3) (bool, error) {
	client := sendgrid.NewSendClient(s.Config.Key)
	response, err := client.Send(mail)
	fmt.Println("Code: ", response.StatusCode)
	fmt.Println("Body: ", response.Body)
	if err != nil {
		return false, err
	}
	return response.StatusCode == http.StatusAccepted, nil
}
