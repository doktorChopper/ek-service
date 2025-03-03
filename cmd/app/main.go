package main

import (
	"fmt"

	"github.com/doktorChopper/ek-service/internal/config"
	"github.com/doktorChopper/ek-service/internal/server"
)


func main() {

    cfg := config.New()

    fmt.Println("start app...")
    srv := server.NewServer(cfg)
    srv.RunServer()
}
