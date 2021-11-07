// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/picture"
	"github.com/depromeet/everybody-backend/rest-api/ent/schema"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/depromeet/everybody-backend/rest-api/ent/video"
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
	// notificationconfigDescMonday is the schema descriptor for monday field.
	notificationconfigDescMonday := notificationconfigFields[1].Descriptor()
	// notificationconfig.DefaultMonday holds the default value on creation for the monday field.
	notificationconfig.DefaultMonday = notificationconfigDescMonday.Default.(bool)
	// notificationconfigDescTuesday is the schema descriptor for tuesday field.
	notificationconfigDescTuesday := notificationconfigFields[2].Descriptor()
	// notificationconfig.DefaultTuesday holds the default value on creation for the tuesday field.
	notificationconfig.DefaultTuesday = notificationconfigDescTuesday.Default.(bool)
	// notificationconfigDescWednesday is the schema descriptor for wednesday field.
	notificationconfigDescWednesday := notificationconfigFields[3].Descriptor()
	// notificationconfig.DefaultWednesday holds the default value on creation for the wednesday field.
	notificationconfig.DefaultWednesday = notificationconfigDescWednesday.Default.(bool)
	// notificationconfigDescThursday is the schema descriptor for thursday field.
	notificationconfigDescThursday := notificationconfigFields[4].Descriptor()
	// notificationconfig.DefaultThursday holds the default value on creation for the thursday field.
	notificationconfig.DefaultThursday = notificationconfigDescThursday.Default.(bool)
	// notificationconfigDescFriday is the schema descriptor for friday field.
	notificationconfigDescFriday := notificationconfigFields[5].Descriptor()
	// notificationconfig.DefaultFriday holds the default value on creation for the friday field.
	notificationconfig.DefaultFriday = notificationconfigDescFriday.Default.(bool)
	// notificationconfigDescSaturday is the schema descriptor for saturday field.
	notificationconfigDescSaturday := notificationconfigFields[6].Descriptor()
	// notificationconfig.DefaultSaturday holds the default value on creation for the saturday field.
	notificationconfig.DefaultSaturday = notificationconfigDescSaturday.Default.(bool)
	// notificationconfigDescSunday is the schema descriptor for sunday field.
	notificationconfigDescSunday := notificationconfigFields[7].Descriptor()
	// notificationconfig.DefaultSunday holds the default value on creation for the sunday field.
	notificationconfig.DefaultSunday = notificationconfigDescSunday.Default.(bool)
	// notificationconfigDescIsActivated is the schema descriptor for is_activated field.
	notificationconfigDescIsActivated := notificationconfigFields[11].Descriptor()
	// notificationconfig.DefaultIsActivated holds the default value on creation for the is_activated field.
	notificationconfig.DefaultIsActivated = notificationconfigDescIsActivated.Default.(bool)
	pictureFields := schema.Picture{}.Fields()
	_ = pictureFields
	// pictureDescCreatedAt is the schema descriptor for created_at field.
	pictureDescCreatedAt := pictureFields[3].Descriptor()
	// picture.DefaultCreatedAt holds the default value on creation for the created_at field.
	picture.DefaultCreatedAt = pictureDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescMotto is the schema descriptor for motto field.
	userDescMotto := userFields[3].Descriptor()
	// user.DefaultMotto holds the default value on creation for the motto field.
	user.DefaultMotto = userDescMotto.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescCreatedAt is the schema descriptor for created_at field.
	videoDescCreatedAt := videoFields[0].Descriptor()
	// video.DefaultCreatedAt holds the default value on creation for the created_at field.
	video.DefaultCreatedAt = videoDescCreatedAt.Default.(func() time.Time)
}
