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
		field.String("key"),
		// taken_at은 사진을 찍은 날짜
		field.Time("taken_at"),
		// uploaded_at은 사진을 서버로 업로드한 날짜
		field.Time("uploaded_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Picture.
func (Picture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("album", Album.Type).
			Ref("picture").
			Unique(),
		edge.From("user", User.Type).
			Ref("picture").
			Required().
			Unique(),
	}
}
