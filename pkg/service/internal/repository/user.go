package repository

import (
	"context"
	"log"

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
	DeleteFriendsTx(ctx context.Context, tx *ent.Tx, user *ent.User) error
}

type user struct {
	client *ent.Client
}

type UserQueryOption func(query *ent.UserQuery)

func UserWithFriends() UserQueryOption {
	return func(query *ent.UserQuery) {
		query.Where(userPredicate.HasFriend())
	}
}

func UserWithIDGT(id int) UserQueryOption {
	return func(query *ent.UserQuery) {
		query.Where(userPredicate.IDGT(id))
	}
}

func UserOrderedBy(field string) UserQueryOption {
	return func(query *ent.UserQuery) {
		query.Order(ent.Desc(field))
	}
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

func (r user) DeleteFriendsTx(ctx context.Context, tx *ent.Tx, user *ent.User) error {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := tx.User.UpdateOne(user).ClearFriend().Exec(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Print(err)
		}

		return err
	}

	return tx.Commit()
}

func (r user) GetMany(ctx context.Context, options ...UserQueryOption) ([]*ent.User, error) {
	query := r.client.User.Query()
	for _, option := range options {
		option(query)
	}

	return query.All(ctx)
}
