// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/depromeet/everybody-backend/rest-api/ent/migrate"

	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/device"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/picture"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/depromeet/everybody-backend/rest-api/ent/video"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Album is the client for interacting with the Album builders.
	Album *AlbumClient
	// Device is the client for interacting with the Device builders.
	Device *DeviceClient
	// NotificationConfig is the client for interacting with the NotificationConfig builders.
	NotificationConfig *NotificationConfigClient
	// Picture is the client for interacting with the Picture builders.
	Picture *PictureClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Video is the client for interacting with the Video builders.
	Video *VideoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Album = NewAlbumClient(c.config)
	c.Device = NewDeviceClient(c.config)
	c.NotificationConfig = NewNotificationConfigClient(c.config)
	c.Picture = NewPictureClient(c.config)
	c.User = NewUserClient(c.config)
	c.Video = NewVideoClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Album:              NewAlbumClient(cfg),
		Device:             NewDeviceClient(cfg),
		NotificationConfig: NewNotificationConfigClient(cfg),
		Picture:            NewPictureClient(cfg),
		User:               NewUserClient(cfg),
		Video:              NewVideoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:             cfg,
		Album:              NewAlbumClient(cfg),
		Device:             NewDeviceClient(cfg),
		NotificationConfig: NewNotificationConfigClient(cfg),
		Picture:            NewPictureClient(cfg),
		User:               NewUserClient(cfg),
		Video:              NewVideoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Album.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Album.Use(hooks...)
	c.Device.Use(hooks...)
	c.NotificationConfig.Use(hooks...)
	c.Picture.Use(hooks...)
	c.User.Use(hooks...)
	c.Video.Use(hooks...)
}

// AlbumClient is a client for the Album schema.
type AlbumClient struct {
	config
}

