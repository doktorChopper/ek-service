package store

import (
	"database/sql"

	"github.com/doktorChopper/ek-service/internal/models"
)

type UserStore struct {
    db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
    return &UserStore{
        db: db,
    }
} 

func (u *UserStore) ComparePassword(password, email string) bool {

    var checkPassword string

    stmt := `SELECT hashed_password FROM users WHERE email = $1` 

    row := u.db.QueryRow(stmt, email)
    row.Scan(&checkPassword)

    if len(checkPassword) != len(password) {
        return false
    }

    for i := range checkPassword {
        if checkPassword[i] != password[i] {
            return false
        } 
    }

    return true
}

func (u *UserStore) FindByEmail(email string) (*models.User, error) {

    var (
        row     *sql.Row
        user    models.User
    )

    stmt := `SELECT id, name, surname, email, hashed_password FROM users WHERE email = $1`
    
    row = u.db.QueryRow(stmt, email)
    
    row.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password)

    return &user, nil
}

func (u *UserStore) Get(id int) ([]models.User, error) {
    
    var (
        rows    *sql.Rows
        err     error
    )

    stmt := `SELECT id, name, surname, email FROM users WHERE id = $1`

    rows, err = u.db.Query(stmt, id)

    if err != nil {
        return nil, err
    }

    var users []models.User

    for rows.Next() {
        var u models.User

        _ = rows.Scan(
            &u.ID,
            &u.Name,
            &u.Surname,
            &u.Email)

        users = append(users, u)
    }
    
    return users, nil

}

func (u *UserStore) Create(user *models.User) (*models.User, error) {

    stmt := `INSERT INTO users (name, surname, email, hashed_password)
    VALUES ($1, $2, $3, $4)`

    res, err := u.db.Exec(stmt, user.Name, user.Surname, user.Email, user.Password)

    if err != nil {
        return &models.User{}, err
    }

    id, _ := res.LastInsertId()
    user.ID = int64(id)

    return user, nil
}

