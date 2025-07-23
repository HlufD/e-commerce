package dto

type ValidateUser struct {
	Token string `json:"token" validate:"required"`
}
