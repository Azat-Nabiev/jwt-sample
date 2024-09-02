package repositories

import (
	"auth-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepositoryImpl struct {
	db     *pgxpool.Pool
	logger *zap.SugaredLogger
}

func NewUserRepositoryImpl(db *pgxpool.Pool, logger *zap.SugaredLogger) UserRepository {
	return &UserRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (u *UserRepositoryImpl) SaveUser(user *models.User) (*models.User, error) {

	return nil, nil
}
