package fixture

import (
	"context"

	"entexample/pkg/ent"
)

type UserOption func(*ent.User)

func UserWithName(name string) UserOption {
	return func(user *ent.User) {
		user.Name = name
	}
}

func NewUser(client *ent.Client, options ...UserOption) *ent.User {
	user := &ent.User{}

	for _, option := range options {
		option(user)
	}

	return client.User.Create().
		SetName(user.Name).
		SaveX(context.Background())
}