// NewAlbumClient returns a client for the Album from the given config.
func NewAlbumClient(c config) *AlbumClient {
	return &AlbumClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `album.Hooks(f(g(h())))`.
func (c *AlbumClient) Use(hooks ...Hook) {
	c.hooks.Album = append(c.hooks.Album, hooks...)
}

// Create returns a create builder for Album.
func (c *AlbumClient) Create() *AlbumCreate {
	mutation := newAlbumMutation(c.config, OpCreate)
	return &AlbumCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Album entities.
func (c *AlbumClient) CreateBulk(builders ...*AlbumCreate) *AlbumCreateBulk {
	return &AlbumCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Album.
func (c *AlbumClient) Update() *AlbumUpdate {
	mutation := newAlbumMutation(c.config, OpUpdate)
	return &AlbumUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlbumClient) UpdateOne(a *Album) *AlbumUpdateOne {
	mutation := newAlbumMutation(c.config, OpUpdateOne, withAlbum(a))
	return &AlbumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlbumClient) UpdateOneID(id int) *AlbumUpdateOne {
	mutation := newAlbumMutation(c.config, OpUpdateOne, withAlbumID(id))
	return &AlbumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Album.
func (c *AlbumClient) Delete() *AlbumDelete {
	mutation := newAlbumMutation(c.config, OpDelete)
	return &AlbumDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AlbumClient) DeleteOne(a *Album) *AlbumDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AlbumClient) DeleteOneID(id int) *AlbumDeleteOne {
	builder := c.Delete().Where(album.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlbumDeleteOne{builder}
}

// Query returns a query builder for Album.
func (c *AlbumClient) Query() *AlbumQuery {
	return &AlbumQuery{
		config: c.config,
	}
}

// Get returns a Album entity by its id.
func (c *AlbumClient) Get(ctx context.Context, id int) (*Album, error) {
	return c.Query().Where(album.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlbumClient) GetX(ctx context.Context, id int) *Album {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Album.
func (c *AlbumClient) QueryUser(a *Album) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, album.UserTable, album.UserColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPicture queries the picture edge of a Album.
func (c *AlbumClient) QueryPicture(a *Album) *PictureQuery {
	query := &PictureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(picture.Table, picture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, album.PictureTable, album.PictureColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVideo queries the video edge of a Album.
func (c *AlbumClient) QueryVideo(a *Album) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, album.VideoTable, album.VideoColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AlbumClient) Hooks() []Hook {
	return c.hooks.Album
}

// DeviceClient is a client for the Device schema.
type DeviceClient struct {
	config
}

// NewDeviceClient returns a client for the Device from the given config.
func NewDeviceClient(c config) *DeviceClient {
	return &DeviceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `device.Hooks(f(g(h())))`.
func (c *DeviceClient) Use(hooks ...Hook) {
	c.hooks.Device = append(c.hooks.Device, hooks...)
}

// Create returns a create builder for Device.
func (c *DeviceClient) Create() *DeviceCreate {
	mutation := newDeviceMutation(c.config, OpCreate)
	return &DeviceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Device entities.
func (c *DeviceClient) CreateBulk(builders ...*DeviceCreate) *DeviceCreateBulk {
	return &DeviceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Device.
func (c *DeviceClient) Update() *DeviceUpdate {
	mutation := newDeviceMutation(c.config, OpUpdate)
	return &DeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeviceClient) UpdateOne(d *Device) *DeviceUpdateOne {
	mutation := newDeviceMutation(c.config, OpUpdateOne, withDevice(d))
	return &DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeviceClient) UpdateOneID(id int) *DeviceUpdateOne {
	mutation := newDeviceMutation(c.config, OpUpdateOne, withDeviceID(id))
	return &DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Device.
func (c *DeviceClient) Delete() *DeviceDelete {
	mutation := newDeviceMutation(c.config, OpDelete)
	return &DeviceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DeviceClient) DeleteOne(d *Device) *DeviceDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DeviceClient) DeleteOneID(id int) *DeviceDeleteOne {
	builder := c.Delete().Where(device.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeviceDeleteOne{builder}
}

// Query returns a query builder for Device.
func (c *DeviceClient) Query() *DeviceQuery {
	return &DeviceQuery{
		config: c.config,
	}
}

// Get returns a Device entity by its id.
func (c *DeviceClient) Get(ctx context.Context, id int) (*Device, error) {
	return c.Query().Where(device.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeviceClient) GetX(ctx context.Context, id int) *Device {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Device.
func (c *DeviceClient) QueryUser(d *Device) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, device.UserTable, device.UserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeviceClient) Hooks() []Hook {
	return c.hooks.Device
}

// NotificationConfigClient is a client for the NotificationConfig schema.
type NotificationConfigClient struct {
	config
}

// NewNotificationConfigClient returns a client for the NotificationConfig from the given config.
func NewNotificationConfigClient(c config) *NotificationConfigClient {
	return &NotificationConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notificationconfig.Hooks(f(g(h())))`.
func (c *NotificationConfigClient) Use(hooks ...Hook) {
	c.hooks.NotificationConfig = append(c.hooks.NotificationConfig, hooks...)
}

// Create returns a create builder for NotificationConfig.
func (c *NotificationConfigClient) Create() *NotificationConfigCreate {
	mutation := newNotificationConfigMutation(c.config, OpCreate)
	return &NotificationConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NotificationConfig entities.
func (c *NotificationConfigClient) CreateBulk(builders ...*NotificationConfigCreate) *NotificationConfigCreateBulk {
	return &NotificationConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NotificationConfig.
func (c *NotificationConfigClient) Update() *NotificationConfigUpdate {
	mutation := newNotificationConfigMutation(c.config, OpUpdate)
	return &NotificationConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotificationConfigClient) UpdateOne(nc *NotificationConfig) *NotificationConfigUpdateOne {
	mutation := newNotificationConfigMutation(c.config, OpUpdateOne, withNotificationConfig(nc))
	return &NotificationConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotificationConfigClient) UpdateOneID(id int) *NotificationConfigUpdateOne {
	mutation := newNotificationConfigMutation(c.config, OpUpdateOne, withNotificationConfigID(id))
	return &NotificationConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NotificationConfig.
func (c *NotificationConfigClient) Delete() *NotificationConfigDelete {
	mutation := newNotificationConfigMutation(c.config, OpDelete)
	return &NotificationConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NotificationConfigClient) DeleteOne(nc *NotificationConfig) *NotificationConfigDeleteOne {
	return c.DeleteOneID(nc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NotificationConfigClient) DeleteOneID(id int) *NotificationConfigDeleteOne {
	builder := c.Delete().Where(notificationconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotificationConfigDeleteOne{builder}
}

// Query returns a query builder for NotificationConfig.
func (c *NotificationConfigClient) Query() *NotificationConfigQuery {
	return &NotificationConfigQuery{
		config: c.config,
	}
}

// Get returns a NotificationConfig entity by its id.
func (c *NotificationConfigClient) Get(ctx context.Context, id int) (*NotificationConfig, error) {
	return c.Query().Where(notificationconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotificationConfigClient) GetX(ctx context.Context, id int) *NotificationConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a NotificationConfig.
func (c *NotificationConfigClient) QueryUser(nc *NotificationConfig) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := nc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationconfig.Table, notificationconfig.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notificationconfig.UserTable, notificationconfig.UserColumn),
		)
		fromV = sqlgraph.Neighbors(nc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NotificationConfigClient) Hooks() []Hook {
	return c.hooks.NotificationConfig
}

// PictureClient is a client for the Picture schema.
type PictureClient struct {
	config
}

// NewPictureClient returns a client for the Picture from the given config.
func NewPictureClient(c config) *PictureClient {
	return &PictureClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `picture.Hooks(f(g(h())))`.
func (c *PictureClient) Use(hooks ...Hook) {
	c.hooks.Picture = append(c.hooks.Picture, hooks...)
}

// Create returns a create builder for Picture.
func (c *PictureClient) Create() *PictureCreate {
	mutation := newPictureMutation(c.config, OpCreate)
	return &PictureCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Picture entities.
func (c *PictureClient) CreateBulk(builders ...*PictureCreate) *PictureCreateBulk {
	return &PictureCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Picture.
func (c *PictureClient) Update() *PictureUpdate {
	mutation := newPictureMutation(c.config, OpUpdate)
	return &PictureUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PictureClient) UpdateOne(pi *Picture) *PictureUpdateOne {
	mutation := newPictureMutation(c.config, OpUpdateOne, withPicture(pi))
	return &PictureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PictureClient) UpdateOneID(id int) *PictureUpdateOne {
	mutation := newPictureMutation(c.config, OpUpdateOne, withPictureID(id))
	return &PictureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Picture.
func (c *PictureClient) Delete() *PictureDelete {
	mutation := newPictureMutation(c.config, OpDelete)
	return &PictureDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PictureClient) DeleteOne(pi *Picture) *PictureDeleteOne {
	return c.DeleteOneID(pi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PictureClient) DeleteOneID(id int) *PictureDeleteOne {
	builder := c.Delete().Where(picture.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PictureDeleteOne{builder}
}

// Query returns a query builder for Picture.
func (c *PictureClient) Query() *PictureQuery {
	return &PictureQuery{
		config: c.config,
	}
}

// Get returns a Picture entity by its id.
func (c *PictureClient) Get(ctx context.Context, id int) (*Picture, error) {
	return c.Query().Where(picture.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PictureClient) GetX(ctx context.Context, id int) *Picture {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAlbum queries the album edge of a Picture.
func (c *PictureClient) QueryAlbum(pi *Picture) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(picture.Table, picture.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, picture.AlbumTable, picture.AlbumColumn),
		)
		fromV = sqlgraph.Neighbors(pi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Picture.
func (c *PictureClient) QueryUser(pi *Picture) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(picture.Table, picture.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, picture.UserTable, picture.UserColumn),
		)
		fromV = sqlgraph.Neighbors(pi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PictureClient) Hooks() []Hook {
	return c.hooks.Picture
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDevices queries the devices edge of a User.
func (c *UserClient) QueryDevices(u *User) *DeviceQuery {
	query := &DeviceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DevicesTable, user.DevicesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNotificationConfig queries the notification_config edge of a User.
func (c *UserClient) QueryNotificationConfig(u *User) *NotificationConfigQuery {
	query := &NotificationConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(notificationconfig.Table, notificationconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.NotificationConfigTable, user.NotificationConfigColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAlbum queries the album edge of a User.
func (c *UserClient) QueryAlbum(u *User) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AlbumTable, user.AlbumColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPicture queries the picture edge of a User.
func (c *UserClient) QueryPicture(u *User) *PictureQuery {
	query := &PictureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(picture.Table, picture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PictureTable, user.PictureColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVideo queries the video edge of a User.
func (c *UserClient) QueryVideo(u *User) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.VideoTable, user.VideoColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VideoClient is a client for the Video schema.
type VideoClient struct {
	config
}

// NewVideoClient returns a client for the Video from the given config.
func NewVideoClient(c config) *VideoClient {
	return &VideoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `video.Hooks(f(g(h())))`.
func (c *VideoClient) Use(hooks ...Hook) {
	c.hooks.Video = append(c.hooks.Video, hooks...)
}

// Create returns a create builder for Video.
func (c *VideoClient) Create() *VideoCreate {
	mutation := newVideoMutation(c.config, OpCreate)
	return &VideoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Video entities.
func (c *VideoClient) CreateBulk(builders ...*VideoCreate) *VideoCreateBulk {
	return &VideoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Video.
func (c *VideoClient) Update() *VideoUpdate {
	mutation := newVideoMutation(c.config, OpUpdate)
	return &VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VideoClient) UpdateOne(v *Video) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideo(v))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VideoClient) UpdateOneID(id int) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideoID(id))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Video.
func (c *VideoClient) Delete() *VideoDelete {
	mutation := newVideoMutation(c.config, OpDelete)
	return &VideoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VideoClient) DeleteOne(v *Video) *VideoDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VideoClient) DeleteOneID(id int) *VideoDeleteOne {
	builder := c.Delete().Where(video.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VideoDeleteOne{builder}
}

// Query returns a query builder for Video.
func (c *VideoClient) Query() *VideoQuery {
	return &VideoQuery{
		config: c.config,
	}
}

// Get returns a Video entity by its id.
func (c *VideoClient) Get(ctx context.Context, id int) (*Video, error) {
	return c.Query().Where(video.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VideoClient) GetX(ctx context.Context, id int) *Video {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Video.
func (c *VideoClient) QueryUser(v *Video) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, video.UserTable, video.UserColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAlbum queries the album edge of a Video.
func (c *VideoClient) QueryAlbum(v *Video) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, video.AlbumTable, video.AlbumColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VideoClient) Hooks() []Hook {
	return c.hooks.Video
}
