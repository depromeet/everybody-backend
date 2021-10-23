package repository

import (
	"context"
	"github.com/pkg/errors"

	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

type albumRepository struct {
	db *ent.Client
}

type AlbumRepositoryInterface interface {
	Create(album *ent.Album) (*ent.Album, error)
	GetAllByUserID(userID int) ([]*ent.Album, error)
	Get(albumID int) (*ent.Album, error)
}

func NewAlbumRepository(db *ent.Client) AlbumRepositoryInterface {
	return &albumRepository{
		db: db,
	}
}

func (r *albumRepository) Create(album *ent.Album) (*ent.Album, error) {
	newAlbum, err := r.db.Album.Create().
		SetUser(album.Edges.User).
		SetName(album.Name).
		Save(context.Background())

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return newAlbum, nil
}

func (r *albumRepository) GetAllByUserID(userID int) ([]*ent.Album, error) {
	albums, err := r.db.Album.Query().
		Where(album.HasUserWith(user.ID(userID))).
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return albums, nil
}

func (r *albumRepository) Get(albumID int) (*ent.Album, error) {
	albumData, err := r.db.Album.Query().
		Where(album.ID(albumID)).
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return albumData, nil
}
