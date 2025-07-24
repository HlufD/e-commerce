package application

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HlufD/users-ms/internals/domain"
	"github.com/HlufD/users-ms/internals/ports"
)

type AuthService struct {
	userRepository ports.UserRepositoryPort
	hashing        ports.HashingPort
	token          ports.TokenPort
}

func NewAuthService(userRepository ports.UserRepositoryPort,
	hashing ports.HashingPort,
	token ports.TokenPort) *AuthService {
	return &AuthService{
		userRepository,
		hashing,
		token,
	}
}

func (au *AuthService) Login(credentials domain.Credentials) (*domain.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var err error

	user, err := au.userRepository.FindByUsername(ctx, credentials.Username)

	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	isValid := au.hashing.Compare(user.Password, credentials.Password)

	if !isValid {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := au.token.Generate(user.Id)

	if err != nil {
		return nil, fmt.Errorf("token generation failed: %w", err)
	}

	return &domain.Token{Token: token}, nil
}

func (au *AuthService) Register(registration domain.Registration) (*domain.User, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	userNameExits, err := au.userRepository.CheckIfUserExists(ctx, "username", registration.Username)

	if err != nil {
		return nil, err
	}

	if userNameExits {
		return nil, domain.ErrUsernameExists
	}

	emailExits, err := au.userRepository.CheckIfUserExists(ctx, "email", registration.Email)

	if err != nil {
		return nil, err
	}

	if emailExits {
		return nil, domain.ErrEmailExists
	}

	hashedPwd, err := au.hashing.Hash(registration.Password)

	if err != nil {
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	user := &domain.User{
		Username:  registration.Username,
		Email:     registration.Email,
		Password:  hashedPwd,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = au.userRepository.Save(ctx, user)
	log.Println(err)

	if err != nil {
		return nil, fmt.Errorf("user saving failed: %w", err)
	}

	return user, nil
}

func (au *AuthService) Validate(token string) (string, error) {
	id, err := au.token.Validate(token)

	if err != nil {
		return "", err
	}

	return id, nil
}
