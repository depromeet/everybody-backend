package repository

import (
	"context"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"time"
)

type NotificationRepository interface {
	CreateNotificationConfig(config *ent.NotificationConfig) (*ent.NotificationConfig, error)
	FindById(id int) (*ent.NotificationConfig, error)
	UpdateInterval(id, interval int) (*ent.NotificationConfig, error)
	UpdateLastNotifiedAt(id int, lastNotifiedAt time.Time) (*ent.NotificationConfig, error)
	UpdateIsActivated(id int, isActivated bool) (*ent.NotificationConfig, error)
}

func NewNotificationRepository(client *ent.Client) NotificationRepository {
	return &notificationRepository{
		db: client,
	}
}

type notificationRepository struct {
	db *ent.Client
}

func (repo *notificationRepository) CreateNotificationConfig(config *ent.NotificationConfig) (*ent.NotificationConfig, error) {
	result, err := repo.db.NotificationConfig.Create().
		SetUser(config.Edges.User).
		SetInterval(config.Interval).
		SetIsActivated(config.IsActivated).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *notificationRepository) FindById(id int) (*ent.NotificationConfig, error) {
	u, err := repo.db.NotificationConfig.Query().
		Where(notificationconfig.ID(id)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (repo *notificationRepository) UpdateInterval(id, interval int) (*ent.NotificationConfig, error) {
	return repo.db.NotificationConfig.UpdateOneID(id).SetInterval(interval).Save(context.Background())

}

func (repo *notificationRepository) UpdateLastNotifiedAt(id int, lastNotifiedAt time.Time) (*ent.NotificationConfig, error) {
	return repo.db.NotificationConfig.UpdateOneID(id).SetLastNotifiedAt(lastNotifiedAt).Save(context.Background())
}

func (repo *notificationRepository) UpdateIsActivated(id int, isActivated bool) (*ent.NotificationConfig, error) {
	return repo.db.NotificationConfig.UpdateOneID(id).SetIsActivated(isActivated).Save(context.Background())
}
