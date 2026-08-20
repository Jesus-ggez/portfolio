package src

import (
    "log"
    "net/http"
    "strings"
)


func LoadBackend() *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("/email", emailHandler)

    return router
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
        return
    }

    message := strings.TrimSpace(r.FormValue("message"))
    email := strings.TrimSpace(r.FormValue("email"))

    if err := messageValidator(message); err != "" {
        http.Error(w, err, http.StatusBadRequest)
        return
    }

    if err := emailValidator(email); err != "" {
        http.Error(w, err, http.StatusBadRequest)
        return
    }

    if !PROD {
        log.Printf(`Email sending correctly:
            Email: %s
            Message: %s`, email, message)
        w.WriteHeader(http.StatusNoContent)
        return
    }

    payload := createMessage(email, message)

    if err := SMTP_SERVE.DialAndSend(payload); err != nil {
        log.Print("Mail not sent: " + err.Error())
        http.Error(w, "Error sending email", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
