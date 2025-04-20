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

func (s *SessionStore) Get(sessionID string) (*models.Session, error) {

    var session models.Session

    stmt := `SELECT id, user_id, expires_at FROM sessions WHERE id = $1`
    row := s.db.QueryRow(stmt, sessionID)

    err := row.Scan(&session.ID, &session.UserID, &session.ExpiresAt)
    if err != nil {
        return nil, err
    }

    return &session, nil
}

func (s *SessionStore) Delete(sessionID string) error {

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

func (s *SessionStore) GC() error {

    stmt := `DELETE FROM sessions WHERE expires_at < Now()`

    _, err := s.db.Exec(stmt)
    if err != nil {
        return err
    }

    return nil
}
