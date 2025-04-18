package middleware

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"

	"github.com/doktorChopper/ek-service/internal/models"
)

func RenderLoginForm(w http.ResponseWriter, r *http.Request) {

    files := []string{
        "templates/html/login.page.tmpl",
        "templates/html/base.layout.tmpl",
    }

    t, err := template.ParseFiles(files...)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, nil)
}

func LoginMiddleware(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        // token := r.Header.Get("Authorization")
        password := r.FormValue("password")
        email := r.FormValue("login")

        if password == "" || email == "" {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
        }

        var checkPassword string
        var id int

        stmt := `SELECT id, hashed_password FROM users WHERE email = $1` 

        row := db.QueryRow(stmt, email)
        row.Scan(&id, &checkPassword)

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
            s := models.NewSession(id, 120 * time.Second)
            stmt := `INSERT INTO sessions (id, user_id, expires_at)
            VALUES ($1, $2, $3)`

            _, err := db.Exec(stmt, &s.SessionID, &s.UserID, &s.ExpiresAt)
            if err != nil {
                return
            }

            cookie := http.Cookie {
                Name:       "session_token",
                Value:      s.SID(),
                Expires:    s.EAT(),
            }

            http.SetCookie(w, &cookie)
        }

        // if token != "secret" {
        //     http.Redirect(w, r, "/login", http.StatusSeeOther)
        //     return
        // }

        next.ServeHTTP(w, r)
    })
}

