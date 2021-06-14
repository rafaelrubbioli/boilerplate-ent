package service

import (
	"context"

	"entexample/pkg/ent"
	"entexample/pkg/service/internal/repository"
)

type User interface {
	Create(ctx context.Context, name string) (*ent.User, error)
	Get(ctx context.Context, id int) (*ent.User, error)
	CreateMany(ctx context.Context, names []string) ([]*ent.User, error)
	GetFriends(ctx context.Context, user *ent.User) ([]*ent.User, error)
	GetName(ctx context.Context, id int) (string, error)
	Update(ctx context.Context, id int, name string) (*ent.User, error)
	Delete(ctx context.Context, user *ent.User) error
	DeleteFriends(ctx context.Context, user *ent.User) error
}

type user struct {
	userRepository repository.User
}
