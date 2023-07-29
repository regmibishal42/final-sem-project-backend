package model

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//replace later with env
var (
	SecretKey = []byte("MyTopSecretKey")
)

const (
	ACCESS_TOKEN_EXPIRE = time.Hour * 24 * 100
)

func (u *User) _AuthToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ACCESS_TOKEN_EXPIRE).Unix()
	claims["token_type"] = "access"
	claims["UserId"] = u.ID
	claims["Email"] = u.Email

	authToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return authToken, nil
}

func (u *User) Token() (*AuthToken, error) {
	if u == nil {
		return nil, nil
	}
	if !u.IsVerified {
		return nil, nil
	}
	authToken, err := u._AuthToken()

	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: authToken,
	}, nil
}

//parse the jwt token
func ParseAuthToken(tokenString string) (*string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiresAt := claims["exp"].(float64)
		if float64(time.Now().Unix()) > expiresAt {
			return nil, errors.New("token expired")
		}
		userId := claims["userId"].(string)
		return &userId, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
