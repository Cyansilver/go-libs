package auth

import (
	"encoding/json"
	"time"

	"github.com/cyansilver/go-lib/auth/token"
)

type AuthCaching struct {
	prefix string
	get    func(key string) (string, error)
	set    func(key string, val string, exp time.Duration) error
	verify func(token string) (*token.SessionTokenClaims, error)
}

// NewAuthCaching returns new instance of AuthCaching
func NewAuthCaching(
	prefix string,
	get func(key string) (string, error),
	set func(key string, val string, exp time.Duration) error,
	verify func(token string) (*token.SessionTokenClaims, error),
) *AuthCaching {
	return &AuthCaching{
		prefix: prefix,
		get:    get,
		set:    set,
		verify: verify,
	}
}

func (ac *AuthCaching) getKey(key string) string {
	return ac.prefix + key
}

// VerifyIDToken get from cache first
func (ac *AuthCaching) VerifyToken(tokenStr string) (*token.SessionTokenClaims, error) {
	key := ac.getKey(tokenStr)
	tokenClaimStr, err := ac.get(key)
	if err == nil {
		var tokenClaim token.SessionTokenClaims
		err = json.Unmarshal([]byte(tokenClaimStr), &tokenClaim)
		return &tokenClaim, err
	}

	tokenClaim, err := ac.verify(tokenStr)
	if err != nil {
		return tokenClaim, err
	}

	err = ac.set(key, tokenStr, time.Hour*4)

	return tokenClaim, err
}
