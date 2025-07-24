package db

import (
	"context"

	"github.com/sawmeraw/gogo/internal/store"
)

var usernames = []string{
	"bob", "alice",
}

func Seed(store store.Storage) error {
	ctx := context.Background()

	users := generateUsers(100)

	return nil
}

func generateUsers(n int) []*store.User {

	users := make([]*store.User, n)

	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username: "",
			Email:    "",
			Password: "",
		}
	}

	return users
}
