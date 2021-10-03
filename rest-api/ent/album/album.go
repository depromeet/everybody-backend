// Code generated by entc, DO NOT EDIT.

package album

import (
	"time"
)

const (
	// Label holds the string label denoting the album type in the database.
	Label = "album"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFolderName holds the string denoting the folder_name field in the database.
	FieldFolderName = "folder_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the album in the database.
	Table = "albums"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "albums"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_album"
)

// Columns holds all SQL columns for album fields.
var Columns = []string{
	FieldID,
	FieldFolderName,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "albums"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_album",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
