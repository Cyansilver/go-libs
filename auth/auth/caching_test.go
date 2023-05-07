package auth

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/cyansilver/go-libs/auth/token"
)

func TestAuthCaching(t *testing.T) {
	t.Run("Verify from cache", func(t *testing.T) {
		// init
		prefix := "test_"
		userId := "test1243"
		get := func(key string) (string, error) {
			val, _ := json.Marshal(token.SessionTokenClaims{UserID: userId})
			return string(val), nil
		}
		set := func(key string, val string, exp time.Duration) error {
			return nil
		}
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{UserID: userId}
			return &st, errors.New("Test")
		}

		svc := NewAuthCaching(prefix, get, set, verifyToken)
		tokenClaim, err := svc.VerifyToken(userId)

		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}
		if tokenClaim.UserID != userId {
			t.Fatalf("Expected %v, actual %v", userId, tokenClaim.UserID)
		}
	})

	t.Run("Verify from verify function", func(t *testing.T) {
		// init
		prefix := "test_"
		userId := "test1243"
		get := func(key string) (string, error) {
			val, _ := json.Marshal(token.SessionTokenClaims{UserID: userId})
			return string(val), errors.New("Not found")
		}
		set := func(key string, val string, exp time.Duration) error {
			return nil
		}
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{UserID: userId}
			return &st, nil
		}

		svc := NewAuthCaching(prefix, get, set, verifyToken)
		tokenClaim, err := svc.VerifyToken(userId)

		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}
		if tokenClaim.UserID != userId {
			t.Fatalf("Expected %v, actual %v", userId, tokenClaim.UserID)
		}
	})

	t.Run("Verify from verify function but store failed", func(t *testing.T) {
		// init
		prefix := "test_"
		userId := "test1243"
		get := func(key string) (string, error) {
			val, _ := json.Marshal(token.SessionTokenClaims{UserID: userId})
			return string(val), errors.New("Not found")
		}
		set := func(key string, val string, exp time.Duration) error {
			return errors.New("Failed to store")
		}
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{UserID: userId}
			return &st, nil
		}

		svc := NewAuthCaching(prefix, get, set, verifyToken)
		tokenClaim, err := svc.VerifyToken(userId)

		// assert
		if err == nil {
			t.Fatal("Expected errror but nil")
		}
		if tokenClaim.UserID != userId {
			t.Fatalf("Expected %v, actual %v", userId, tokenClaim.UserID)
		}
	})
}
