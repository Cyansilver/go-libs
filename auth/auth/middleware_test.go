package auth

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cyansilver/go-libs/auth/token"
)

func TestHandleBearerAuth(t *testing.T) {
	t.Run("Exclude Path", func(t *testing.T) {
		// init
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{}
			return &st, errors.New("Test")
		}
		excludePath := map[string]int8{"/test": 1}
		req, err := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("Error %v", err)
		}

		err = HandleBearerAuth(req, excludePath, verifyToken)

		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}
	})

	t.Run("Missing authorization header", func(t *testing.T) {
		// init
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{}
			return &st, nil
		}
		excludePath := map[string]int8{}
		req, err := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("Error %v", err)
		}

		err = HandleBearerAuth(req, excludePath, verifyToken)
		// assert
		if err == nil {
			t.Fatal("Expected error but receive nil")
		}
	})

	t.Run("Missing prefix Bearer", func(t *testing.T) {
		// init
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{}
			return &st, nil
		}
		excludePath := map[string]int8{}
		req, err := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("Error %v", err)
		}
		req.Header.Add("authorization", "test")

		err = HandleBearerAuth(req, excludePath, verifyToken)
		// assert
		if err == nil {
			t.Fatal("Expected error but receive nil")
		}
	})

	t.Run("Failed to verify token", func(t *testing.T) {
		// init
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{}
			return &st, errors.New("Test")
		}
		excludePath := map[string]int8{}
		req, err := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("Error %v", err)
		}
		req.Header.Add("authorization", "Bearer test")

		err = HandleBearerAuth(req, excludePath, verifyToken)
		// assert
		if err == nil {
			t.Fatal("Expected error but receive nil")
		}
	})

	t.Run("Auth Success", func(t *testing.T) {
		// init
		verifyToken := func(tokenStr string) (*token.SessionTokenClaims, error) {
			st := token.SessionTokenClaims{}
			return &st, nil
		}
		excludePath := map[string]int8{}
		req, err := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Fatalf("Error %v", err)
		}
		req.Header.Add("authorization", "Bearer test")

		err = HandleBearerAuth(req, excludePath, verifyToken)
		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}
	})
}
