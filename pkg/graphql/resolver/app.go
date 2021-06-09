package resolver

import (
	"entexample/pkg/graphql/gqlgen"

	"entgo.io/ent/examples/start/ent"
)

func New(client *ent.Client) *Resolver {
	return &Resolver{
		client: client,
	}
}

type Resolver struct{ client *ent.Client }

func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}
