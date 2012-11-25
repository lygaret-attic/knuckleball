package knuckleball

import (
	netmail "net/mail"
)

type Message struct {
	Sender     string
	Recipients []string
	Message    netmail.Message
}
