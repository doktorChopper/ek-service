package store

import (
	"database/sql"

	"github.com/doktorChopper/ek-service/internal/models"
)

type SessionStorer struct {
    db *sql.DB
}


func (s *SessionStorer) Get(sessionID int) (models.Session, error) {

    var session models.Session

    stmt := `SELECT * FROM sessions WHERE id = $1`
    row := s.db.QueryRow(stmt, sessionID)

    row.Scan(&session.SessionID, &session.UserID, &session.ExpiresAt)

    return session, nil
}

func (s *SessionStorer) Create(session *models.Session) error {

    stmt := `INSERT INTO sessions (id, user_id, expires_at)
    VALUES ($1, $2, $3)`

    _, err := s.db.Exec(stmt, &session.SessionID, &session.UserID, &session.ExpiresAt)
    if err != nil {
        return err
    }

    return nil
}
