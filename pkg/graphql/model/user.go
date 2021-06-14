package model

import "entexample/pkg/ent"

func NewUser(user *ent.User) *User {
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}
