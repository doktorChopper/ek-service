package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doktorChopper/ek-service/internal/config"
)


func ConnectToPostgre(c *config.Config) (*sql.DB, error) {

    dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s",
        c.Database.User,
        c.Database.Password,
        c.Database.Name,
        c.Database.Port)

    db, err := sql.Open(c.Database.Driver, dsn)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    if err = db.Ping(); err != nil {
        log.Fatalln(err)
        return nil, err
    }

    return db, nil
}

