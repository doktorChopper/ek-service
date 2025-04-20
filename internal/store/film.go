package store

import (
	"database/sql"

	"github.com/doktorChopper/ek-service/internal/models"
)

type FilmStorer struct {
    db *sql.DB
} 

func NewFilm(db *sql.DB) FilmStorer {
    return FilmStorer{
        db: db,
    }
}

func (f FilmStorer) GetFilmByUser(user_id int) ([]models.Film, error) {
    var (
        rows *sql.Rows
        err error
    )

    stmt := `SELECT id, name, genre, review, rate FROM films WHERE user_id = $1`

    rows, err = f.db.Query(stmt, user_id)

    if err != nil {
        return nil, err
    }

    var films []models.Film

    for rows.Next() {
        var f models.Film
        _ = rows.Scan(
            &f.ID,
            &f.Name,
            &f.Genre,
            &f.Review,
            &f.Rate)

        films = append(films, f)
    }

    return films, nil
}

func (f FilmStorer) AddFilmToUser(film models.Film) (models.Film, error) {

    stmt := `INSERT INTO films (name, genre, review, rate, user_id)
    VALUES($1, $2, $3, $4, $5)`

    res, err := f.db.Exec(stmt,
        film.Name,
        film.Genre,
        film.Review,
        film.Rate,
        film.UserID)

    if err != nil {
        return models.Film{}, nil
    }

    id, _ := res.LastInsertId()

    film.ID = id

    return film, nil
}

/*
func (f FilmStorer) DeleteFilm() {

}
*/
