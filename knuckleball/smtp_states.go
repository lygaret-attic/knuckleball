package knuckleball

import ()

/*

	Implement the minimum allowable command set for an SMTP server,
	according to the [RFC](1).

		4.5.1.  MINIMUM IMPLEMENTATION

			 In order to make SMTP workable, the following minimum
			 implementation is required for all receivers:

				COMMANDS -- HELO
							MAIL
							RCPT
							DATA
							RSET
							NOOP
							QUIT

	This corresponds to a state machine with four states:
		* hello
		  This is the initial state; helo, rset, noop and quit are allowed
		* mail
		  This occurs after hello, and is where rset sends you. it's the beginning 
		  of a message session.
		* rcpt
		  This state handles both the rcpt and initial data commands.
		* data
		  This state loops until the end of the message

	1: http://tools.ietf.org/html/rfc821#page-41
*/

/*
	Initial connection.
 */
func helo(line *string, context *context) (handlerf, error) {

	cmd, rest, err := read_cmd(line)
	if err != nil {
		return nil, err
	}

	// handle allowable commands
	if is_special(cmd) {
		return handle_special(cmd, rest, helo, context)
	}

	// special case ehlo
	if cmd == "EHLO" {
		return helo, context.send_reply(502)
	}

	// not helo, no idea who they are
	if cmd != "HELO" {
		return helo, context.send_reply(503)
	}

	context.clientName = rest
	return mail, context.send_ok()

}

/*
	Get the sender (the relay, really)
 */
func mail(line *string, context *context) (handlerf, error) {

	cmd, rest, err := read_cmd(line)
	if err != nil {
		return nil, err
	}

	if is_special(cmd) {
		return handle_special(cmd, rest, mail, context)
	}

	// not mail, bad sequence
	if cmd != "MAIL" {
		return mail, context.send_reply(503)
	}

	params := read_params(rest)
	if params == nil {
		return mail, context.send_reply(501)
	}

	// we're starting a new message now!
	context.curMessage = Message{
		Sender: params["FROM"],
	}

	return rcpt, context.send_ok()
}

/*
	Get the recipients. This repeats until DATA
 */
func rcpt(line *string, context *context) (handlerf, error) {

	cmd, rest, err := read_cmd(line)
	if err != nil {
		return nil, err
	}

	if cmd == "DATA" { 
		if len(context.curMessage.Recipients) > 0 {
		return data, context.send_reply(354)
	}
else {

}
}

}

func data(line *string, context *context) (handlerf, error) {

}


func is_special(cmd string) bool {
	return cmd == "RSET" || cmd == "NOOP" || cmd == "QUIT"
}

func handle_special(cmd string, rest string, handlerf current, context *context) (handlerf, error) {
	switch cmd {
	case "RSET":
		curMessage = nil
		return mail, context.send_ok()

	case "NOOP":
		return current, context.send_ok()

	case "QUIT":
		return nil, context.send_ok()
	}
}
