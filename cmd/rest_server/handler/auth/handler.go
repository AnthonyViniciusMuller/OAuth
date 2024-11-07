package auth

import (
	"encoding/json"
	"net/http"

	authCodeService "github.com/AnthonyViniciusMuller/OAuth/internal/domain/service"
	authCodeRepository "github.com/AnthonyViniciusMuller/OAuth/internal/infrastructure/repository/auth_code"
	userRepository "github.com/AnthonyViniciusMuller/OAuth/internal/infrastructure/repository/user"
)

type Handler struct {
	service *authCodeService.Service
}

func New() *Handler {
	return &Handler{
		service: authCodeService.New(authCodeRepository.New(), userRepository.New()),
	}
}

func (handler *Handler) Authorize(w http.ResponseWriter, r *http.Request) {
	username, password := r.URL.Query().Get("username"), r.URL.Query().Get("password")

	authCode, err := handler.service.Authorize(username, password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"authorization_code": authCode})
}

func (handler *Handler) Token(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	token, err := handler.service.Token(authCode)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}
