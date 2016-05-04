package main

import (
	"flag"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	var (
		server     string
		port       string
		passwd     string
		from       string
		to         string
		subject    string
		body       string
		attachFile string
		attachDir  string
	)

	flag.StringVar(&server, "server", "", "smtp server without port to use")
	flag.StringVar(&port, "port", "587", "smtp port to use")
	flag.StringVar(&passwd, "passwd", "", "password for smtp auth")
	flag.StringVar(&from, "from", "", "email to send from")
	flag.StringVar(&to, "to", "", "email to send to")
	flag.StringVar(&subject, "subject", "Your Files", "subject line")
	flag.StringVar(&body, "body", "See attached", "some body text")
	flag.StringVar(&attachFile, "attach", "", "single file to attach")
	flag.StringVar(&attachDir, "attachDir", "", "directory to attach files from")

	flag.Parse()

	if (server == "") || (port == "") || (passwd == "") || (from == "") || (to == "") {
		panic("Required flags: -server, -passwd, -port, -from, -to")
	}

	sendserver := server + ":" + port

	e := email.NewEmail()
	e.From = "<" + from + ">"
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	if attachFile != "" {
		e.AttachFile(attachFile)
	}

	if attachDir != "" {

	}

	e.Send(sendserver, smtp.PlainAuth("", from, passwd, server))
}
