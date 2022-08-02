package auth

import (
	"context"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func InitFirebase() *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs(os.Getenv("FIREBASE_PRIVATE_KEY"))
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Fireauth load error")
	}
	return auth
}
