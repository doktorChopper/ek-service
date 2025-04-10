package store

// import (
// 	"database/sql"
// 	"log"
//
// 	"github.com/doktorChopper/ek-service/internal/models"
// )
//
// type UserCredStorer struct {
//     db  *sql.DB
// }
//
// func NewUserCredStorer(db *sql.DB) *UserCredStorer {
//     return &UserCredStorer{
//         db: db,
//     }
// }
//
// func (u *UserCredStorer) Login(login string) (models.User, error) {
//
//     var (
//         row *sql.Row
//         err error
//     )
//
//     stmt := `SELECT hashed_password, login FROM users WHERE login = $1`
//
//     row = u.db.QueryRow(stmt, login)
//
//     var user models.User
//
//     err = row.Scan(&user.HashedPassword, &user.Login)
//     log.Println(err)
//
//     return user, nil
// }
//
//
//
