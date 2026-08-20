package src

import (

    "gopkg.in/gomail.v2"
)

// === === === Email send === === === #
func createMessage(email, message string) *gomail.Message {
    mail := gomail.NewMessage()

    mail.SetHeader("From", mail.FormatAddress(USERNAME, "portfolio"))
    mail.SetHeader("To", USERNAME)
    mail.SetHeader("Subject", email)

    mail.SetBody("text/html", createMailView(email, message))

    return mail
}

