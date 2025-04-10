package controller

import (
	"encoding/json"
	// "fmt"
	"html/template"
	// "log"
	"net/http"
	"strconv"

	"github.com/doktorChopper/ek-service/internal/models"
	"github.com/doktorChopper/ek-service/internal/store"
)

var flag bool

type UserController struct {
    store store.User
}

// type UserCredHandler struct {
//     store store.UserCredStorer 
// }
//
// func NewUserCredHandler(cred store.UserCredStorer) UserCredHandler {
//     return UserCredHandler{
//         store: cred,
//     }
// }

func NewUserController(user store.User) UserController {
    return UserController{
        store: user,
    }
}

func Home(w http.ResponseWriter, r* http.Request) {
    flag = false
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

func (u *UserController) Get(w http.ResponseWriter, r * http.Request) {

    if r.Method == http.MethodGet {
        id := r.PathValue("id")

        i, err := strconv.Atoi(id)

        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        resp, err := u.store.Get(i)
        
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        body, _ := json.Marshal(resp)

        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _,_ = w.Write(body)

    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }

}


func (u *UserController) CreateUser(w http.ResponseWriter, r * http.Request) {
    user := models.User{
        Name:       r.PostFormValue("name"),
        Surname:    r.PostFormValue("surname"),
        Email:      r.PostFormValue("email"),
        }

    _, err := u.store.Create(user)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/user/create", http.StatusMovedPermanently)
}

func (u *UserController) RenderCreateUserForm(w http.ResponseWriter, r *http.Request) {

    files := []string{
        "templates/html/user.form.tmpl",
        "templates/html/base.layout.tmpl",
    }

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, nil)
}

// func (u *UserCredHandler) Login(w http.ResponseWriter, r *http.Request) {
//
//     uc := models.UserCred{
//         Login: r.FormValue("login"),
//         Password: r.FormValue("password"),
//     }
//
//     res, _ := u.store.Login(uc.Login)
//
//     fmt.Println(uc.Password)
//     fmt.Println(res.HashedPassword)
//
//     if !comparePasswords(uc.Password, res.HashedPassword) {
//         w.Write([]byte("ERROR"))
//         return
//     }
//
//     flag = true
//     http.Redirect(w, r, "/home", http.StatusFound)
// }

func comparePasswords(formPass string, DBPass string) bool {

    for i := 0; i < max(len(formPass), len(DBPass)); i += 1 {
        if formPass[i] != DBPass[i] {
            return false
        }
    }

    return true
}

func Authorized(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r * http.Request) {
        if !flag {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next(w, r)
    }
}

// func (u *UserCredHandler) LoginForm(w http.ResponseWriter, r * http.Request) {
//
//     files := []string{
//         "templates/html/login.page.tmpl",
//         "templates/html/base.layout.tmpl",
//     }
//
//     t, err := template.ParseFiles(files...)
//     if err != nil {
//         log.Println(err)
//         return
//     }
//
//     err = t.Execute(w, nil)
//     if err != nil {
//         log.Println(err)
//         return
//     }
//
// }
//
//
