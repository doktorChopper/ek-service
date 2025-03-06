package store

import "github.com/doktorChopper/ek-service/internal/models"

type User interface {
    Get(int) ([]models.User, error)
    Create(models.User) (models.User, error)
}
