package routes

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/doktorChopper/ek-service/internal/controller"
	"github.com/doktorChopper/ek-service/internal/middleware"
	"github.com/doktorChopper/ek-service/internal/store"
	"github.com/doktorChopper/ek-service/internal/views"
)



func AddRouters(mux *http.ServeMux, db *sql.DB) {

    st := store.NewUserStore(db)
    userController := controller.NewUserController(st)

    sessionStore := store.NewSessionStore(db)
    authStore := store.NewAuthStore(st)

    authController := controller.NewAuthController(sessionStore, authStore)

    // credStore := store.NewUserCredStorer(db)
    // credHandler := store.NewUserCredHandler(*credStore)

    // mux.HandleFunc("/home", controller.Authorized(controller.Home))

    mux.HandleFunc("/register", middleware.LoggerMiddleware(authController, authController.Register))
    mux.HandleFunc("/login", middleware.LoggerMiddleware(authController, authController.Login))
    mux.HandleFunc("/logout", middleware.LoggerMiddleware(authController, authController.Logout))

    mux.HandleFunc("/id/{id}", middleware.AuthMiddleware(sessionStore, userController.Home))

    // mux.HandleFunc("/user/{id}", userController.Get)
    mux.HandleFunc("/user/{id}", middleware.AuthMiddleware(sessionStore, userController.Get))
    mux.HandleFunc("/user/create", middleware.AuthMiddleware(sessionStore, views.RenderCreateUserForm))
    mux.HandleFunc("/user/create/submit", userController.CreateUser)

    store := store.NewFilm(db)

    filmController := controller.NewFilmController(&store)

    mux.HandleFunc("/user/{id}/films", filmController.GetFilms)
    // mux.HandleFunc("/user/{id}/film/add", filmController.RenderAddFilmForm)
    mux.HandleFunc("/user/{id}/film/add/submit", filmController.AddFilm)

    go func() {
        for {
            time.Sleep(1 * time.Minute)
            log.Println("start session GC")
            sessionStore.GC()
        }
    } ()
}
