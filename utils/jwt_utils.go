package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"goLangJwtPrac/structures"
	"time"
)

var (
	secretKey = []byte("cos")
)

type Claims struct {
	Username    string `json:"username"`
	Authorities string `json:"authorities"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *structures.User) (string, error) {
	var claims Claims
	claims.Username = user.Username
	claims.Authorities = user.Authorities
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, err
}

func CreateRefreshToken(user *structures.User) (string, error) {
	var claims Claims
	claims.Username = user.Username
	claims.Authorities = user.Authorities
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * 7 * 2 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return refreshToken, err
}

func VerifyAccessToken(accessToken string) (bool, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["Username"], claims["Authorities"])
	} else {
		return false, fmt.Errorf("invalid token")
	}
	parseResult, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("cos"), nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := parseResult.Claims.(*Claims); ok && parseResult.Valid {
		fmt.Print(claims)
	} else {
		fmt.Println(err)
	}
	return true, nil
}
