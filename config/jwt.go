package config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func getSecret() string {
	secret := os.Getenv("secret")
	if secret == "" {
		return "secret"
	}
	return secret
}

type claims struct {
	UserName  string
	Email     string
	CanDelete bool
	jwt.StandardClaims
}

func GenerateToken(username string, email string, canDelete bool) string {
	claims := &claims{
		UserName:  username,
		Email:     email,
		CanDelete: canDelete,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "myIssuer",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(getSecret()))

	if err != nil {
		panic(err)
	}

	return t
}

func ValidateToken(token string) bool {
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token")

		}
		return []byte(getSecret()), nil
	})

	if err != nil {
		return false
	} else {
		return res.Valid
	}
}
