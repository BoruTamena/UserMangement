package services

import (
	"fmt"
	"os"
	"time"

	"github.com/BoruTamena/UserManagement/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(userclaim models.UserLogIn) (string, error) {

	// creating token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userclaim.UserName,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})

	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	token_str, err := token.SignedString([]byte(os.Getenv("SCERATEKEY")))

	if err != nil {

		return "", err

	}

	return token_str, nil

}

func ParseAccessToken(accessToken string) error {

	parse_token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {

		err := godotenv.Load(".env")
		if err != nil {
			return "", err
		}
		key := []byte(os.Getenv("SCERATEKEY"))

		return key, nil
	})

	if err != nil {
		return err
	}

	if !parse_token.Valid {

		return fmt.Errorf("Invalid token ")
	}
	return nil

}
