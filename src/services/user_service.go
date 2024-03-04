package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"rest-api-golang-jwt/src/config"
	"rest-api-golang-jwt/src/entity"
	"rest-api-golang-jwt/src/helpers"
	"rest-api-golang-jwt/src/repositories"
	"time"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	db             *sql.DB
}

func NewUserService(userRepository *repositories.UserRepository, db *sql.DB) *UserService {
	return &UserService{UserRepository: userRepository, db: db}
}

func (s UserService) Login(ctx context.Context, user entity.UserLoginRequest) (entity.LoginResponse, error) {
	userData, err := s.UserRepository.GetByEmail(ctx, s.db, user.Username)
	if err != nil {
		return entity.LoginResponse{}, errors.New("user not found")
	}
	err = helpers.ComparePassword(userData.Password, user.Password)
	if err != nil {
		return entity.LoginResponse{}, errors.New("password not match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userData.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.JwtExp))),
	})

	tokenString, err := token.SignedString([]byte(config.JwtSecret))

	if err != nil {
		return entity.LoginResponse{}, errors.New("error generate token")
	}
	return entity.LoginResponse{
		Token: tokenString,
		User: entity.UserResponse{
			ID:      userData.ID,
			Name:    userData.Name,
			Address: userData.Address,
		},
	}, nil
}

func (s UserService) Get(ctx context.Context) ([]entity.Users, error) {
	users, err := s.UserRepository.Get(ctx, s.db)
	if err != nil {
		return []entity.Users{}, nil
	}
	return users, nil
}

func (s UserService) Insert(ctx context.Context, user entity.UserRequest) (int, error) {
	// set data id to user from google uuid
	user.ID = uuid.New().String()
	passwrod, err := helpers.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = passwrod

	if err != nil {
		return 0, err
	}
	result, err := s.UserRepository.Insert(ctx, s.db, user)
	if err != nil {
		return 0, err
	}
	return result, nil
}
