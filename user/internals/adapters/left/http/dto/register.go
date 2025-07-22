package dto

import "github.com/HlufD/users-ms/internals/domain"

// RegisterUserDto represents user registration data
// @Description User registration request payload
type RegisterUserDto struct {
	// @Example: "john_doe"
	Username string `json:"username" validate:"required,min=3,max=20" example:"john_doe" minLength:"3" maxLength:"20"`

	// @Example: "user@example.com"
	Email string `json:"email" validate:"required,email" example:"user@example.com" format:"email"`

	// @Example: "P@ssw0rd123"
	Password string `json:"password" validate:"required,min=6" example:"P@ssw0rd123" minLength:"6"`
}

// MapToDomainEntity converts DTO to domain model
func (ru *RegisterUserDto) MapToDomainEntity(user domain.Registration) *domain.Registration {
	return &domain.Registration{
		Username: ru.Username,
		Email:    ru.Email,
		Password: ru.Password,
	}
}
