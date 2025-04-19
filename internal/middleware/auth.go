package middleware

import (
	"net/http"
	"time"

	"github.com/doktorChopper/ek-service/internal/models"
	"github.com/doktorChopper/ek-service/internal/store"
)


func AuthMiddleware(s *store.SessionStore, next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        password := r.FormValue("password")
        email := r.FormValue("login")

        if password == "" || email == "" {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
        }

        var checkPassword string
        var id int

        // stmt := `SELECT id, hashed_password FROM users WHERE email = $1` 

        // row := s.QueryRow(stmt, email)
        // row.Scan(&id, &checkPassword)

        flag := true

        if len(checkPassword) != len(password) {
            flag = false
        }

        for i := range checkPassword {
            if checkPassword[i] != password[i] {
               flag = false
            } 
        }

        if !flag {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        } else {
            session := models.NewSession(id, 120 * time.Second)

            err := s.Create(session)
            if err != nil {
                return
            }

            cookie := http.Cookie {
                Name:       "session_token",
                Value:      session.ID,
                Expires:    session.ExpiresAt,
            }

            http.SetCookie(w, &cookie)
        }

        next.ServeHTTP(w, r)
    })
}

