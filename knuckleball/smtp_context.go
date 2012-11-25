package knuckleball

import (
	"fmt"
)

type context struct {
	client     IClient
	clientName string
	curMessage Message
	messages   chan Message
}

func (context *context) send_ok() error {
	return context.send_reply(250)
}

func (context *context) send_reply(code int) error {
	return context.client.Writeln(fmt.Sprintf("%v: %v", code, get_message(code)))
}

func get_message(code int) string {
	switch code {
	case 211:
		return "System status, or system help reply"
	case 214:
		return "Help message [this reply is useful only to the human user]"
	case 220:
		return "<domain> Service ready"
	case 221:
		return "<domain> Service closing transmission channel"
	case 250:
		return "Ok."
	case 251:
		return "User not local; will forward to <forward-path>"
	case 252:
		return "Cannot VRFY user, but will accept message and attempt delivery"

	case 354:
		return "Start mail input; end with <CRLF>.<CRLF>"

	case 421:
		return "<domain> Service not available, closing transmission channel"
	case 450:
		return "Requested mail action not taken: mailbox unavailable"
	case 451:
		return "Requested action aborted: local error in processing"
	case 452:
		return "Requested action not taken: insufficient system storage"

	case 500:
		return "Syntax error, command unrecognized"
	case 501:
		return "Syntax error in parameters or arguments"
	case 502:
		return "Command not implemented"
	case 503:
		return "Bad sequence of commands"
	case 504:
		return "Command parameter not implemented"
	case 550:
		return "Requested action not taken: mailbox unavailable"
	case 551:
		return "User not local; please try <forward-path>"
	case 552:
		return "Requested mail action aborted: exceeded storage allocation"
	case 553:
		return "Requested action not taken: mailbox name not allowed"
	case 554:
		return "Transaction failed"
	}

	return "Unknown Status"
}
