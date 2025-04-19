package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
    ID          string
    UserID      int64
    ExpiresAt   time.Time
}


func NewSession(userID int, expiresAfter time.Duration) *Session {

    sessionID := uuid.New().String()
    expiresAt := time.Now().Add(expiresAfter)

    session := &Session {
        ID:  sessionID,
        UserID:     int64(userID),
        ExpiresAt:  expiresAt,
    }

    return session
}

func (s *Session) isExpired() bool {
    return s.ExpiresAt.Before(time.Now())
}
