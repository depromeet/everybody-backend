package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NotificationConfig holds the schema definition for the NotificationConfig entity.
type NotificationConfig struct {
	ent.Schema
}

// Fields of the NotificationConfig.
func (NotificationConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("interval").Optional(),
		field.Time("last_notified_at").Optional(),
		field.Bool("is_activated").Default(true),
	}
}

// Edges of the NotificationConfig.
func (NotificationConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("notification_config").
			Unique(),
	}
}
