package sendgrid

import (
	"context"
	"fmt"

	"github.com/machilan1/plpr2/internal/business/sdk/mailer"
	"github.com/machilan1/plpr2/internal/framework/logger"
	"github.com/machilan1/plpr2/internal/framework/tracer"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.opentelemetry.io/otel/attribute"
)

type Config struct {
	APIKey      string
	FromName    string
	FromAddress string
}

type sendgridMailer struct {
	log         *logger.Logger
	sgClient    *sendgrid.Client
	fromName    string
	fromAddress string
}

func New(log *logger.Logger, cfg Config) mailer.Mailer {
	c := sendgrid.NewSendClient(cfg.APIKey)

	return &sendgridMailer{
		log:         log,
		sgClient:    c,
		fromName:    cfg.FromName,
		fromAddress: cfg.FromAddress,
	}
}

func (sg *sendgridMailer) Send(ctx context.Context, messages ...mailer.Message) (int, error) {
	ctx, span := tracer.AddSpan(ctx, "business.mailer.sendgrid.send", attribute.Int("messages", len(messages)))
	defer span.End()

	sentEmailsCount := 0

	for _, msg := range messages {
		err := sg.sendMessage(ctx, msg)
		if err != nil {
			return sentEmailsCount, err
		}

		sentEmailsCount++
	}

	return sentEmailsCount, nil
}

func (sg *sendgridMailer) sendMessage(ctx context.Context, msg mailer.Message) error {
	from := mail.NewEmail(sg.fromName, sg.fromAddress)
	subject := msg.Subject

	m := mail.NewV3Mail()
	m.SetFrom(from)
	m.Subject = subject

	content := mail.NewContent("text/html", msg.Body)
	m.AddContent(content)

	p := mail.NewPersonalization()
	for _, to := range msg.To {
		p.AddTos(mail.NewEmail("", to))
	}
	m.AddPersonalizations(p)

	response, err := sg.sgClient.SendWithContext(ctx, m)
	if err != nil {
		return fmt.Errorf("sendgrid: client error: %w", err)
	}

	sg.log.Info(ctx, "sendgrid client send", "status", response.StatusCode, "body", response.Body, "headers", response.Headers)

	if response.StatusCode >= 400 {
		return fmt.Errorf("sendgrid: send email failed: %s", response.Body)
	}

	return nil
}
