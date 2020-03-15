package lib

import (
	"bytes"
	"fmt"
	"github.com/getsentry/raven-go"
	"os"

	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

/**
	Modified from https://gist.github.com/jpillora/cb46d183eca0710d909a
	Thank you very much.
**/

const (
	/**
		Gmail SMTP Server
	**/
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = "587"
)

type Sender struct {
	Form     string
	User     string
	Password string
}

func EmailSent(form string, to []string, name string, body string) {

	sender := NewSender(form, os.Getenv("EMAIL_SERVER_NAME") , os.Getenv("EMAIL_SERVER_PASS")  )

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := to

	Subject := name
	message := `
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
		<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>
		` + body + `
	</body>
	</html>
	`
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)

	sender.SendMail(Receiver, Subject, bodyMessage)

}
func NewSender(Form, Username, Password string) Sender {

	return Sender{Form, Username, Password}
}

func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	fmt.Printf("====Sent mail====")
	msg := "From: " + sender.Form + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":"+SMTPPort,
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Printf("smtp error: %s", err)

		return
	}

	fmt.Println("Mail sent successfully!")
}

func (sender Sender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {
	return sender.WriteEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *Sender) WritePlainEmail(dest []string, subject, bodyMessage string) string {
	return sender.WriteEmail(dest, "text/plain", subject, bodyMessage)
}
