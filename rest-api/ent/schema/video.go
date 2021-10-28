package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.String("key"),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("video").
			Required().
			Unique(),
		edge.From("album", Album.Type).
			Ref("video").
			Unique(),
	}

}
