package server

import (
	"log"
	"net/http"

	"github.com/doktorChopper/ek-service/internal/config"
	"github.com/doktorChopper/ek-service/internal/database"
	handlerUser "github.com/doktorChopper/ek-service/internal/handlers/users"
)

type Server struct {
    cfg     *config.Config
}

func NewServer(cfg *config.Config) *Server {

    return &Server{
        cfg: cfg,
    }
}

func (s *Server) RunServer() {

    mux := http.NewServeMux()

    db, err := database.ConnectToPostgre(s.cfg)
    if err != nil {
        log.Println("could not connect to sql, err:", err)
        return
    }

    handlerUser.AddRouters(mux, db)


    srv := http.Server{
        Addr: s.cfg.Server.Addr + s.cfg.Server.Port,
        Handler: mux,
    }

    fs := http.FileServer(http.Dir("templates/static/"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("launching server...")
    err = srv.ListenAndServe()
    if err != nil {
        log.Println(err.Error())
    }

}
