package utilities

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/seigaalghi/majoo-pos/models"
)

var _ error = godotenv.Load()
var secret string = os.Getenv("SECRET_TOKEN")

func GenerateToken(payload *models.Payload) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["payload"] = payload
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return token, err
}
