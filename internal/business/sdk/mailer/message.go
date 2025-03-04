package mailer

// Message is the data structure for sending emails.
type Message struct {
	To      []string
	From    string
	Subject string
	Body    string
}
