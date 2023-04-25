package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	t.Run("Hash and verify password successfully", func(t *testing.T) {
		// init
		password := "123456789"
		hashPwd, err := GenerateHashPassword(password)
		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}

		res, err := VerifyHashPassword(password, hashPwd)
		if res == false {
			t.Fatalf("Error %v", res)
		}

		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}
	})

	t.Run("Mismatch password", func(t *testing.T) {
		// init
		password := "123456789"
		wPassword := "1234567890"
		hashPwd, err := GenerateHashPassword(password)
		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}

		res, err := VerifyHashPassword(wPassword, hashPwd)

		// assert
		if res == true {
			t.Fatalf("Error %v", res)
		}
		// assert
		if err != nil {
			t.Fatalf("Error %v", err)
		}
	})
}
