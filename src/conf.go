package src

import (
    "embed"
    "html/template"
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/joho/godotenv"
    "gopkg.in/gomail.v2"
)



var (
    EMBED_FS embed.FS

    USERNAME string
    PASSWORD string
    ADDRESS string
    PORT string

    GMAIL_SMTP_SSL string
    GMAIL_SSL_PORT int

    indexTempl *template.Template
    emailTempl *template.Template

    SMTP_SERVE *gomail.Dialer
    PROD bool
)

func InitApp(appFS embed.FS) {
    if err := godotenv.Load(); err != nil {
        log.Print("\nWarning: Not load dotenv file\n")
    }

    EMBED_FS = appFS

    USERNAME            = os.Getenv("USERNAME")
    PASSWORD            = os.Getenv("PASSWORD")
    ADDRESS             = os.Getenv("ADDRESS")
    PORT                = getPort()

    GMAIL_SMTP_SSL      = os.Getenv("GMAIL_SMTP_SSL")
    GMAIL_SSL_PORT, _   = strconv.Atoi(os.Getenv("GMAIL_SSL_PORT"))

    indexTempl          = loadTemplate("root")
    emailTempl          = loadTemplate("email")

    SMTP_SERVE = createEmailConf()
    PROD = os.Getenv("PROD") != "0"
}


// === === === Types === === === #
type ListStr []string

type Stylesheet struct {
    CSSFiles ListStr
}


func getPort() string {
    port := "3000"

    if envPort := os.Getenv("PORT"); envPort != "" {
        port = envPort
    }

    return port
}

func Serve(frontend, backend *http.ServeMux) {
    app := http.NewServeMux()

    app.Handle("/api/", http.StripPrefix("/api", backend))
    app.Handle("/", frontend)

    log.Printf("init serve at port %s", PORT)
    log.Fatal(http.ListenAndServe(":" + PORT, app))
}
