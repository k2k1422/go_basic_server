package SMTP

import (
	"fmt"
	"net/smtp"
	"server/Logging"
)

func Send(receiver []string, subject string, body string) bool {
	/*
		SMTP server
	*/

	// Formatting  the body mail meta and body

	header := make(map[string]string)
	header["From"] = Credential["Sender"]
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; text/css; charset=\"utf-8\""

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + body

	if err := smtp.SendMail(
		Credential["AuthServer"],
		smtp.PlainAuth(
			"",
			Credential["Sender"],
			Credential["AppKey"],
			Credential["Host"],
		),
		Credential["Sender"],
		receiver,
		[]byte(message)); err != nil {
		// Failed to send the mail
		Logging.ERROR.Println("SMTP error- ", err, receiver)
		return false
	} else {
		// Successfully sent the mail
		Logging.INFO.Println("SMTP success- ", receiver)
		return true
	}
}
