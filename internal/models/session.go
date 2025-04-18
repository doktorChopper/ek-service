package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
    SessionID   string
    UserID      int
    ExpiresAt   time.Time
}

func (s *Session) SID() string {
    return s.SessionID
}

func (s *Session) EAT() time.Time {
    return s.ExpiresAt
}

func NewSession(userID int, expiresAfter time.Duration) *Session {

    sessionID := uuid.New().String()
    expiresAt := time.Now().Add(expiresAfter)

    session := &Session {
        SessionID:  sessionID,
        UserID:     userID,
        ExpiresAt:  expiresAt,
    }

    return session
}

func (s *Session) isExpired() bool {
    return s.ExpiresAt.Before(time.Now())
}
