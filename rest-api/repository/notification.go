package repository

import (
	"context"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/pkg/errors"
)

type NotificationRepository interface {
	CreateNotificationConfig(config *ent.NotificationConfig) (*ent.NotificationConfig, error)
	FindById(id int) (*ent.NotificationConfig, error)
	FindByUser(user int) (*ent.NotificationConfig, error)
	Update(id int, config *ent.NotificationConfig) (*ent.NotificationConfig, error)
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
		SetMonday(config.Monday).
		SetTuesday(config.Tuesday).
		SetWednesday(config.Wednesday).
		SetThursday(config.Thursday).
		SetFriday(config.Friday).
		SetSaturday(config.Saturday).
		SetSunday(config.Sunday).
		SetPreferredTimeHour(config.PreferredTimeHour).
		SetPreferredTimeMinute(config.PreferredTimeMinute).
		SetNillableLastNotifiedAt(config.LastNotifiedAt).
		SetIsActivated(config.IsActivated).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (repo *notificationRepository) FindById(id int) (*ent.NotificationConfig, error) {
	u, err := repo.db.NotificationConfig.Query().
		Where(notificationconfig.ID(id)).
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func (repo *notificationRepository) FindByUser(userID int) (*ent.NotificationConfig, error) {
	u, err := repo.db.NotificationConfig.Query().
		Where(notificationconfig.HasUserWith(user.ID(userID))).
		WithUser().
		First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}

func (repo *notificationRepository) Update(id int, config *ent.NotificationConfig) (*ent.NotificationConfig, error) {
	return repo.db.NotificationConfig.UpdateOneID(id).
		SetUser(config.Edges.User).
		SetMonday(config.Monday).
		SetTuesday(config.Tuesday).
		SetWednesday(config.Wednesday).
		SetThursday(config.Thursday).
		SetFriday(config.Friday).
		SetSaturday(config.Saturday).
		SetSunday(config.Sunday).
		SetPreferredTimeHour(config.PreferredTimeHour).
		SetPreferredTimeMinute(config.PreferredTimeMinute).
		SetNillableLastNotifiedAt(config.LastNotifiedAt).
		SetIsActivated(config.IsActivated).
		Save(context.Background())
}
