package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/doktorChopper/ek-service/internal/models"
	"github.com/doktorChopper/ek-service/internal/store"
	"github.com/doktorChopper/ek-service/internal/views"
)

type UserController struct {
    name    string
    store   *store.UserStore
}

func NewUserController(user *store.UserStore) UserController {
    return UserController{
        name:   "UserController",
        store:  user,
    }
}

func (u *UserController) LoggerName() string {
    return u.name
}

func (u *UserController) Home(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodGet {

        userIDint, err := strconv.Atoi(r.URL.Path[len("/id/"):])
        if err != nil {
            return
        }

        user, err := u.store.Get(userIDint)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        films, err := u.store.GetFilmsByUserID(user.ID)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        user.Films = films

        views.RenderHomeTemplate(w, r, user)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
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


func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

    user := models.User{
        Name:       r.PostFormValue("name"),
        Surname:    r.PostFormValue("surname"),
        Email:      r.PostFormValue("email"),
    }

    _, err := u.store.Create(&user)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/user/create", http.StatusMovedPermanently)
}

