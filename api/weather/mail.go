package weather

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
)

type IMailClient interface {
	SendMail(body string) error
}

type Config struct {
	sender   string
	receiver string
	password string
	smtphost string
	smtpport int
}

type MailClient struct {
	config *Config
}

var _ IMailClient = (*MailClient)(nil)

func NewMailClient(config Config) *MailClient {
	return &MailClient{
		config: &config,
	}
}

func (mc MailClient) SendMail(body string) error {
	server := mail.NewSMTPClient()
	server.Host = mc.config.smtphost
	server.Port = mc.config.smtpport
	server.Username = mc.config.sender
	server.Password = mc.config.password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom(mc.config.sender)
	email.AddTo(mc.config.receiver)
	email.SetSubject("Weather Report from Mantil!")

	email.SetBody(mail.TextPlain, body)

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
