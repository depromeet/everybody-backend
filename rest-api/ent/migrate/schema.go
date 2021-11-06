// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AlbumsColumns holds the columns for the "albums" table.
	AlbumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_album", Type: field.TypeInt, Nullable: true},
	}
	// AlbumsTable holds the schema information for the "albums" table.
	AlbumsTable = &schema.Table{
		Name:       "albums",
		Columns:    AlbumsColumns,
		PrimaryKey: []*schema.Column{AlbumsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "albums_users_album",
				Columns:    []*schema.Column{AlbumsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "device_token", Type: field.TypeString},
		{Name: "push_token", Type: field.TypeString},
		{Name: "device_os", Type: field.TypeEnum, Enums: []string{"ANDROID", "IOS"}},
		{Name: "user_devices", Type: field.TypeInt, Nullable: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_users_devices",
				Columns:    []*schema.Column{DevicesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotificationConfigsColumns holds the columns for the "notification_configs" table.
	NotificationConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "monday", Type: field.TypeBool, Default: false},
		{Name: "tuesday", Type: field.TypeBool, Default: false},
		{Name: "wednesday", Type: field.TypeBool, Default: false},
		{Name: "thursday", Type: field.TypeBool, Default: false},
		{Name: "friday", Type: field.TypeBool, Default: false},
		{Name: "saturday", Type: field.TypeBool, Default: false},
		{Name: "sunday", Type: field.TypeBool, Default: false},
		{Name: "preferred_time_hour", Type: field.TypeInt, Nullable: true},
		{Name: "preferred_time_minute", Type: field.TypeInt, Nullable: true},
		{Name: "last_notified_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_activated", Type: field.TypeBool, Default: true},
		{Name: "user_notification_config", Type: field.TypeInt, Nullable: true},
	}
	// NotificationConfigsTable holds the schema information for the "notification_configs" table.
	NotificationConfigsTable = &schema.Table{
		Name:       "notification_configs",
		Columns:    NotificationConfigsColumns,
		PrimaryKey: []*schema.Column{NotificationConfigsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notification_configs_users_notification_config",
				Columns:    []*schema.Column{NotificationConfigsColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PicturesColumns holds the columns for the "pictures" table.
	PicturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "body_part", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "taken_at", Type: field.TypeTime},
		{Name: "uploaded_at", Type: field.TypeTime},
		{Name: "album_picture", Type: field.TypeInt, Nullable: true},
		{Name: "user_picture", Type: field.TypeInt, Nullable: true},
	}
	// PicturesTable holds the schema information for the "pictures" table.
	PicturesTable = &schema.Table{
		Name:       "pictures",
		Columns:    PicturesColumns,
		PrimaryKey: []*schema.Column{PicturesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pictures_albums_picture",
				Columns:    []*schema.Column{PicturesColumns[5]},
				RefColumns: []*schema.Column{AlbumsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pictures_users_picture",
				Columns:    []*schema.Column{PicturesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "motto", Type: field.TypeString, Default: "눈바디와 함께 꾸준히 운동할테야!"},
		{Name: "height", Type: field.TypeInt, Nullable: true},
		{Name: "weight", Type: field.TypeInt, Nullable: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"SIMPLE", "KAKAO", "APPLE", "NAVER", "GOOGLE"}},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VideosColumns holds the columns for the "videos" table.
	VideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString},
		{Name: "user_video", Type: field.TypeInt, Nullable: true},
	}
	// VideosTable holds the schema information for the "videos" table.
	VideosTable = &schema.Table{
		Name:       "videos",
		Columns:    VideosColumns,
		PrimaryKey: []*schema.Column{VideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "videos_users_video",
				Columns:    []*schema.Column{VideosColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlbumsTable,
		DevicesTable,
		NotificationConfigsTable,
		PicturesTable,
		UsersTable,
		VideosTable,
	}
)

func init() {
	AlbumsTable.ForeignKeys[0].RefTable = UsersTable
	DevicesTable.ForeignKeys[0].RefTable = UsersTable
	NotificationConfigsTable.ForeignKeys[0].RefTable = UsersTable
	PicturesTable.ForeignKeys[0].RefTable = AlbumsTable
	PicturesTable.ForeignKeys[1].RefTable = UsersTable
	VideosTable.ForeignKeys[0].RefTable = UsersTable
}
