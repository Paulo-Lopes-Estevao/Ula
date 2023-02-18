package email

import (
	"bytes"
	"crypto/tls"
	"errors"

	"gopkg.in/gomail.v2"
)

type Email struct {
	From             string
	To               string
	Cc               string
	Subject          string
	Body             string
	IEmailCredential EmailCredentialInterface
}

type IEmail interface {
	NewMessage() *gomail.Message
	DialPoolConnection() *gomail.Dialer
	SendEmailPlain()
	SendEmailHtml()
}

const (
	ErrTo      = "to is required"
	ErrSubject = "subject is required"
	ErrBody    = "body is required"
)

func NewEmail(to, subject, body string, iEmailCredential EmailCredentialInterface) (*Email, error) {
	e := &Email{
		To:               to,
		Subject:          subject,
		Body:             body,
		IEmailCredential: iEmailCredential,
	}

	if err := e.validate(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e Email) NewMessage() *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", e.IEmailCredential.GetFrom())
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", e.Subject)

	return m
}

func (e Email) DialPoolConnection() *gomail.Dialer {
	d := gomail.NewDialer(e.IEmailCredential.GetHost(), e.IEmailCredential.GetPort(), e.IEmailCredential.GetFrom(), e.IEmailCredential.GetPassword())

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d

}

func (e Email) SendEmailPlain() {

	m := e.NewMessage()
	m.SetBody("text/plain", e.Body)

	dialPool := e.DialPoolConnection()

	if err := dialPool.DialAndSend(m); err != nil {
		panic(err)
	}
}

func (e Email) SendEmailHtml() {

	m := e.NewMessage()
	m.SetBody("text/html", e.Body)

	dialPool := e.DialPoolConnection()

	if err := dialPool.DialAndSend(m); err != nil {
		panic(err)
	}
}

func (e Email) SendEmailTemplateHtml() {
	var body bytes.Buffer

	if err := TemplateFileName(&body, e.Body); err != nil {
		panic(err)
	}

	m := e.NewMessage()
	m.SetBody("text/html", body.String())

	defer body.Reset()

	dialPool := e.DialPoolConnection()

	if err := dialPool.DialAndSend(m); err != nil {
		panic(err)
	}

}

func (e Email) validate() error {
	if e.To == "" {
		return errors.New(ErrTo)
	}

	if e.Subject == "" {
		return errors.New(ErrSubject)
	}

	if e.Body == "" {
		return errors.New(ErrBody)
	}

	return nil
}
