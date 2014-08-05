package main

import (
    "html/template"
    "net/http"
)

type Dat struct {
    Title string
    Body []string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    funcMap := template.FuncMap{
        "safehtml": func(text string) template.HTML {
                        return template.HTML(text)
                    },
    }

    templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/layout.html", "templates/view.html"))

    dat := Dat{
        Title: "ほげやん",
        Body: []string{"1: ほげではない。", "2: ほげではない。"},
    }

    err := templates.ExecuteTemplate(w, "layout", dat)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8888", nil)
}
