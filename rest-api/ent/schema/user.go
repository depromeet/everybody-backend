package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		// struct tag를 정의 안 해주면 deviceToken으로 직렬화됨
		field.String("deviceToken").StructTag("json:\"device_token\""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
