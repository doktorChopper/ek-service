package middleware

import (
	"log"
	"net/http"

	"github.com/doktorChopper/ek-service/internal/store"
)


func AuthMiddleware(s *store.SessionStore, next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

        cookie, err := r.Cookie("session_token")
        if err != nil {
            // http.Error(w, "No cookie", http.StatusInternalServerError)
            log.Println(err.Error())

            http.Redirect(w, r, "/login", http.StatusUnauthorized)
            return
        }

        session, err := s.Get(cookie.Value)
        if err != nil {
            // http.Error(w, "No session", http.StatusInternalServerError)
            log.Println(err.Error())

            http.Redirect(w, r, "/login", http.StatusUnauthorized)
            return
        }

        if session.IsExpired() {
            err = s.Delete(cookie.Value)
            if err != nil {
                w.WriteHeader(http.StatusUnauthorized)
                return
            }

            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }


        next.ServeHTTP(w, r)
    })
}

