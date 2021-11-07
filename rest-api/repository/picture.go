package repository

import (
	"context"

	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/picture"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/pkg/errors"
)

type pictureRepository struct {
	db *ent.Client
}

type PictureRepositoryInterface interface {
	Save(picture *ent.Picture) (*ent.Picture, error)
	GetAllByUserID(userID int) ([]*ent.Picture, error)
	GetAllByAlbumID(albumID int) ([]*ent.Picture, error)
	Get(pictureID int) (*ent.Picture, error)
	FindByAlbumIDAndBodyPart(albumID int, bodyPart string) ([]*ent.Picture, error)
}

func NewPictureRepository(db *ent.Client) PictureRepositoryInterface {
	return &pictureRepository{
		db: db,
	}
}

func (r *pictureRepository) Save(picture *ent.Picture) (*ent.Picture, error) {
	p, err := r.db.Picture.Create().
		SetUser(picture.Edges.User).
		SetAlbum(picture.Edges.Album).
		SetBodyPart(picture.BodyPart).
		SetKey(picture.Key).
		SetTakenAt(picture.TakenAt).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := p.QueryUser().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	album, err := p.QueryAlbum().First(context.Background())
	if err != nil {
		return nil, err
	}
	p.Edges.User = user
	p.Edges.Album = album

	return p, nil
}

func (r *pictureRepository) GetAllByUserID(userID int) ([]*ent.Picture, error) {
	pictures, err := r.db.Picture.Query().
		Where(picture.HasUserWith(user.ID(userID))).
		WithUser().
		WithAlbum().
		Order(ent.Asc("taken_at")).
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pictures, nil
}

// GetAllByAlbumID는 각 album 안에 있는 모든 사진의 데이터를 조회
func (r *pictureRepository) GetAllByAlbumID(albumID int) ([]*ent.Picture, error) {
	pictures, err := r.db.Picture.Query().
		Where(picture.HasAlbumWith(album.ID(albumID))).
		WithUser().
		WithAlbum().
		Order(ent.Asc("taken_at")).
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pictures, nil
}

func (r *pictureRepository) Get(pictureID int) (*ent.Picture, error) {
	p, err := r.db.Picture.Query().
		Where(picture.ID(pictureID)).
		WithUser().
		WithAlbum().
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return p, nil
}

// FindByAlbumIDAndBodyParts은 albumID 와 특정 신체 부위에 해당하는 사진들을 조회
func (r *pictureRepository) FindByAlbumIDAndBodyPart(albumID int, bodyPart string) ([]*ent.Picture, error) {
	pictures, err := r.db.Picture.Query().
		Where(picture.And(picture.HasAlbumWith(album.ID(albumID)), picture.BodyPart(bodyPart))).
		WithUser().
		WithAlbum().
		Order(ent.Asc("taken_at")).
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pictures, nil
}
