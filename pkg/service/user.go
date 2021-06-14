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
	AddFriend(ctx context.Context, user, friend *ent.User) error
}

type user struct {
	userRepository repository.User
}

func NewUser(client *ent.Client) User {
	return user{userRepository: repository.NewUser(client)}
}

func (u user) Create(ctx context.Context, name string) (*ent.User, error) {
	return u.userRepository.Create(ctx, name)
}

func (u user) Get(ctx context.Context, id int) (*ent.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u user) CreateMany(ctx context.Context, names []string) ([]*ent.User, error) {
	return u.userRepository.CreateMany(ctx, names)
}

func (u user) GetFriends(ctx context.Context, user *ent.User) ([]*ent.User, error) {
	return u.userRepository.GetFriends(ctx, user)
}

func (u user) GetName(ctx context.Context, id int) (string, error) {
	return u.userRepository.GetName(ctx, id)
}

func (u user) Update(ctx context.Context, id int, name string) (*ent.User, error) {
	return u.userRepository.Update(ctx, id, name)
}

func (u user) Delete(ctx context.Context, user *ent.User) error {
	return u.userRepository.Delete(ctx, user)
}

func (u user) DeleteFriends(ctx context.Context, user *ent.User) error {
	return u.userRepository.DeleteFriendsTx(ctx, user)
}

func (u user) AddFriend(ctx context.Context, user, friend *ent.User) error {
	return u.userRepository.AddFriend(ctx, user, friend)
}
