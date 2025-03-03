package server

import (
	"log"
	"net/http"

	"github.com/doktorChopper/ek-service/internal/config"
	"github.com/doktorChopper/ek-service/internal/handlers"
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

    mux.HandleFunc("/home", handlers.Home)
    
    srv := http.Server{
        Addr: s.cfg.Server.Addr + s.cfg.Server.Port,
        Handler: mux,
    }

    log.Println("launching server...")
    err := srv.ListenAndServe()
    if err != nil {
        log.Println(err.Error())
    }

}
