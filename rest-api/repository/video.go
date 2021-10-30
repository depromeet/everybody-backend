package repository

import (
	"context"

	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/depromeet/everybody-backend/rest-api/ent/video"
	"github.com/pkg/errors"
)

type videoRepository struct {
	db *ent.Client
}

type VideoRepositoryInterface interface {
	Save(video *ent.Video) (*ent.Video, error)
	Get(videoID int) (*ent.Video, error)
	GetAllByUserID(userID int) ([]*ent.Video, error)
}

func NewVideoRepository(db *ent.Client) VideoRepositoryInterface {
	return &videoRepository{
		db: db,
	}
}

func (r *videoRepository) Save(video *ent.Video) (*ent.Video, error) {
	v, err := r.db.Video.Create().
		SetUser(video.Edges.User).
		SetKey(video.Key).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// album, err := v.QueryAlbum().First(context.Background())
	// if err != nil {
	// 	// rollback 해줘야 하나...
	// 	return nil, err
	// }

	// v.Edges.Album = album
	return v, nil
}

func (r *videoRepository) Get(videoID int) (*ent.Video, error) {
	v, err := r.db.Video.Query().
		Where(video.ID(videoID)).
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return v, nil
}

func (r *videoRepository) GetAllByUserID(userID int) ([]*ent.Video, error) {
	videos, err := r.db.Video.Query().
		Where(video.HasUserWith(user.ID(userID))).
		WithUser().
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return videos, nil
}
