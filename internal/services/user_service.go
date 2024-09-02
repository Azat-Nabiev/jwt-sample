package services

import (
	"auth-service/internal/handlers/dto/authenticate"
	"auth-service/internal/handlers/dto/reg"
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepository
	logger         *zap.SugaredLogger
}

func NewUserService(userRepository repositories.UserRepository, logger *zap.SugaredLogger) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (h *UserService) Register(regRequest *reg.UserRequest) (*reg.UserResponse, error) {
	var user, err = mapToModel(regRequest)

	if err != nil {
		h.logger.Errorw("Error during mapping the user:", err.Error())
		return nil, err
	}
	var savedUser *models.User

	savedUser, err = h.userRepository.SaveUser(&user)
	if err != nil {
		h.logger.Errorw("Error during saving the user", err.Error())
		return nil, err
	}

	return mapToResponse(savedUser), nil
}

func (h *UserService) Authenticate(authRequest *authenticate.UserAuthRequest) (*authenticate.UserAuthResponse, error) {

	return nil, nil
}

func mapToModel(regRequest *reg.UserRequest) (models.User, error) {
	var hashedPassword, err = hashPassword(regRequest.Password)

	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Name:     regRequest.Name,
		Login:    regRequest.Login,
		Password: hashedPassword,
	}, nil
}

func mapToResponse(user *models.User) *reg.UserResponse {
	return &reg.UserResponse{
		Name:  user.Name,
		Login: user.Login,
	}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
