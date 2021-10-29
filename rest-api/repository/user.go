package repository

import (
	"context"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(user *ent.User) (*ent.User, error)
	FindById(id int) (*ent.User, error)
	Update(id int, user *ent.User) (*ent.User, error)
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{
		db: client,
	}
}

type userRepository struct {
	db *ent.Client
}

func (repo *userRepository) Create(user *ent.User) (*ent.User, error) {
	result, err := repo.db.User.Create().
		SetNickname(user.Nickname).
		SetMotto(user.Motto).
		SetKind(user.Kind).
		SetNillableHeight(user.Height).
		SetNillableWeight(user.Weight).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (repo *userRepository) FindById(id int) (*ent.User, error) {
	u, err := repo.db.User.Query().
		Where(user.ID(id)).
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func (repo *userRepository) Update(id int, user *ent.User) (*ent.User, error) {
	result, err := repo.db.User.UpdateOneID(id).
		SetNickname(user.Nickname).
		SetMotto(user.Motto).
		SetNillableHeight(user.Height).
		SetNillableWeight(user.Weight).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
