package mails

import (
	"github.com/wmd/utils"
	"net/smtp"
)

var (
	sender = "mail"
)

// SendMail func
func SendMail(user, mail, rand string) {
	from := sender
	pass := "password"
	to := mail

	msg :=
		"\nFrom: WMD Postman\n" +
			"To: " + mail + "\n" +
			"Subject: Verifying your Account\n" +
			"Hey " + user + ", you registration needs just one more step.\n\n" +
			"Verification Key : " + rand +
			"\n\nFor activating your Account please click on the following link and paste your Verification Key : \n\n" +
			"http://localhost:8080/verif"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		utils.LoggingErrorFile(err.Error())
	}
}
