package main

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

const (
	_INTERNET_ERROR_MSG = "</br>This email was attempted to be sent on %s but there was not internet.</br>"
)

var (
	emailAdditionalInfo string
	emailSent           bool
)

func sendEmailReportingDockerFailure() {

	appendLog("[Email]: Trying send notification")

	creds := config.Credentials
	email := config.Email

	m := gomail.NewMessage()
	m.SetHeader("From", creds.Email)
	m.SetHeader("To", email.To...)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", emailAdditionalInfo+email.Body)

	d := gomail.NewDialer(creds.SmtpHost, creds.SmtpPort, creds.Email, creds.Password)

	if err := d.DialAndSend(m); err != nil {
		strErr := err.Error()

		internetProblem := strings.Contains(strErr, "no such host")
		if internetProblem && emailAdditionalInfo == "" {
			emailAdditionalInfo = fmt.Sprintf(_INTERNET_ERROR_MSG, time.Now().Format(config.Log.Format))
		}
		appendLog("[Email]: Error: " + err.Error())

	} else {
		appendLog("[Email]: Success")
		emailSent = true
		emailAdditionalInfo = ""
	}
}
