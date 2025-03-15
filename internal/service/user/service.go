package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	"github.com/pujidjayanto/goginboilerplate/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)
	Register(context.Context, dto.RegisterRequest) error
}

type service struct {
	userRepository user.Repository
}

func (s *service) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredential
	}

	token, err := generateJWT(user.Id)
	if err != nil {
		return nil, fmt.Errorf("error generate token, %v", err)
	}

	return &dto.LoginResponse{Token: token}, nil
}

func (s *service) Register(ctx context.Context, req dto.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generate password, %v", err)
	}

	newUser := &user.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepository.Create(ctx, newUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrUserAlreadyExisted
		}
		return fmt.Errorf("error create user, %v", err)
	}

	return nil
}

func generateJWT(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key")) // need to pass the env to here
}

func NewService(ur user.Repository) Service {
	return &service{userRepository: ur}
}
