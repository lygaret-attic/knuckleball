package knuckleball

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
	The handle to the connected client that's passed into the handler of the
	server. It's a line oriented interface.
*/
type IClient interface {
	Readln() (s *string, err error)
	Writeln(s string) (err error)
}

/*
	Starts a server at the target address (http://localhost:2345) that will
	handle connections by calling the handler in a go routine. The client
	in the handler is an IClient, and has method to read and write in a line
	oriented fashion.

	Returns an error if there was one, or nil otherwise
*/
func Listen(target string, handler func(client IClient)) error {

	// resolve the address
	addr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		return err
	}

	// listen on that TCP address
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	// accept and handle connections
	for {
		connection, err := listener.AcceptTCP()
		if err != nil {
			return err
		}

		go func() {
			log.Println("Accepted incoming connection: ", connection.RemoteAddr(), " -> ", connection.LocalAddr())

			client := sClient{connection, bufio.NewReader(connection), false}
			defer client.close()

			handler(client)
		}()
	}

	return nil
}

func (client sClient) Readln() (*string, error) {

	// read the line from the buffer
	line, err := client.reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading from ", client.connection.RemoteAddr(), " : ", err)
		client.close()
	}

	return &line, err
}

func (client sClient) Writeln(line string) error {

	// write to the connection stream
	_, err := fmt.Fprintln(client.connection, line)
	if err != nil {
		log.Println("Error writing to ", client.connection.RemoteAddr(), " : ", err)
		client.close()
	}

	return err
}

func (client sClient) close() {
	client.connection.Close()
	client.closed = true
}

// private

type sClient struct {
	connection *net.TCPConn
	reader     *bufio.Reader
	closed     bool
}
