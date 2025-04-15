package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("A-Str1ng-s3cr3t-At-l3ast-256-b1ts-long")

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewJWTService() Service {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
