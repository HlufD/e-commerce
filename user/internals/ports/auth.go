package ports

import "github.com/HlufD/users-ms/internals/domain"

type AuthPort interface {
	Login(credentials domain.Credentials) (domain.Token, error)
	Register(registration domain.Registration) (*domain.User, error)
}
