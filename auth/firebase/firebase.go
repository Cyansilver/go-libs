package firebase

import (
	"context"
	"errors"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"github.com/cyansilver/go-libs/auth/token"
)

// AuthFirebase wraps firebase function
type AuthFirebase struct {
	app    *firebase.App
	Client *auth.Client
}

// NewAuthFirebase returns new instance of AuthFirebase
func NewAuthFirebase(cfgFile string) *AuthFirebase {
	opt := option.WithCredentialsFile(cfgFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return nil
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("error initializing client: %v\n", err)
		return nil
	}

	return &AuthFirebase{
		app:    app,
		Client: client,
	}
}

// CreateUser create a new user
func (af *AuthFirebase) CreateUser(
	uid string,
	username string,
	password string,
) error {
	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(username).
		EmailVerified(false).
		Password(password).
		Disabled(false)
	_, err := af.Client.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

// VerifyToken uses firebase admin sdk to verify id token
func (af *AuthFirebase) VerifyToken(idToken string) (*token.SessionTokenClaims, error) {
	tokenFb, err := af.Client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			return nil, errors.New("The authentication failed. Please sign in again")
		}
		return nil, err
	}

	userInfo, err := af.GetUser(tokenFb.UID)
	if err != nil {
		return nil, err
	}
	tokenClaim := token.SessionTokenClaims{
		UserID:   tokenFb.UID,
		Username: userInfo["email"],
		CustProps: map[string]string{
			"name":        userInfo["name"],
			"createdAt":   userInfo["createdAt"],
			"lastLogInTS": userInfo["lastLogInTS"],
			"avatarURL":   userInfo["avatarURL"],
		},
		IsVerified: true,
	}

	return &tokenClaim, nil
}

// GetUser get user info from firebase
func (af *AuthFirebase) GetUser(userID string) (map[string]string, error) {
	user, err := af.Client.GetUser(context.Background(), userID)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"email":       user.Email,
		"name":        user.DisplayName,
		"createdAt":   fmt.Sprintf("%d", user.UserMetadata.CreationTimestamp),
		"avatarURL":   user.PhotoURL,
		"lastLogInTS": fmt.Sprintf("%d", user.UserMetadata.LastLogInTimestamp),
	}, nil
}
