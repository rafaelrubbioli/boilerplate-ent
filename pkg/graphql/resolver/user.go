package resolver

import (
	"context"
	"log"

	"entexample/pkg/graphql/gqlerror"
	"entexample/pkg/graphql/model"
)

type userResolver struct{ *Resolver }

func (u userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	user, err := u.userService.Get(ctx, obj.ID)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	friends, err := u.userService.GetFriends(ctx, user)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	results := make([]*model.User, 0)
	for _, friend := range friends {
		results = append(results, model.NewUser(friend))
	}

	return results, nil
}
