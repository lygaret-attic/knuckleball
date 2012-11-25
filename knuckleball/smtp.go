package knuckleball

import (
	"fmt"
)

type handlerf func(line *string, context *context) (handlerf, error)

/*
	Starts a server at the target address (http://localhost:2345) that will
	handle smtp sessions, and pass incoming messages back to the caller via
	the passed in channel.

	Returns an error if there was one, or nil otherwise
*/
func ListenSMTP(target string, messages chan Message) error {

	Listen(target, func(client IClient) {
		client.Writeln("220 knuckleball")

		context := context{client: client, messages: messages}

		// state machine!
		state := helo
		for state != nil {
			line, err := client.Readln()
			if err != nil {
				fmt.Printf("Error reading from client! %v\n", err)
				break
			}

			state, err = state(line, &context)
			if err != nil {
				fmt.Println("Error!", err)
				break
			}
		}
	})

	return nil
}
