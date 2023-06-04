package data

import (
	"net/smtp"

	"github.com/hashicorp/go-hclog"
)

const smtpAddr = "smtp.gmail.com"

type MailServer struct {
	log        hclog.Logger
	DevAddress []byte
}

func NewMailService(log hclog.Logger, Address []byte) *MailServer {
	msrv := &MailServer{log: log, DevAddress: Address}
	return msrv
}

func (m *MailServer) SendMail(message []byte) error {
	auth := smtp.PlainAuth(
		"",
		string(m.DevAddress),
		"", // AppPassword
		smtpAddr,
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		string(m.DevAddress),
		[]string{string(m.DevAddress)},
		message,
	)

	if err != nil {
		m.log.Error("Unable to send a mail", "error", err)
		return err
	}

	return nil
}
