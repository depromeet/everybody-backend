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
		// 요일을 또 다른 테이블로 정의한 뒤 Join하는 것보단 7개로 정적이다보니
		// 그냥 column으로 나열하는 게 좋을 듯
		field.Bool("monday").Default(false),
		field.Bool("tuesday").Default(false),
		field.Bool("thursday").Default(false),
		field.Bool("friday").Default(false),
		field.Bool("saturday").Default(false),
		field.Bool("sunday").Default(false),
		// 원하는 알림 시각
		field.String("preferred_time_hour").Optional(),
		field.String("preferred_time_minute").Optional(),
		// 최근 알림 보낸 게 오늘(같은 요일) 이전일 때에만 알림을 보낸다.
		// 최근 알림을 보낸 게 언젠지 기록해야 오늘의 다음 Loop 때 이 유저에게 보낼지 말지를 판단할 수 있음.
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
