package src

import (
    "bytes"
    "html/template"
    "io/fs"
    "log"
    "strings"

    "gopkg.in/gomail.v2"
)


func loadTemplate(name string) *template.Template {
    dirPath := "templates/" + name + "/"
    dirPartsPath := dirPath + "parts/"
    dynamicCssDir := "templates/shared/*.html"

    return template.Must(template.ParseFS(EMBED_FS,
        dirPath + "*.html",
        dirPartsPath + "*.html",
        dynamicCssDir,
    ))
}

func createEmailConf() *gomail.Dialer {
    return gomail.NewDialer(
        GMAIL_SMTP_SSL,
        GMAIL_SSL_PORT,
        USERNAME,
        PASSWORD,
    )
}

func createMailView(email, content string) string {
    var body bytes.Buffer

    if err := emailTempl.ExecuteTemplate(&body, "index.html", struct {
        Email string
        Message string
        CSSFiles ListStr
    }{ Email: email, Message: content, CSSFiles: getCSS("email") }); err != nil {
        log.Print("Error sending messag: " + err.Error())
        return err.Error() + "\n\n" + email + "\n\n" + content
    }

    return body.String()
}


// === === === CSS module === === === #
var cssCache = make(map[string]ListStr)
func getCSS(dir string) (css ListStr) {
    if cached, ok := cssCache[dir]; ok { return cached }

    staticSub, err := fs.Sub(EMBED_FS, "static/css/" + dir)
    if err != nil { panic(err.Error()) }

    entries, err := fs.ReadDir(staticSub, ".")
    if err != nil { panic(err.Error()) }

    for _, entry := range entries {
        if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".css") {
            css = append(css, dir + "/" + entry.Name())
        }
    }
    cssCache[dir] = css
    return
}

