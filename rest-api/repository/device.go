package repository

import (
	"context"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/device"
	"github.com/pkg/errors"
)

type DeviceRepository interface {
	CreateDevice(device *ent.Device) (*ent.Device, error)
	FindById(id int) (*ent.Device, error)
	FindByDeviceToken(deviceToken string) (*ent.Device, error)
	Update(id int, device *ent.Device) error
}

func NewDeviceRepository(client *ent.Client) DeviceRepository {
	return &deviceRepository{
		db: client,
	}
}

type deviceRepository struct {
	db *ent.Client
}

func (repo *deviceRepository) CreateDevice(device *ent.Device) (*ent.Device, error) {
	result, err := repo.db.Device.Create().
		SetUser(device.Edges.User).
		SetDeviceToken(device.DeviceToken).
		SetPushToken(device.PushToken).
		SetDeviceOs(device.DeviceOs).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (repo *deviceRepository) FindById(id int) (*ent.Device, error) {
	d, err := repo.db.Device.Query().
		Where(device.ID(id)).
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return d, nil
}

func (repo *deviceRepository) FindByDeviceToken(deviceToken string) (*ent.Device, error) {
	d, err := repo.db.Device.Query().
		Where(device.DeviceToken(deviceToken)).
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return d, nil
}

func (repo *deviceRepository) Update(id int, device *ent.Device) error {
	err := repo.db.Device.UpdateOneID(id).
		SetUser(device.Edges.User).
		SetPushToken(device.PushToken).
		Exec(context.Background())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
