package resolver

import (
	"entexample/pkg/ent"
	"entexample/pkg/graphql/gqlgen"
	"entexample/pkg/service"
)

func New(client *ent.Client) *Resolver {
	return &Resolver{
		userService: service.NewUser(client),
	}
}

type Resolver struct{ userService service.User }

func (r *Resolver) User() gqlgen.UserResolver {
	return userResolver{r}
}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return mutationResolver{r}
}

func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}
