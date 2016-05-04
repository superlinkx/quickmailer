package main

import (
	"flag"
	"net/smtp"
	"path/filepath"

	"github.com/jordan-wright/email"
)

type emailParams struct {
	server     string
	sendserver string
	port       string
	passwd     string
	from       string
	to         string
	subject    string
	body       string
	attachFile string
	attachDir  string
}

func main() {
	settings := getArgs()

	e := email.NewEmail()
	e.From = "<" + settings.from + ">"
	e.To = []string{settings.to}
	e.Subject = settings.subject
	e.Text = []byte(settings.body)

	if settings.attachFile != "" {
		e.AttachFile(settings.attachFile)
	}

	if settings.attachDir != "" {
		files, _ := filepath.Glob(settings.attachDir + "/*")
		for _, elem := range files {
			e.AttachFile(elem)
		}
	}

	e.Send(settings.sendserver, smtp.PlainAuth("", settings.from, settings.passwd, settings.server))
}

func getArgs() emailParams {
	var params emailParams

	flag.StringVar(&params.server, "server", "", "smtp server without port to use")
	flag.StringVar(&params.port, "port", "587", "smtp port to use")
	flag.StringVar(&params.passwd, "passwd", "", "password for smtp auth")
	flag.StringVar(&params.from, "from", "", "email to send from")
	flag.StringVar(&params.to, "to", "", "email to send to")
	flag.StringVar(&params.subject, "subject", "Your Files", "subject line")
	flag.StringVar(&params.body, "body", "See attached", "some body text")
	flag.StringVar(&params.attachFile, "attach", "", "single file to attach")
	flag.StringVar(&params.attachDir, "attachDir", "", "directory to attach files from")

	flag.Parse()

	if (params.server == "") || (params.port == "") || (params.passwd == "") || (params.from == "") || (params.to == "") {
		panic("Required flags: -server, -passwd, -port, -from, -to")
	}

	params.sendserver = params.server + ":" + params.port

	return params
}
