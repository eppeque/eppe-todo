package security

import (
	"errors"
	"time"

	"github.com/eppeque/todo-server/models"
	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func CreateToken(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(models.ServerConfig.SecretKey)
	signed, _ := token.SignedString(secretKey)
	return signed
}

func VerifyToken(token string) (int, error) {
	claims := &customClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.ServerConfig.SecretKey), nil
	})

	if err != nil {
		return -1, err
	}

	if parsedToken.Valid {
		return claims.Id, nil
	}

	return -1, errors.New("invalid claims")
}
