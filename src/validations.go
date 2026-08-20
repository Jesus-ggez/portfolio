package src

import (
    "log"
    "net/mail"
)


func emailValidator(email string) string {
    if len(email) == 0 { return "email is required" }

    if _, err := mail.ParseAddress(email); err != nil {
        log.Print("error parsing email: ", err.Error())
        return "invalid client email format"
    }

    return ""
}

func messageValidator(msg string) string {
    if len(msg) == 0 { return "message is required" }

    if len(msg) > 3000 { return "message is too long" }

    return ""
}



