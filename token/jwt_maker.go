package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

type JwtMaker struct {
	secret string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size must be atleast %d characters", minSecretKeySize)
	}

	return &JwtMaker{secret: secretKey}, nil
}

func (maker *JwtMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secret))
}
func (maker *JwtMaker) VerifyToken(token string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, tokenInvalidErr
		}
		return []byte(maker.secret), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, tokenExpiredErr) {
			return nil, tokenExpiredErr
		}
		return nil, tokenInvalidErr
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, tokenInvalidErr
	}

	return payload, nil

}
