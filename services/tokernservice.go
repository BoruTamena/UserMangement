package services

import (
	"fmt"
	"os"
	"time"

	"github.com/BoruTamena/UserManagement/models"
	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpireDuration  = time.Minute * 15
	refreshTokenExpireDuration = time.Hour * 24 * 7
)

func CreateToken(userClaim models.UserReg) (string, string, error) {
	// Creating access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userClaim.Id, // Change "UserId" to "userId" for consistency
		"exp":    time.Now().Add(accessTokenExpireDuration).Unix(),
	})

	// Creating refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userClaim.Id,
		"exp":    time.Now().Add(refreshTokenExpireDuration).Unix(),
	})

	// Signing tokens
	accessTokenStr, err := accessToken.SignedString([]byte(os.Getenv("SCERATEKEY")))
	if err != nil {
		return "", "", err
	}

	refreshTokenStr, err := refreshToken.SignedString([]byte(os.Getenv("SCERATEKEY")))
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(accessTokenExpireDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SCERATEKEY")))
}

func ParseAccessToken(accessToken string) error {
	return parseToken(accessToken)
}

func ParseRefreshToken(refreshToken string) error {
	return parseToken(refreshToken)
}

func parseToken(tokenString string) error {
	// Parse token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCERATEKEY")), nil
	})

	// Check for parsing errors
	if err != nil {
		return err
	}

	// Check token validity
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
