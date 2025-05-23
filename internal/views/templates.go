package views

import (
	"html/template"
	"log"
	"net/http"

	"github.com/doktorChopper/ek-service/internal/models"
	// "github.com/doktorChopper/ek-service/internal/models"
)

const (
    BASE            string = "templates/html/base.layout.tmpl"
    HOME            string = "templates/html/home.page.tmpl"
    USER_FORM       string = "templates/html/user.form.tmpl"
    REGISTER_FORM   string = "templates/html/register.form.tmpl"
    LOGIN_FORM      string = "templates/html/login.page.tmpl"
)

func RenderHomeTemplate(w http.ResponseWriter, r *http.Request, u *models.User) {

    files := []string{HOME, BASE}

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, u)
    if err != nil {
        log.Println(err)
    }
}

func RenderLoginForm(w http.ResponseWriter, r *http.Request) {

    files := []string{LOGIN_FORM, BASE}

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, nil)
}

func RenderCreateUserForm(w http.ResponseWriter, r *http.Request) {

    files := []string{USER_FORM, BASE}

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, nil)
}

func RenderRegisterForm(w http.ResponseWriter, r *http.Request) {

    files := []string{REGISTER_FORM, BASE}

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, nil)
}
