package email

import (
	"gopkg.in/gomail.v2"
)

type Client struct {
	Sender sender
}

type sender struct {
	smpt     string
	port     int
	username string
	password string
}

type IEmail interface {
	CreateMessage(email string, message string) *gomail.Message
	SendMail(m *gomail.Message) error
}

// NewClient creates a new client configured with the given sender options
func NewClient(smtp string, port int, username string, password string) *Client {
	sndr := &sender{
		smpt:     smtp,
		port:     port,
		username: username,
		password: password,
	}
	client := &Client{Sender: *sndr}
	return client
}

func (c *Client) CreateMessage(email string, message string) *gomail.Message {
	m := gomail.NewMessage()

	// from
	m.SetHeader("From", c.Sender.username)

	// to
	m.SetHeader("To", email)

	// subject
	m.SetHeader("Subject", "Your One Time Password")

	m.SetBody("text/plain", message)

	return m
}

func (c *Client) SendMail(m *gomail.Message) error {
	// gmail smtp
	d := gomail.NewDialer(c.Sender.smpt, c.Sender.port, c.Sender.username, c.Sender.password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
