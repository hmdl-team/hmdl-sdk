package sdk

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
)

type EmailConfig struct {
	SMTPServer string
	Port       string
	UserName   string
	Password   string
}

func (config *EmailConfig) EmailSent(form string, to []string, name string, body string) error {

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
	bodyMessage := config.writeHTMLEmail(form, Receiver, Subject, message)

	return config.sendMail(form, Receiver, Subject, bodyMessage)

}

func (config *EmailConfig) sendMail(from string, Dest []string, Subject, bodyMessage string) error {
	msg := "From: " + from + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(
		config.SMTPServer+":"+config.Port,
		smtp.PlainAuth("", config.UserName, config.Password, config.SMTPServer),
		config.UserName,
		Dest,
		[]byte(msg))

	if err != nil {
		fmt.Printf("smtp error: %s", err)
		return err
	}

	return nil
}

func (config *EmailConfig) writeEmail(from string, dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = from

	receipient := ""

	for _, user := range dest {
		receipient = receipient + ","+ user
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
	_, _ = finalMessage.Write([]byte(bodyMessage))
	_ = finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (config *EmailConfig) writeHTMLEmail(from string, dest []string, subject, bodyMessage string) string {
	return config.writeEmail(from, dest, "text/html", subject, bodyMessage)
}

func (config *EmailConfig) writePlainEmail(from string, dest []string, subject, bodyMessage string) string {
	return config.writeEmail(from, dest, "text/plain", subject, bodyMessage)
}
