package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

var kinds = []string{"SIMPLE", "KAKAO", "APPLE", "NAVER", "GOOGLE"}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("nickname"),
		field.Int("height").Optional().Nillable(),
		field.Int("weight").Optional().Nillable(),
		field.Enum("kind").Values(kinds...),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("notification_config", NotificationConfig.Type),
		edge.To("album", Album.Type),
		edge.To("picture", Picture.Type),
		edge.To("video", Video.Type),
	}
}
