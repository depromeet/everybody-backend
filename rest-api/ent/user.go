// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProfileImage holds the value of the "profile_image" field.
	ProfileImage string `json:"profile_image,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Motto holds the value of the "motto" field.
	Motto string `json:"motto,omitempty"`
	// Height holds the value of the "height" field.
	Height *int `json:"height,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight *int `json:"weight,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind user.Kind `json:"kind,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Devices holds the value of the devices edge.
	Devices []*Device `json:"devices,omitempty"`
	// NotificationConfig holds the value of the notification_config edge.
	NotificationConfig []*NotificationConfig `json:"notification_config,omitempty"`
	// Album holds the value of the album edge.
	Album []*Album `json:"album,omitempty"`
	// Picture holds the value of the picture edge.
	Picture []*Picture `json:"picture,omitempty"`
	// Video holds the value of the video edge.
	Video []*Video `json:"video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DevicesOrErr() ([]*Device, error) {
	if e.loadedTypes[0] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// NotificationConfigOrErr returns the NotificationConfig value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationConfigOrErr() ([]*NotificationConfig, error) {
	if e.loadedTypes[1] {
		return e.NotificationConfig, nil
	}
	return nil, &NotLoadedError{edge: "notification_config"}
}

// AlbumOrErr returns the Album value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AlbumOrErr() ([]*Album, error) {
	if e.loadedTypes[2] {
		return e.Album, nil
	}
	return nil, &NotLoadedError{edge: "album"}
}

// PictureOrErr returns the Picture value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PictureOrErr() ([]*Picture, error) {
	if e.loadedTypes[3] {
		return e.Picture, nil
	}
	return nil, &NotLoadedError{edge: "picture"}
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VideoOrErr() ([]*Video, error) {
	if e.loadedTypes[4] {
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldHeight, user.FieldWeight:
			values[i] = new(sql.NullInt64)
		case user.FieldProfileImage, user.FieldNickname, user.FieldMotto, user.FieldKind:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldProfileImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_image", values[i])
			} else if value.Valid {
				u.ProfileImage = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldMotto:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field motto", values[i])
			} else if value.Valid {
				u.Motto = value.String
			}
		case user.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				u.Height = new(int)
				*u.Height = int(value.Int64)
			}
		case user.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				u.Weight = new(int)
				*u.Weight = int(value.Int64)
			}
		case user.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				u.Kind = user.Kind(value.String)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryDevices queries the "devices" edge of the User entity.
func (u *User) QueryDevices() *DeviceQuery {
	return (&UserClient{config: u.config}).QueryDevices(u)
}

// QueryNotificationConfig queries the "notification_config" edge of the User entity.
func (u *User) QueryNotificationConfig() *NotificationConfigQuery {
	return (&UserClient{config: u.config}).QueryNotificationConfig(u)
}

// QueryAlbum queries the "album" edge of the User entity.
func (u *User) QueryAlbum() *AlbumQuery {
	return (&UserClient{config: u.config}).QueryAlbum(u)
}

// QueryPicture queries the "picture" edge of the User entity.
func (u *User) QueryPicture() *PictureQuery {
	return (&UserClient{config: u.config}).QueryPicture(u)
}

// QueryVideo queries the "video" edge of the User entity.
func (u *User) QueryVideo() *VideoQuery {
	return (&UserClient{config: u.config}).QueryVideo(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", profile_image=")
	builder.WriteString(u.ProfileImage)
	builder.WriteString(", nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", motto=")
	builder.WriteString(u.Motto)
	if v := u.Height; v != nil {
		builder.WriteString(", height=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := u.Weight; v != nil {
		builder.WriteString(", weight=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", kind=")
	builder.WriteString(fmt.Sprintf("%v", u.Kind))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
