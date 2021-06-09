package repository

import (
	"context"

	"entexample/pkg/ent"
	userPredicate "entexample/pkg/ent/user"
)

type User interface {
	Create(ctx context.Context, name string) (*ent.User, error)
	CreateMany(ctx context.Context, names []string) ([]*ent.User, error)
	Get(ctx context.Context, id int) (*ent.User, error)
	GetFriends(ctx context.Context, user *ent.User) ([]*ent.User, error)
	GetName(ctx context.Context, id int) (string, error)
	Update(ctx context.Context, id int, name string) (*ent.User, error)
	Delete(ctx context.Context, user *ent.User) error
}

type user struct {
	client *ent.Client
}

func (r user) Create(ctx context.Context, name string) (*ent.User, error) {
	return r.client.User.Create().
		SetName(name).
		Save(ctx)
}

func (r user) Get(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}

func (r user) GetFriends(ctx context.Context, user *ent.User) ([]*ent.User, error) {
	return user.QueryFriend().All(ctx)
}

func (r user) GetName(ctx context.Context, id int) (string, error) {
	return r.client.User.Query().
		Where(userPredicate.IDEQ(id)).
		Select(userPredicate.FieldName).
		String(ctx)
}

func (r user) GetMany(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().
		Where(userPredicate.IDGT(2)).
		Order(ent.Asc(userPredicate.FieldID)).
		All(ctx)
}

func (r user) CreateMany(ctx context.Context, names []string) ([]*ent.User, error) {
	bulk := make([]*ent.UserCreate, 0)
	for _, name := range names {
		bulk = append(bulk, r.client.User.Create().SetName(name))
	}

	return r.client.User.CreateBulk(bulk...).Save(ctx)
}

func (r user) Update(ctx context.Context, id int, name string) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).SetName(name).Save(ctx)
}

func (r user) Delete(ctx context.Context, user *ent.User) error {
	return r.client.User.DeleteOne(user).Exec(ctx)
}
