package middleware

import (
	"log"
	"net/http"

	"github.com/doktorChopper/ek-service/internal/controller"
)


func LoggerMiddleware(c controller.Controller, next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Started %s: %s  %s\n", c.LoggerName(), r.Method, r.URL)
        next.ServeHTTP(w, r)
        log.Printf("Completed %s\n", c.LoggerName())
    })
}
