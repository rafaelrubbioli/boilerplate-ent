package resolver

import (
	"context"

	"entexample/pkg/graphql/model"
)

type queryResolver struct{ *Resolver }

func (q queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic("implement me")
}
