package src

import "net/http"

func LoadFrontend() *http.ServeMux {
    router:= http.NewServeMux()

    router.Handle("/static/", http.FileServerFS(EMBED_FS))

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := Stylesheet { CSSFiles: getCSS("root") }

        if err := indexTempl.ExecuteTemplate(w, "index.html", data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    return router
}
