package repositories

import "auth-service/internal/models"

type UserRepository interface {
	SaveUser(user *models.User) (*models.User, error)
}
