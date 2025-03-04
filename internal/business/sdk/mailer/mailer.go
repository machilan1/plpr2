package mailer

import "context"

// Mailer is the interface that provides support for sending email messages.
type Mailer interface {
	// Send sends email messages and returns the number of messages sent.
	Send(ctx context.Context, messages ...Message) (int, error)
}
