package main

import "github.com/tangingw/go_smtp"

func main() {

	sender := NewSender("tsuki4u@gmail.com", "Dumble#14dore11")

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"tsuki4u@gmail.com"}

	Subject := "Testing email from golang"
	bodyMessage := "Sending email using Golang. Yeah"

	sender.SendMail(Receiver, Subject, bodyMessage)
}
