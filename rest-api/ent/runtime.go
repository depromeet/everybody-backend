// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/picture"
	"github.com/depromeet/everybody-backend/rest-api/ent/schema"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	albumFields := schema.Album{}.Fields()
	_ = albumFields
	// albumDescCreatedAt is the schema descriptor for created_at field.
	albumDescCreatedAt := albumFields[1].Descriptor()
	// album.DefaultCreatedAt holds the default value on creation for the created_at field.
	album.DefaultCreatedAt = albumDescCreatedAt.Default.(func() time.Time)
	notificationconfigFields := schema.NotificationConfig{}.Fields()
	_ = notificationconfigFields
	// notificationconfigDescIsActivated is the schema descriptor for is_activated field.
	notificationconfigDescIsActivated := notificationconfigFields[3].Descriptor()
	// notificationconfig.DefaultIsActivated holds the default value on creation for the is_activated field.
	notificationconfig.DefaultIsActivated = notificationconfigDescIsActivated.Default.(bool)
	pictureFields := schema.Picture{}.Fields()
	_ = pictureFields
	// pictureDescCreatedAt is the schema descriptor for created_at field.
	pictureDescCreatedAt := pictureFields[1].Descriptor()
	// picture.DefaultCreatedAt holds the default value on creation for the created_at field.
	picture.DefaultCreatedAt = pictureDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
