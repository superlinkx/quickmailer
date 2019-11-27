package main

import (
	"flag"
	"log"
	"net/smtp"
	"net/textproto"
	"path/filepath"

	"github.com/jordan-wright/email"
)

type emailParams struct {
	server     string
	sendserver string
	port       string
	passwd     string
	from       string
	fromname   string
	to         string
	subject    string
	body       string
	attachFile string
	attachDir  string
}

func main() {
	settings := getArgs()

	fromline := "<" + settings.from + ">"

	if settings.fromname != "" {
		fromline = settings.fromname + " " + fromline
	}

	e := &email.Email{
		To:      []string{settings.to},
		From:    fromline,
		Subject: settings.subject,
		Text:    []byte(settings.body),
		Headers: textproto.MIMEHeader{},
	}

	if settings.attachFile != "" {
		_, err := e.AttachFile(settings.attachFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	if settings.attachDir != "" {
		files, _ := filepath.Glob(settings.attachDir + "/*")
		for _, elem := range files {
			_, err := e.AttachFile(elem)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err := e.Send(settings.sendserver, smtp.PlainAuth("", settings.from, settings.passwd, settings.server))

	if err != nil {
		log.Fatal(err)
	}
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
	flag.StringVar(&params.fromname, "fromName", "", "name to send email as")

	flag.Parse()

	if (params.server == "") || (params.port == "") || (params.passwd == "") || (params.from == "") || (params.to == "") {
		log.Fatal("Required flags: -server, -passwd, -port, -from, -to")
	}

	params.sendserver = params.server + ":" + params.port

	return params
}
