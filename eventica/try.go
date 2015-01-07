package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	email := "adoaonqj@sharklasers.com"
	auth := smtp.PlainAuth(
		"",
		"webadmin@techkriti.org",
		"rememberyourpassword",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"noreply@techkriti.org",
		[]string{email},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
