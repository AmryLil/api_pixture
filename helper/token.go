package helper

import (
	models "api/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var MyKey = []byte("kunciku")

type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(models models.UserAccounts) (string, error) {
	claims := JWTClaims{
		models.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MyKey)
	return ss, err

}

func ValidateToken(tokenstr string) (*int, error) {
	token, err := jwt.ParseWithClaims(tokenstr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("your token was expired")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("your token was expired")
	}
	return &claims.ID, nil
}
