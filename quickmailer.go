package main

import (
	"flag"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	server := flag.String("server", "", "smtp server without port to use")
	port := flag.String("port", "587", "smtp port to use")
	passwd := flag.String("passwd", "", "password for smtp auth")
	from := flag.String("from", "", "email to send from")
	to := flag.String("to", "", "email to send to")
	subject := flag.String("subject", "Your Files", "subject line")
	body := flag.String("body", "See attached", "some body text")
	attachFile := flag.String("attach", "", "single file to attach")
	attachDir := flag.String("attachDir", "", "directory to attach files from")

	if (*server == "") || (*port == "") || (*passwd == "") || (*from == "") || (*to == "") {
		panic("Required flags: -server, -passwd, -port, -from, -to")
	}

	sendserver := *server + ":" + *port

	e := email.NewEmail()
	e.From = "<" + *from + ">"
	e.To = []string{*to}
	e.Subject = *subject
	e.Text = []byte(*body)

	if *attachFile != "" {
		e.AttachFile(*attachFile)
	}

	if *attachDir != "" {

	}

	e.Send(sendserver, smtp.PlainAuth("", *from, *passwd, *server))
}
