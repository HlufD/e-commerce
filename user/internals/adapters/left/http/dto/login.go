package dto

import "github.com/HlufD/users-ms/internals/domain"

// Login represents user authentication credentials
// @Description User login request payload
type Login struct {
	// @Example: "john_doe"
	Username string `json:"username" validate:"required,min=3,max=20" example:"john_doe" minLength:"3" maxLength:"20"`

	// @Example: "P@ssw0rd123"
	Password string `json:"password" validate:"required,min=5" example:"P@ssw0rd123" minLength:"5"`
}

// MapToEntity converts Login DTO to domain Credentials
// @Summary Converts login DTO to domain model
// @Description Maps the login data transfer object to domain credentials
func (lg *Login) MapToEntity(credentials domain.Credentials) *domain.Credentials {
	return &domain.Credentials{
		Username: lg.Username,
		Password: lg.Password,
	}
}
