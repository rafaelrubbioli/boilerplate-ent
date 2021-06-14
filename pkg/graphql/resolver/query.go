package resolver

import (
	"context"
	"log"

	"entexample/pkg/graphql/gqlerror"
	"entexample/pkg/graphql/model"
)

type queryResolver struct{ *Resolver }

func (q queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user, err := q.userService.Get(ctx, id)
	if err != nil {
		log.Print(err)
		return nil, gqlerror.ErrServiceUnavailable
	}

	return model.NewUser(user), nil
}
