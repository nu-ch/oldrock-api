package auth

import "context"

type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	login(ctx context.Context, username, password string) (string, error)
}

func login() {}
