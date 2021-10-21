package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Picture holds the schema definition for the Picture entity.
type Picture struct {
	ent.Schema
}

// Fields of the Picture.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.String("body_part"),
		// TODO: 이 부분은 location보다는 key가 좋을 것 같은데 어떨까요
		field.String("location"),
		field.Int("album_id"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Picture.
func (Picture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("album", Album.Type).
			Ref("picture").
			Unique().
			Required().
			Field("album_id"),
		edge.From("user", User.Type).
			Ref("picture").
			Unique(),
	}
}
