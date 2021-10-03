package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("nickname"),
		field.Int("height"),
		field.Int("weight"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("device", Device.Type),
		edge.To("notification_config", NotificationConfig.Type),
	}
}
