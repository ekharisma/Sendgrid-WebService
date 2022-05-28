package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ekharisma/sendgrid-web-service/internals/service"
)

type IController interface {
	SendEmail(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	EmailService service.IEmailClient
}

func NewSendGridController(emailService service.IEmailClient) IController {
	return &Controller{
		EmailService: emailService,
	}
}

func (c *Controller) SendEmail(w http.ResponseWriter, r *http.Request) {
	var email service.Email
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	mail := c.EmailService.Generate(&email)
	isSent, err := c.EmailService.Send(mail)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isSent {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to send email"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Email sent"))
}
