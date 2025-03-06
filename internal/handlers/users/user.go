package users

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/doktorChopper/ek-service/internal/handlers"
	"github.com/doktorChopper/ek-service/internal/models"
	"github.com/doktorChopper/ek-service/internal/store"
	"github.com/doktorChopper/ek-service/internal/store/users"
)

type UserHandler struct {
    store store.User
}

func AddRouters(mux *http.ServeMux, db *sql.DB) {

    store := users.New(db)

    userHandler := New(&store)

    mux.HandleFunc("/home", handlers.Home)
    mux.HandleFunc("/user/{id}", userHandler.Get)
    mux.HandleFunc("/user/create", userHandler.RenderCreateUserForm)
    mux.HandleFunc("/user/create/submit", userHandler.CreateUser)
}

func New(user store.User) UserHandler {
    return UserHandler{
        store: user,
    }
}

func (u *UserHandler) Get(w http.ResponseWriter, r * http.Request) {

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


func (u *UserHandler) CreateUser(w http.ResponseWriter, r * http.Request) {
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

func (u *UserHandler) RenderCreateUserForm(w http.ResponseWriter, r *http.Request) {

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


