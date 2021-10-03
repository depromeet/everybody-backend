// Code generated by entc, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/depromeet/everybody-backend/rest-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DeviceToken applies equality check predicate on the "device_token" field. It's identical to DeviceTokenEQ.
func DeviceToken(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceToken), v))
	})
}

// PushToken applies equality check predicate on the "push_token" field. It's identical to PushTokenEQ.
func PushToken(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPushToken), v))
	})
}

// DeviceTokenEQ applies the EQ predicate on the "device_token" field.
func DeviceTokenEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenNEQ applies the NEQ predicate on the "device_token" field.
func DeviceTokenNEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenIn applies the In predicate on the "device_token" field.
func DeviceTokenIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeviceToken), v...))
	})
}

// DeviceTokenNotIn applies the NotIn predicate on the "device_token" field.
func DeviceTokenNotIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeviceToken), v...))
	})
}

// DeviceTokenGT applies the GT predicate on the "device_token" field.
func DeviceTokenGT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenGTE applies the GTE predicate on the "device_token" field.
func DeviceTokenGTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenLT applies the LT predicate on the "device_token" field.
func DeviceTokenLT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenLTE applies the LTE predicate on the "device_token" field.
func DeviceTokenLTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenContains applies the Contains predicate on the "device_token" field.
func DeviceTokenContains(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenHasPrefix applies the HasPrefix predicate on the "device_token" field.
func DeviceTokenHasPrefix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenHasSuffix applies the HasSuffix predicate on the "device_token" field.
func DeviceTokenHasSuffix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenEqualFold applies the EqualFold predicate on the "device_token" field.
func DeviceTokenEqualFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDeviceToken), v))
	})
}

// DeviceTokenContainsFold applies the ContainsFold predicate on the "device_token" field.
func DeviceTokenContainsFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDeviceToken), v))
	})
}

// PushTokenEQ applies the EQ predicate on the "push_token" field.
func PushTokenEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPushToken), v))
	})
}

// PushTokenNEQ applies the NEQ predicate on the "push_token" field.
func PushTokenNEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPushToken), v))
	})
}

// PushTokenIn applies the In predicate on the "push_token" field.
func PushTokenIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPushToken), v...))
	})
}

// PushTokenNotIn applies the NotIn predicate on the "push_token" field.
func PushTokenNotIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPushToken), v...))
	})
}

// PushTokenGT applies the GT predicate on the "push_token" field.
func PushTokenGT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPushToken), v))
	})
}

// PushTokenGTE applies the GTE predicate on the "push_token" field.
func PushTokenGTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPushToken), v))
	})
}

// PushTokenLT applies the LT predicate on the "push_token" field.
func PushTokenLT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPushToken), v))
	})
}

// PushTokenLTE applies the LTE predicate on the "push_token" field.
func PushTokenLTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPushToken), v))
	})
}

// PushTokenContains applies the Contains predicate on the "push_token" field.
func PushTokenContains(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPushToken), v))
	})
}

// PushTokenHasPrefix applies the HasPrefix predicate on the "push_token" field.
func PushTokenHasPrefix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPushToken), v))
	})
}

// PushTokenHasSuffix applies the HasSuffix predicate on the "push_token" field.
func PushTokenHasSuffix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPushToken), v))
	})
}

// PushTokenEqualFold applies the EqualFold predicate on the "push_token" field.
func PushTokenEqualFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPushToken), v))
	})
}

// PushTokenContainsFold applies the ContainsFold predicate on the "push_token" field.
func PushTokenContainsFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPushToken), v))
	})
}

// DeviceOsEQ applies the EQ predicate on the "device_os" field.
func DeviceOsEQ(v DeviceOs) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceOs), v))
	})
}

// DeviceOsNEQ applies the NEQ predicate on the "device_os" field.
func DeviceOsNEQ(v DeviceOs) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeviceOs), v))
	})
}

// DeviceOsIn applies the In predicate on the "device_os" field.
func DeviceOsIn(vs ...DeviceOs) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeviceOs), v...))
	})
}

// DeviceOsNotIn applies the NotIn predicate on the "device_os" field.
func DeviceOsNotIn(vs ...DeviceOs) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeviceOs), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		p(s.Not())
	})
}