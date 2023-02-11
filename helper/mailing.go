package helper

import (
	"strconv"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/templates"
	"gopkg.in/gomail.v2"
)

type Mailing struct {
	host       string
	port       string
	senderName string
	email      string
	password   string
}

func NewMailing() *Mailing {
	return &Mailing{
		host:       configs.GetConfig("SMTP_HOST"),
		port:       configs.GetConfig("SMTP_PORT"),
		email:      configs.GetConfig("EMAIL"),
		password:   configs.GetConfig("PASSWORD_EMAIL"),
		senderName: configs.GetConfig("EMAIL_SENDER_NAME"),
	}
}

func (mc *Mailing) SendForgotPasswordMail(recipient, subject, token string) error {
	data := map[string]string{
		"token":    token,
		"email":    recipient,
		"base_url": configs.GetConfig("FE_BASE_URL"),
	}

	result, err := RenderHTMLToString(templates.ForgotPassword, data)

	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", mc.senderName)
	mail.SetHeader("To", recipient)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", result)

	port, _ := strconv.Atoi(mc.port)

	dialer := gomail.NewDialer(mc.host, port, mc.email, mc.password)

	err = dialer.DialAndSend(mail)

	if err != nil {
		return err
	}

	return nil
}
