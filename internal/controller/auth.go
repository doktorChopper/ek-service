package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/doktorChopper/ek-service/internal/models"
	"github.com/doktorChopper/ek-service/internal/store"
	"github.com/doktorChopper/ek-service/internal/views"
	"github.com/google/uuid"
)

type AuthController struct {
    name        string
    session     *store.SessionStore
    authStore   *store.AuthStore
}

func NewAuthController(s *store.SessionStore, a *store.AuthStore) *AuthController {
    return &AuthController {
        name:       "AuthController",
        session:    s,
        authStore:  a,
    }
}

func (a *AuthController) LoggerName() string {
    return a.name
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        views.RenderLoginForm(w, r)
    } else if r.Method == http.MethodPost {

        email := r.FormValue("email")
        password := r.FormValue("password")

        user, err := a.authStore.Login(email, password)
        if err != nil {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        sessionID := uuid.NewString()
        expiresAt := time.Now().Add(600 * time.Second)

        session := models.Session {
            ID:         sessionID,
            UserID:     user.ID,
            ExpiresAt:  expiresAt,
        }

        if err := a.session.Create(&session); err != nil {
            http.Error(w, "Failed to create session", http.StatusInternalServerError)
            return
        }

        cookie := http.Cookie {
            Name:       "session_token",
            Value:      sessionID,
            Expires:    expiresAt,
        }
        
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/id/" + strconv.Itoa(user.ID), http.StatusSeeOther)
    }
}

func (a *AuthController) Register(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodGet {
        views.RenderRegisterForm(w, r)
    } else if r.Method == http.MethodPost {

        user := models.User {
            Name:           r.FormValue("name"),
            Surname:        r.FormValue("surname"),
            Email:          r.FormValue("email"),
            Password:       r.FormValue("password"),
        }

        err := a.authStore.Register(&user)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/register", http.StatusSeeOther)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func (a *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
    
    if r.Method == http.MethodGet {

        cookie, err := r.Cookie("session_token")
        if err != nil {
            return
        }

        err = a.session.Delete(cookie.Value)
        if err != nil {
            return
        }

    } else {
        w.WriteHeader(http.StatusBadRequest)
    }
}



