package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(context.Context, dto.LoginRequest) (string, error)
}

type service struct {
	userRepository user.Repository
}

func (s *service) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	usr, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJWT(usr.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key"))
}

func NewService(ur user.Repository) Service {
	return &service{userRepository: ur}
}
