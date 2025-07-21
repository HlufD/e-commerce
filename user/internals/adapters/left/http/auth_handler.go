package http

import (
	"encoding/json"
	"net/http"

	"github.com/HlufD/users-ms/common"
	"github.com/HlufD/users-ms/internals/application"
	"github.com/HlufD/users-ms/internals/domain"
)

type AuthHandler struct {
	authService application.AuthService
}

func NewAuthHandler(authService application.AuthService) *AuthHandler {
	return &AuthHandler{
		authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	user := domain.Registration{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdUser, err := h.authService.Register(user)

	if err != nil {
		common.RespondWithError(w, http.StatusConflict, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, createdUser)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var creds domain.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(creds)

	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, token)
}
