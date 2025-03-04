package smtp

import (
	"context"
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"

	"github.com/machilan1/plpr2/internal/business/sdk/mailer"
)

type Config struct {
	From string
	// Host is the SMTP server host.
	Host string
	// Port is the SMTP server port.
	Port int
	// User is the SMTP server user.
	User string
	// Pass is the SMTP server password.
	Pass string
}

// smtpMailer is a Mailer implementation that sends emails using SMTP.
type smtpMailer struct {
	from string
	addr string
	auth smtp.Auth
}

func New(cfg Config) mailer.Mailer {
	return &smtpMailer{
		from: cfg.From,
		addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		auth: smtp.CRAMMD5Auth(cfg.User, cfg.Pass),
	}
}

func (sc *smtpMailer) Send(_ context.Context, messages ...mailer.Message) (int, error) {
	sentEmailsCount := 0

	for _, msg := range messages {
		err := sc.sendMessage(msg)
		if err != nil {
			return sentEmailsCount, err
		}

		sentEmailsCount++
	}

	return sentEmailsCount, nil
}

func (sc *smtpMailer) sendMessage(msg mailer.Message) error {
	m := sc.buildEmail(msg)
	return smtp.SendMail(sc.addr, sc.auth, sc.from, msg.To, m)
}

// buildEmail builds an email message in the HTML mime format.
func (sc *smtpMailer) buildEmail(msg mailer.Message) []byte {
	header := make(map[string]any)
	header["From"] = sc.from
	header["To"] = msg.To
	header["Subject"] = mime.QEncoding.Encode("UTF-8", msg.Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	m := ""
	for k, v := range header {
		m += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	m += "\r\n" + base64.StdEncoding.EncodeToString([]byte(msg.Body))

	return []byte(m)
}
