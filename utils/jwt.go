package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	ID string `json:"id"`
	Rut string `json:"rut"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id string, rut string, email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		ID:	id,
		Rut: rut,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(getJWTSecret()))
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getJWTSecret()), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("no se pudo leer los claims")
		return
	}
	if claims.ExpiresAt.Unix() < jwt.NewNumericDate(time.Now().Local()).Unix() {
		err = errors.New("token venció")
		return
	}
	return
}

func getJWTSecret() string {
	jwtSecret, exists := os.LookupEnv("JWT_SECRET")
    if !exists {
        jwtSecret = "jwtsecret"
    }
	return jwtSecret
}
