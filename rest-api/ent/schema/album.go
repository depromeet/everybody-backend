package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("folder_name"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("album").
			Unique(),
	}
}
