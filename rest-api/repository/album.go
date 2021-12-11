package repository

import (
	"context"

	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/pkg/errors"
)

type albumRepository struct {
	db *ent.Client
}

type AlbumRepositoryInterface interface {
	Create(album *ent.Album) (*ent.Album, error)
	GetAllByUserID(userID int) ([]*ent.Album, error)
	Get(albumID int) (*ent.Album, error)
	Update(albumID int, album *ent.Album) (*ent.Album, error)
	Delete(albumID int) error
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
	// TODO: 이 부분에서 근데 EagerLoading 안 잡아주면 에러날 것 같은데

	return newAlbum, nil
}

func (r *albumRepository) GetAllByUserID(userID int) ([]*ent.Album, error) {
	albums, err := r.db.Album.Query().
		Where(album.HasUserWith(user.ID(userID))).
		WithUser().
		WithPicture(func(query *ent.PictureQuery) {
			query.WithAlbum(func(query *ent.AlbumQuery) {
				query.Select("id")
			}).
				WithUser().
				Order(ent.Asc("taken_at", "created_at"))
		}).
		Order(ent.Desc("created_at")).
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return albums, nil
}

func (r *albumRepository) Get(albumID int) (*ent.Album, error) {
	albumData, err := r.db.Album.Query().
		Where(album.ID(albumID)).
		WithUser().
		WithPicture(func(query *ent.PictureQuery) {
			query.WithAlbum(func(query *ent.AlbumQuery) {
				query.Select("id")
			}).
				WithUser().
				Order(ent.Asc("taken_at", "created_at"))
		}).
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return albumData, nil
}

func (r *albumRepository) Update(id int, a *ent.Album) (*ent.Album, error) {
	updated, err := r.db.Album.UpdateOneID(id).
		SetName(a.Name).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return updated, nil
}

func (r *albumRepository) Delete(id int) error {
	err := r.db.Album.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
