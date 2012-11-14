package main

import (
	"github.com/lygaret/knuckleball"
	"log"
	"net/mail"
)

func main() {

	server := &knuckleball.Server{
		Addr:   ":25",
		Option: true,
		Other:  false,
	}

	server.HandleMail(func(message *mail.Message) {
		// do something
	})

	server.HandleError(func(error err) error {
		// bad error, bail
		return err

		// handlable, no big deal
		return nil
	})

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
