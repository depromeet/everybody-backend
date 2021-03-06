package schema

import (
	"time"

	"entgo.io/ent/dialect/entsql"

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
		field.String("profile_image").Optional(),
		field.String("nickname"),
		field.String("motto").Default("눈바디와 함께 꾸준히 운동할테야!"),
		field.Int("height").Optional().Nillable(),
		field.Int("weight").Optional().Nillable(),
		field.Enum("kind").Values(kinds...),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("download_completed").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("notification_config", NotificationConfig.Type).
			Annotations(entsql.Annotation{
				// 유저를 지울 때 notification config도 지움
				OnDelete: entsql.Cascade,
			}).
			Unique(),
		edge.To("album", Album.Type),
		edge.To("picture", Picture.Type),
		edge.To("video", Video.Type),
	}
}
