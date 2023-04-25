package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/cyansilver/go-lib/auth/token"
)

func HandleBearerAuth(
	r *http.Request,
	excludePath map[string]int8,
	verifyToken func(token string) (*token.SessionTokenClaims, error),
) error {
	if _, ok := excludePath[r.URL.Path]; ok {
		return nil
	}
	auth := r.Header.Get("authorization")
	if len(auth) == 0 {
		return errors.New("Missing authentication header")
	}
	const prefix = "Bearer "
	if !strings.HasPrefix(auth, prefix) {
		return errors.New("Missing prefix Bearer")
	}
	claims, err := verifyToken(auth[len(prefix):])
	if err != nil {
		return err
	}
	custProps, _ := json.Marshal(claims.CustProps)
	r.Header.Set("account-id", claims.UserID)
	r.Header.Set("account-username", claims.Username)
	r.Header.Set("account-props", string(custProps))

	return nil
}
