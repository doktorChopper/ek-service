package store

import (
	"database/sql"

	"github.com/doktorChopper/ek-service/internal/models"
)

type SessionStore struct {
    db *sql.DB
}

func NewSessionStore(db *sql.DB) *SessionStore {
    return &SessionStore {
        db: db,
    }
}


func (s *SessionStore) Get(sessionID int) (models.Session, error) {

    var session models.Session

    stmt := `SELECT * FROM sessions WHERE id = $1`
    row := s.db.QueryRow(stmt, sessionID)

    row.Scan(&session.ID, &session.UserID, &session.ExpiresAt)

    return session, nil
}

func (s *SessionStore) Delete(sessionID int) error {

    stmt := `DELETE FROM sessions WHERE id = $1`

    _, err := s.db.Exec(stmt, sessionID)
    if err != nil {
        return err
    }

    return nil
}

func (s *SessionStore) Create(session *models.Session) error {

    stmt := `INSERT INTO sessions (id, user_id, expires_at)
    VALUES ($1, $2, $3)`

    _, err := s.db.Exec(stmt, &session.ID, &session.UserID, &session.ExpiresAt)
    if err != nil {
        return err
    }

    return nil
}
