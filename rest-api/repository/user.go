package repository

import (
	"context"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

type UserRepository interface {
	Create(user *ent.User) (*ent.User, error)
	FindById(id int) (*ent.User, error)
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
		SetType(user.Type).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *userRepository) FindById(id int) (*ent.User, error) {
	u, err := repo.db.User.Query().
		Where(user.ID(id)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}
