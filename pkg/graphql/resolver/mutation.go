package resolver

import (
	"context"
	"log"

	"entexample/pkg/graphql/gqlerror"
	"entexample/pkg/graphql/model"
)

type mutationResolver struct{ *Resolver }

func (m mutationResolver) AddFriend(ctx context.Context, id int, friendID int) (*model.User, error) {
	user, err := m.userService.Get(ctx, id)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	friend, err := m.userService.Get(ctx, friendID)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	err = m.userService.AddFriend(ctx, user, friend)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	return model.NewUser(user), nil
}

func (m mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	user, err := m.userService.Create(ctx, name)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	return model.NewUser(user), nil
}

func (m mutationResolver) CreateManyUsers(ctx context.Context, names []string) ([]*model.User, error) {
	users, err := m.userService.CreateMany(ctx, names)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	results := make([]*model.User, 0)
	for _, user := range users {
		results = append(results, model.NewUser(user))
	}

	return results, nil
}

func (m mutationResolver) UpdateUser(ctx context.Context, id int, newName string) (*model.User, error) {
	user, err := m.userService.Update(ctx, id, newName)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	return model.NewUser(user), nil
}

func (m mutationResolver) DeleteUserFriends(ctx context.Context, id int) (int, error) {
	user, err := m.userService.Get(ctx, id)
	if err != nil {
		log.Print(err)
		return 0, gqlerror.ErrServiceUnavailable
	}

	err = m.userService.DeleteFriends(ctx, user)
	if err != nil {
		log.Print(err)
		return 0, gqlerror.ErrServiceUnavailable
	}

	return id, nil
}

func (m mutationResolver) DeleteUser(ctx context.Context, id int) (int, error) {
	user, err := m.userService.Get(ctx, id)
	if err != nil {
		log.Print(err)
		return 0, gqlerror.ErrServiceUnavailable
	}

	err = m.userService.Delete(ctx, user)
	if err != nil {
		log.Print(err)
		return 0, gqlerror.ErrServiceUnavailable
	}

	return id, nil
}
