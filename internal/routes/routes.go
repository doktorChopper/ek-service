package routes

import (
	"database/sql"
	"net/http"

	"github.com/doktorChopper/ek-service/internal/controller"
	"github.com/doktorChopper/ek-service/internal/middleware"
	"github.com/doktorChopper/ek-service/internal/store"
)

func AddRouters(mux *http.ServeMux, db *sql.DB) {

    st := store.NewUser(db)
    userHandler := controller.NewUserController(st)

    // credStore := store.NewUserCredStorer(db)
    // credHandler := store.NewUserCredHandler(*credStore)

    // mux.HandleFunc("/home", controller.Authorized(controller.Home))

    mux.HandleFunc("/register", userHandler.Register)

    mux.HandleFunc("/login", middleware.RenderLoginForm)
    mux.HandleFunc("/login/submit", middleware.LoginMiddleware(db, userHandler.RenderCreateUserForm))
    mux.HandleFunc("/user/{id}", userHandler.Get)
    mux.HandleFunc("/user/create", middleware.LoginMiddleware(db, userHandler.RenderCreateUserForm))
    mux.HandleFunc("/user/create/submit", userHandler.CreateUser)
    // mux.HandleFunc("/login/submit", credHandler.Login)
    // mux.HandleFunc("/login", credHandler.LoginForm)


    store := store.NewFilm(db)

    filmHandler := controller.NewFilmController(store)

    mux.HandleFunc("/user/{id}/films", filmHandler.GetFilms)
    mux.HandleFunc("/user/{id}/film/add", filmHandler.RenderAddFilmForm)
    mux.HandleFunc("/user/{id}/film/add/submit", filmHandler.AddFilm)
}
