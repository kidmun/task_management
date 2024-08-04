package usecase

import (
	"context"
	"task_management/internal/core/models"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.StandardClaims
}
type userUsecase struct {
	userRepository models.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository models.UserRepository, timeout time.Duration) models.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}


func (uu *userUsecase) RegisterUser(c context.Context, user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c,uu.contextTimeout)
	defer cancel()
	return uu.userRepository.RegisterUser(ctx, user)
}

func (uu *userUsecase) LoginUser(c context.Context, userInput models.UserInput) (string, error) {
	ctx, cancel := context.WithTimeout(c,uu.contextTimeout)
	defer cancel()
	return uu.userRepository.LoginUser(ctx, userInput)
}

func (uu *userUsecase) RegisterAdmin(c context.Context, user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c,uu.contextTimeout)
	defer cancel()
	return uu.userRepository.RegisterAdmin(ctx, user)
}
