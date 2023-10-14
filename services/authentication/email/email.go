package email

import (
	"net/smtp"

	"github.com/spf13/viper"
)

type Email struct {
	From     string
	Password string
	To       []string
	Text     string
}

func (s *Email) Send() error {
	smtpHost := viper.Get("SMTPHOST").(string)
	smtpPort := viper.Get("SMTPPORT").(string)

	auth := smtp.PlainAuth("", s.From, s.Password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, s.From, s.To, []byte(s.Text))
	if err != nil {
		return err
	}
	return nil
}
