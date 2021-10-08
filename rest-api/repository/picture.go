package repository

import (
	"context"

	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/picture"
)

type pictureRepository struct {
	db *ent.Client
}

type PictureRepositoryInterface interface {
	Save(picture *ent.Picture) error
	GetAllByAlbumID(albumID int) ([]*ent.Picture, error)
	Get(pictureID int) (*ent.Picture, error)
	FindByAlbumIDAndBodyPart(albumID int, bodyPart string) ([]*ent.Picture, error)
}

func NewPictureRepository(db *ent.Client) PictureRepositoryInterface {
	return &pictureRepository{
		db: db,
	}
}

func (r *pictureRepository) Save(picture *ent.Picture) error {
	_, err := r.db.Picture.Create().
		SetAlbumID(picture.Edges.Album.ID).
		SetBodyPart(picture.BodyPart).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// GetAllByAlbumID는 각 album 안에 있는 모든 사진의 데이터를 조회
func (r *pictureRepository) GetAllByAlbumID(albumID int) ([]*ent.Picture, error) {
	pictures, err := r.db.Picture.Query().
		Where(picture.HasAlbumWith(album.ID(albumID))).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return pictures, nil
}

func (r *pictureRepository) Get(pictureID int) (*ent.Picture, error) {
	picture, err := r.db.Picture.Query().
		Where(picture.ID(pictureID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return picture, nil
}

// FindByAlbumIDAndBodyParts은 albumID 와 특정 신체 부위에 해당하는 사진들을 조회
func (r *pictureRepository) FindByAlbumIDAndBodyPart(albumID int, bodyPart string) ([]*ent.Picture, error) {
	pictures, err := r.db.Picture.Query().
		Where(picture.And(picture.HasAlbumWith(album.ID(albumID)), picture.BodyPart(bodyPart))).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return pictures, nil
}
