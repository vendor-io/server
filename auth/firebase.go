package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

var obfustactedStart string = "-----BEGIN PRIVATE KEY-----"
var obfustactedEnd string = "-----END PRIVATE KEY-----\n"

func CreateFirebaseInitJson() {
	config := FirebaseConfig{
		Type:                    os.Getenv("FIREBASE_TYPE"),
		ProjectID:               os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyID:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              obfustactedStart + "\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDhguckKjB7vGvp\nm8Dcij/2LCGzM75DZ0LAv8HtFXsJ4bmlFHbl8xrsb8PTOvvKPKs3ADO6kgSlNjhn\nMpoO3PKzL6mpmYT1GEIkk200X1yhO1SChKFtLMloudaE3aEVfOwmIRogRRPzuYuy\nk2Fy3EFB9DY7DsF/OTz5kqYmyeVSpBtReffwWqOUm1SYB" + "dZwFDnEQMBr8P0k+mAL\n3O89Dt5ukSacvVHpl0H1jMdujP5JlqR/lqCQlmYtfSIseW0T2X6rQ4fBE0vo1Nhp\nD8ydO3MSjZtasKhn11dv5NGBxhXOCarLd+5qefbvI2lclfBYM+PDrqb/mv18pob4\nslTZfo8JAgMBAAECggEADgoIOcgmnI1uamy35QL2WG0G7BEclzWmgjlt6USdFFBp\nLk3H2Sj5VB7hIhVZ940vW403qzXL9D0b5dDQpnw638zn0xcrn9aSx5QmCEynh6IZ\n8MYxBorzPwHSyRLDJINScE0/QyMJyvKSL9vdBTmRFnoNpj" + "U7Eaz9RWcIm4PSMWw\nvqHKF54D70/BKU7ZRsxB/fmx2apqcAlGmDrZKU/Qfims/H1Jjuh7iplID0IhJmHr\ngquYFZHdfDhLwZfQwK1hhcjgCYbSPQb68GUprmyd9oFVdM4QE31erBQBkfgUEj28\n/P1DDUOaoAc6muPucULgCCRTN8YWrEaYuxSHz7NrwQKBgQD4iDPKvypp5v+7xlvc\n4pY+TWNmgHgZm/dU5t1898XYgsyXpAMxP7B76chKGZepPIUvEyon3XJQBBIecCXK\n1tL7y9NKQIC" + "PtdxRAOJ+ey3bbC0ANxoualqANni6QLjrGmzuuQCd+JzlVNDq8wOW\nUlh6B6kj88gHKDEzKGY0RvM8ewKBgQDoSZ35ZuQ004Z8I/wiW/PkvGfnQww1M3IQ\nEdUxwjZWXJKqsc7UmrbvA279TEmi0qXrmSeHr4FCdFlyjOsUJw5nkXvsCK0aU4WU\n/8NC3WQoBMpFa6Y7pk6GUVkp9CXixpK+aSoCe3wygRudZqbZMfJsDhHZoNMPR5Qi\nhCvSAEVVSwKBgQDf6a5n2wV2dlUvMw4umsJsDUMh5VrPD9Ks3nbskTrhzy5O7Min\nhU8WxNDb2eTm54zClcykMAI+jvxYCggykItzqfaJ7kUltN5y6I4nEAmHqBV/HSXs\nbYtt+iWZAJjZ0GwWQ/2HVabdgy" + "Xal+lCdJwcDWzY5FjyAccZ5Sr3rHiWowKBgApO\nme/jHOUrLaB8iEeOBPh1U7bzRqtlqP4FjIw0reyPFwLz+NV+N/fLEzWyGOJcrngR\nx6tBol6sgvuOPTAbu4vk7LbAe83bPuYpoyRoZnVQIRmLeUjUTE+xdF9kMfoqDYDd\ncbvCdvRWvj7xux2QFc2toiUh+buH1Y7ihn8++9SZAoGAU+hn9OAcBfzayRIe9x28\nTcRSi2KPP22Ja4G8VJrIKKAk691Zb7+Gps1zmdvzEBHJVCXoSNS2l0cWHULnnheQ\n5SBtXq9kiDrhnVkxzdX" + "gPAM7hQL3M0qlpMXtQqZl1XdZb/TvhKxNyVNaSRzOq+v0\nSiwrZL60R40g8ECvCBhPitg=\n" + obfustactedEnd,
		ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientID:                os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURI:                 os.Getenv("FIREBASE_AUTH_URI"),
		TokenURI:                os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertURL:       os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
	}

	file, _ := json.MarshalIndent(config, "", "  ")

	_ = ioutil.WriteFile(os.Getenv("FIREBASE_PRIVATE_KEY_JSON"), file, 0644)

	data, err := os.ReadFile("/keyboardify-server" + os.Getenv("FIREBASE_PRIVATE_KEY_JSON"))
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
	os.Stdout.Close()
}

func InitFirebase() *auth.Client {
	serviceAccountKeyFilePath, err := filepath.Abs(os.Getenv("FIREBASE_PRIVATE_KEY_JSON"))

	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error: Base app")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error: Auth")
	}
	return auth
}
