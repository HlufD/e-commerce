package http

import (
	"encoding/json"
	"net/http"

	"github.com/HlufD/users-ms/common"
	"github.com/HlufD/users-ms/internals/adapters/left/http/dto"
	"github.com/HlufD/users-ms/internals/application"
	"github.com/HlufD/users-ms/internals/domain"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService application.AuthService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(authService application.AuthService) *AuthHandler {
	return &AuthHandler{
		authService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserDto true "Registration data"
// @Success 201 {object} domain.User "Successfully created user"
// @Failure 400 {object} map[string]interface{} "Invalid request format"
// @Failure 409 {object} map[string]interface{} "User already exists"
// @Router /api/v1/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.Registration
	var registerUserDto dto.RegisterUserDto

	if err := json.NewDecoder(r.Body).Decode(&registerUserDto); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := dto.Validate(registerUserDto); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user = *registerUserDto.MapToDomainEntity(user)
	createdUser, err := h.authService.Register(user)

	if err != nil {
		common.RespondWithError(w, http.StatusConflict, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, createdUser)
}

// Login godoc
// @Summary Authenticate user
// @Description Login with email and password
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dto.Login true "Login credentials"
// @Success 200 {object} domain.Token "Authentication token"
// @Failure 400 {object} map[string]interface{} "Invalid credentials"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/v1/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto dto.Login
	var creds domain.Credentials

	if err := json.NewDecoder(r.Body).Decode(&loginDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := dto.Validate(loginDto); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	creds = *loginDto.MapToEntity(creds)
	token, err := h.authService.Login(creds)

	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, token)
}

// Validate godoc
// @Summary      Validate user token
// @Description  Validates a user based on token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ValidateUser true "User Token Payload"
// @Success      200 {string} string "user_id"
// @Failure      400 {object} map[string]string "Bad Request"
// @Router       /api/v1/validate [post]
func (h *AuthHandler) Validate(w http.ResponseWriter, r *http.Request) {
	var validationDto dto.ValidateUser

	if err := json.NewDecoder(r.Body).Decode(&validationDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := dto.Validate(validationDto); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.authService.Validate(validationDto.Token)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response := map[string]string{
		"userId": id,
	}

	common.RespondWithJSON(w, http.StatusOK, response)
}
