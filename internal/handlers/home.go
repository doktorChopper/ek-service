package handlers

import (
	"html/template"
	"net/http"
)


func Home(w http.ResponseWriter, r* http.Request) {
    renderHomeTemplate(w)
}

func renderHomeTemplate(w http.ResponseWriter) {

    files := []string{
        "templates/html/home.page.tmpl",
        "templates/html/base.layout.tmpl",
    }

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    t.Execute(w, nil)
}

