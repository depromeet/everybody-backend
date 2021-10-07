// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

// Album is the model entity for the Album schema.
type Album struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FolderName holds the value of the "folder_name" field.
	FolderName string `json:"folder_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AlbumQuery when eager-loading is set.
	Edges      AlbumEdges `json:"edges"`
	user_album *string
}

// AlbumEdges holds the relations/edges for other nodes in the graph.
type AlbumEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Picture holds the value of the picture edge.
	Picture []*Picture `json:"picture,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AlbumEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PictureOrErr returns the Picture value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) PictureOrErr() ([]*Picture, error) {
	if e.loadedTypes[1] {
		return e.Picture, nil
	}
	return nil, &NotLoadedError{edge: "picture"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Album) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			values[i] = new(sql.NullInt64)
		case album.FieldFolderName:
			values[i] = new(sql.NullString)
		case album.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case album.ForeignKeys[0]: // user_album
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Album", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Album fields.
func (a *Album) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case album.FieldFolderName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field folder_name", values[i])
			} else if value.Valid {
				a.FolderName = value.String
			}
		case album.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case album.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_album", values[i])
			} else if value.Valid {
				a.user_album = new(string)
				*a.user_album = value.String
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Album entity.
func (a *Album) QueryUser() *UserQuery {
	return (&AlbumClient{config: a.config}).QueryUser(a)
}

// QueryPicture queries the "picture" edge of the Album entity.
func (a *Album) QueryPicture() *PictureQuery {
	return (&AlbumClient{config: a.config}).QueryPicture(a)
}

// Update returns a builder for updating this Album.
// Note that you need to call Album.Unwrap() before calling this method if this Album
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Album) Update() *AlbumUpdateOne {
	return (&AlbumClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Album entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Album) Unwrap() *Album {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Album is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Album) String() string {
	var builder strings.Builder
	builder.WriteString("Album(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", folder_name=")
	builder.WriteString(a.FolderName)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Albums is a parsable slice of Album.
type Albums []*Album

func (a Albums) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
