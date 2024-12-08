package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

type JWTPayloadClaims struct {
	Payload
	jwt.RegisteredClaims
}

func NewJWTPayloadClaims(payload *Payload) *JWTPayloadClaims {
	return &JWTPayloadClaims{
		Payload: *payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
			NotBefore: jwt.NewNumericDate(payload.IssuedAt),
			Issuer:    "simplebank",
			Subject:   payload.Username,
			ID:        payload.ID.String(),
			Audience:  []string{"clients"},
		},
	}
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, NewJWTPayloadClaims(payload))
	signedString, err := jwtToken.SignedString([]byte(m.secretKey))
	return signedString, err
}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtClaims := &JWTPayloadClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, jwtClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		} else if errors.Is(err, ErrInvalidToken) {
			return nil, ErrInvalidToken
		} else {
			return nil, err
		}
	}

	payloadClaims, ok := jwtToken.Claims.(*JWTPayloadClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &payloadClaims.Payload, nil
}
