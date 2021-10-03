package repository

import (
	"context"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type albumRepository struct {
	db *ent.Client
}

type AlbumRepositoryInterface interface {
	Create(album *ent.Album) (*ent.Album, error)
}

func NewAlbumRepository(db *ent.Client) AlbumRepositoryInterface {
	return &albumRepository{
		db: db,
	}
}

func (r *albumRepository) Create(album *ent.Album) (*ent.Album, error) {
	newAlbum, err := r.db.Album.Create().
		SetUser(album.Edges.User).
		SetFolderName(album.FolderName).
		SetCreatedAt(album.CreatedAt).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}
