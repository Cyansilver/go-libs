package token

import (
	"crypto"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SessionTokenClaims struct {
	UserID     string            `json:"uid,omitempty"`
	Username   string            `json:"username,omitempty"`
	IsVerified bool              `json:"isVerified,omitempty"`
	CustProps  map[string]string `json:"custProps,omitempty"`
	ExpiresAt  int64             `json:"exp,omitempty"`
}

func (stc *SessionTokenClaims) Valid() error {
	// Verify expiry.
	if stc.ExpiresAt <= time.Now().UTC().Unix() {
		vErr := new(jwt.ValidationError)
		vErr.Inner = errors.New("Token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
		return vErr
	}
	return nil
}

func Get(encryptionKey []byte, stc *SessionTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stc)
	return token.SignedString(encryptionKey)
}

func Parse(hmacSecretByte []byte, tokenString string) (*SessionTokenClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &SessionTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if s, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || s.Hash != crypto.SHA256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecretByte, nil
	})
	if err != nil {
		return nil, false
	}
	claims, ok := token.Claims.(*SessionTokenClaims)
	if !ok || !token.Valid {
		return nil, false
	}

	return claims, true
}
