// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldMotto holds the string denoting the motto field in the database.
	FieldMotto = "motto"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeNotificationConfig holds the string denoting the notification_config edge name in mutations.
	EdgeNotificationConfig = "notification_config"
	// EdgeAlbum holds the string denoting the album edge name in mutations.
	EdgeAlbum = "album"
	// EdgePicture holds the string denoting the picture edge name in mutations.
	EdgePicture = "picture"
	// Table holds the table name of the user in the database.
	Table = "users"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "devices"
	// DevicesInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DevicesInverseTable = "devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "user_devices"
	// NotificationConfigTable is the table that holds the notification_config relation/edge.
	NotificationConfigTable = "notification_configs"
	// NotificationConfigInverseTable is the table name for the NotificationConfig entity.
	// It exists in this package in order to avoid circular dependency with the "notificationconfig" package.
	NotificationConfigInverseTable = "notification_configs"
	// NotificationConfigColumn is the table column denoting the notification_config relation/edge.
	NotificationConfigColumn = "user_notification_config"
	// AlbumTable is the table that holds the album relation/edge.
	AlbumTable = "albums"
	// AlbumInverseTable is the table name for the Album entity.
	// It exists in this package in order to avoid circular dependency with the "album" package.
	AlbumInverseTable = "albums"
	// AlbumColumn is the table column denoting the album relation/edge.
	AlbumColumn = "user_album"
	// PictureTable is the table that holds the picture relation/edge.
	PictureTable = "pictures"
	// PictureInverseTable is the table name for the Picture entity.
	// It exists in this package in order to avoid circular dependency with the "picture" package.
	PictureInverseTable = "pictures"
	// PictureColumn is the table column denoting the picture relation/edge.
	PictureColumn = "user_picture"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldNickname,
	FieldMotto,
	FieldHeight,
	FieldWeight,
	FieldKind,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultMotto holds the default value on creation for the "motto" field.
	DefaultMotto string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindSIMPLE Kind = "SIMPLE"
	KindKAKAO  Kind = "KAKAO"
	KindAPPLE  Kind = "APPLE"
	KindNAVER  Kind = "NAVER"
	KindGOOGLE Kind = "GOOGLE"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindSIMPLE, KindKAKAO, KindAPPLE, KindNAVER, KindGOOGLE:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for kind field: %q", k)
	}
}
