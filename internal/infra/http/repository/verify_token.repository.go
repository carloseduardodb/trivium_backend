package repository

import (
	"context"
	"fmt"
	"log"
	"trivium/internal/presentation/dto"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type VerifyTokenRepository struct {
	authClient *auth.Client
}

var config *VerifyTokenRepository

func NewVerifyTokenRepository() (*VerifyTokenRepository, error) {
	opt := option.WithCredentialsFile("config/gloveritas-apps-firebase-adminsdk-nn1lz-fbde4c3002.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase Auth client: %v", err)
	}

	fmt.Println("Firebase Auth client initialized")

	config = &VerifyTokenRepository{
		authClient: authClient,
	}

	return &VerifyTokenRepository{
		authClient: authClient,
	}, nil
}

func (f *VerifyTokenRepository) VerifyIdToken(ctx context.Context, token string) (dto.User, error) {
	tokenInfo, err := f.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return dto.User{}, fmt.Errorf("error verifying ID token: %v", err)
	}

	user := dto.User{
		Name:      tokenInfo.Claims["name"].(string),
		Email:     tokenInfo.Claims["email"].(string),
		PhotoPath: tokenInfo.Claims["picture"].(string),
	}

	return user, nil
}

func GetAuthClient() *auth.Client {
	if config == nil || config.authClient == nil {
		var err error
		config, err = NewVerifyTokenRepository()
		if err != nil {
			log.Fatal(err)
		}
	}
	return config.authClient
}
