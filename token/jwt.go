package token

import (
	"errors"
	"fmt"
	"time"

	"beebeewijaya.com/util"
	"github.com/golang-jwt/jwt/v4"
)

const minSecretKeyLength = 16

type JWTMaker struct {
	secretKey string
}

func NewMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeyLength {
		return nil, fmt.Errorf("error: secretkey length is less than %d", minSecretKeyLength)
	}

	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

func (j *JWTMaker) GenerateToken(username int64, email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, email, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, util.ErrTokenInvalid
		}

		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verified, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verified.Inner, util.ErrTokenExpired) {
			return nil, util.ErrTokenExpired
		}
		return nil, util.ErrTokenInvalid
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, util.ErrTokenInvalid
	}

	return payload, nil
}
