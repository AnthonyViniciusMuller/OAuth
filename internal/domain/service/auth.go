package auth

import (
	"fmt"
	"math/rand/v2"
	"time"

	authcode "github.com/AnthonyViniciusMuller/OAuth/internal/domain/repository/auth_code"
	"github.com/AnthonyViniciusMuller/OAuth/internal/domain/repository/user"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	authcodeRepository authcode.Repository
	userRepository     user.Repository
}

func New(authcodeRepository authcode.Repository, userRepository user.Repository) *Service {
	return &Service{
		authcodeRepository: authcodeRepository,
		userRepository:     userRepository,
	}
}

func (service Service) Authorize(username, password string) (string, error) {
	// TODO: use password hash
	user, err := service.userRepository.GetByUserName(username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("%w for user %s", ErrPasswordMismatch, username)
	}

	// TODO: implement real code generation.
	authCode := fmt.Sprintf("%d", rand.Int())

	err = service.authcodeRepository.Insert(authCode, user.ID)
	if err != nil {
		return "", err
	}

	return authCode, nil
}

func (service Service) Token(authCode string) (string, error) {
	// TODO: remove auth code after used
	userID, err := service.authcodeRepository.GetUserID(authCode)
	if err != nil {
		return "", err
	}

	// TODO: get expiration from else where.
	expirationTime := time.Now().Add(5 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   fmt.Sprintf("%d", userID),
	})

	// TODO: get from a private.pem or something.
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
