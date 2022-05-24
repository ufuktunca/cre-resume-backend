package email

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

const mail = "cre-resume@yandex.com"

type EmailInterface interface {
	SendMail(registeredUser, message string) error
}

func SendMail(registeredUser, message string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mail)
	m.SetHeader("To", registeredUser)
	m.SetHeader("Subject", "Register verificiation")

	m.SetBody("text/plain", message)
	d := gomail.NewDialer("smtp.gmail.com", 465, mail, "vqeqebnvtfckvmnm")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)

	return err
}
