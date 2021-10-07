package repository

import (
	"context"

	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

type albumRepository struct {
	db *ent.Client
}

type AlbumRepositoryInterface interface {
	Create(album *ent.Album) (*ent.Album, error)
	GetAllAlbumsByUserID(userID string) ([]*ent.Album, error)
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
		SetFolderName(album.FolderName).
		SetCreatedAt(album.CreatedAt).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}

func (r *albumRepository) GetAllAlbumsByUserID(userID string) ([]*ent.Album, error) {
	albums, err := r.db.Album.Query().
		Where(album.HasUserWith(user.ID(userID))).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// Get은 각 album의 picture 정보도 다 조회해야 함
func (r *albumRepository) Get(albumID int) (*ent.Album, error) {
	albumData, err := r.db.Album.Query().
		Where(album.ID(albumID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return albumData, nil
}

// func (r *albumRepository) FindByIDAndUserID(userID string, albumID int) (*ent.Album, error) {
// 	albumData, err := r.db.Album.Query().
// 		Where(album.And(album.HasUserWith(user.ID(userID)), album.ID(albumID))).
// 		Only(context.Background())

// 	if err != nil {
// 		return nil, err
// 	}

// 	return albumData, nil
// }
