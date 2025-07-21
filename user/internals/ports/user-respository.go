package ports

import (
	"context"

	"github.com/HlufD/users-ms/internals/domain"
)

type UserRepositoryPort interface {
	Save(ctx context.Context, user *domain.User) (*domain.User, error)
	FindById(ctx context.Context, id string) (*domain.User, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	CheckIfUserExists(ctx context.Context, filterKye, filterValue string) (bool, error)
}
