package service

import (
	"context"

	"entexample/pkg/ent"
	"entexample/pkg/service/internal/repository"
)

type User interface {
	Create(ctx context.Context, name string) (*ent.User, error)
	Get(ctx context.Context, id int) (*ent.User, error)
}

type user struct {
	userRepository repository.User
}
