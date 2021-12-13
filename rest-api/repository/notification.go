package repository

import (
	"context"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/pkg/errors"
	"time"
)

type NotificationRepository interface {
	CreateNotificationConfig(config *ent.NotificationConfig) (*ent.NotificationConfig, error)
	FindAll() ([]*ent.NotificationConfig, error)
	FindById(id int) (*ent.NotificationConfig, error)
	FindByUser(user int) (*ent.NotificationConfig, error)
	Update(id int, config *ent.NotificationConfig) (*ent.NotificationConfig, error)
	UpdateLastNotifiedAt(id int, lastNotifiedAt time.Time) (*ent.NotificationConfig, error)
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

// TODO(umi0410): 유저가 엄~청 많아지면 Pagination이 필요할 수도....
func (repo *notificationRepository) FindAll() ([]*ent.NotificationConfig, error) {
	users, err := repo.db.NotificationConfig.Query().
		WithUser(func(query *ent.UserQuery) {
			// device 정보도 같이 fetch 하자.
			query.WithDevices()
		}).
		All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return users, nil
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
	builder := repo.db.NotificationConfig.UpdateOneID(id).
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
		SetIsActivated(config.IsActivated)
	if config.LastNotifiedAt == nil {
		builder = builder.ClearLastNotifiedAt()
	} else {
		builder = builder.SetLastNotifiedAt(*config.LastNotifiedAt)
	}
	cfg, err := builder.Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cfg, nil
}

func (repo *notificationRepository) UpdateLastNotifiedAt(id int, lastNotifiedAt time.Time) (*ent.NotificationConfig, error) {
	return repo.db.NotificationConfig.UpdateOneID(id).
		SetLastNotifiedAt(lastNotifiedAt).
		Save(context.Background())
}
