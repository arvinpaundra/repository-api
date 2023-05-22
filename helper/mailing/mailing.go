package mailing

import (
	"strconv"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/helper"
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

type User struct {
	Fullname       string
	Email          string
	IdentityNumber string
	Departement    string
	StudyProgram   string
	Role           string
	YearGen        string
}

type Repository struct {
	Title         string
	Authors       string
	Collection    string
	Category      string
	Departement   string
	DateValidated string
}

// NewMailing returns a new Mailing struct
func NewMailing(host, port, email, password, senderName string) *Mailing {
	return &Mailing{
		host:       host,
		port:       port,
		email:      email,
		password:   password,
		senderName: senderName,
	}
}

func (mc *Mailing) SendForgotPasswordMail(recipient, subject, token string) error {
	data := map[string]interface{}{
		"token":    token,
		"email":    recipient,
		"base_url": configs.GetConfig("FE_BASE_URL"),
	}

	result, err := helper.RenderHTMLToString(templates.ForgotPassword, nil, data)

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

func (mc *Mailing) SendVerifiedRegisterMail(subject string, user User) error {
	data := map[string]interface{}{
		"fullname":        user.Fullname,
		"email":           user.Email,
		"identity_number": user.IdentityNumber,
		"departement":     user.Departement,
		"study_program":   user.StudyProgram,
		"role":            user.Role,
		"year_gen":        user.YearGen,
		"base_url":        configs.GetConfig("FE_BASE_URL"),
	}

	result, err := helper.RenderHTMLToString(templates.VerifiedRegister, nil, data)

	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", mc.senderName)
	mail.SetHeader("To", user.Email)
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

func (mc *Mailing) SendDeniedRegisterMail(recepient, subject, reasons string) error {
	data := map[string]interface{}{
		"reasons":  reasons,
		"base_url": configs.GetConfig("FE_BASE_URL"),
	}

	result, err := helper.RenderHTMLToString(templates.DeniedRegister, nil, data)

	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", mc.senderName)
	mail.SetHeader("To", recepient)
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

func (mc *Mailing) SendVerifiedRepositoryMail(recepient, subject string, repository Repository) error {
	data := map[string]interface{}{
		"title":          repository.Title,
		"authors":        repository.Authors,
		"collection":     repository.Collection,
		"category":       repository.Category,
		"departement":    repository.Departement,
		"date_validated": repository.DateValidated,
		"base_url":       configs.GetConfig("FE_BASE_URL"),
	}

	result, err := helper.RenderHTMLToString(templates.VerifiedRepository, nil, data)

	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", mc.senderName)
	mail.SetHeader("To", recepient)
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

func (mc *Mailing) SendDeniedRepositoryMail(recepient, subject string, repository Repository) error {
	data := map[string]interface{}{
		"title":          repository.Title,
		"authors":        repository.Authors,
		"collection":     repository.Collection,
		"category":       repository.Category,
		"departement":    repository.Departement,
		"date_validated": repository.DateValidated,
		"base_url":       configs.GetConfig("FE_BASE_URL"),
	}

	result, err := helper.RenderHTMLToString(templates.DeniedRepository, nil, data)

	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", mc.senderName)
	mail.SetHeader("To", recepient)
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
