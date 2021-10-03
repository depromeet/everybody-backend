package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

var (
	deviceOSes = []string{"ANDROID", "IOS"}
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("device_token"),
		field.String("push_token"),
		field.Enum("device_os").
			Values(deviceOSes...),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("device").
			Unique(),
	}
}
