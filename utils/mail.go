package utils

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpGmail     = "smtp.gmail.com"
	smtpAddServer = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		Subject string,
		Content string,
		To []string,
		Bcc []string,
		Cc []string,
		AttachedFiles []string,
	) error
}

type GmailSender struct {
	Address  string
	Password string
	AppName  string
}

func NewGmailSender(address string, password string, appName string) EmailSender {
	return &GmailSender{Address: address, Password: password, AppName: appName}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	bcc []string,
	cc []string,
	attachedFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.AppName, sender.Address)
	e.HTML = []byte(content)
	e.To = to
	e.Bcc = bcc
	e.Cc = cc
	e.Subject = subject

	for _, file := range attachedFiles {
		_, err := e.AttachFile(file)
		if err != nil {
			return fmt.Errorf("FAILED to  attachfile %s", file)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.Address, sender.Password, smtpGmail)

	return e.Send(smtpAddServer, smtpAuth)
}
