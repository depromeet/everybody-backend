// Code generated by entc, DO NOT EDIT.

package picture

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/depromeet/everybody-backend/rest-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BodyPart applies equality check predicate on the "body_part" field. It's identical to BodyPartEQ.
func BodyPart(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBodyPart), v))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// BodyPartEQ applies the EQ predicate on the "body_part" field.
func BodyPartEQ(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBodyPart), v))
	})
}

// BodyPartNEQ applies the NEQ predicate on the "body_part" field.
func BodyPartNEQ(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBodyPart), v))
	})
}

// BodyPartIn applies the In predicate on the "body_part" field.
func BodyPartIn(vs ...string) predicate.Picture {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Picture(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBodyPart), v...))
	})
}

// BodyPartNotIn applies the NotIn predicate on the "body_part" field.
func BodyPartNotIn(vs ...string) predicate.Picture {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Picture(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBodyPart), v...))
	})
}

// BodyPartGT applies the GT predicate on the "body_part" field.
func BodyPartGT(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBodyPart), v))
	})
}

// BodyPartGTE applies the GTE predicate on the "body_part" field.
func BodyPartGTE(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBodyPart), v))
	})
}

// BodyPartLT applies the LT predicate on the "body_part" field.
func BodyPartLT(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBodyPart), v))
	})
}

// BodyPartLTE applies the LTE predicate on the "body_part" field.
func BodyPartLTE(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBodyPart), v))
	})
}

// BodyPartContains applies the Contains predicate on the "body_part" field.
func BodyPartContains(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBodyPart), v))
	})
}

// BodyPartHasPrefix applies the HasPrefix predicate on the "body_part" field.
func BodyPartHasPrefix(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBodyPart), v))
	})
}

// BodyPartHasSuffix applies the HasSuffix predicate on the "body_part" field.
func BodyPartHasSuffix(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBodyPart), v))
	})
}

// BodyPartEqualFold applies the EqualFold predicate on the "body_part" field.
func BodyPartEqualFold(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBodyPart), v))
	})
}

// BodyPartContainsFold applies the ContainsFold predicate on the "body_part" field.
func BodyPartContainsFold(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBodyPart), v))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Picture {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Picture(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Picture {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Picture(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Picture {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Picture(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Picture {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Picture(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasAlbum applies the HasEdge predicate on the "album" edge.
func HasAlbum() predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlbumTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AlbumTable, AlbumColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlbumWith applies the HasEdge predicate on the "album" edge with a given conditions (other predicates).
func HasAlbumWith(preds ...predicate.Album) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlbumInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AlbumTable, AlbumColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
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
func And(predicates ...predicate.Picture) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Picture) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
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
func Not(p predicate.Picture) predicate.Picture {
	return predicate.Picture(func(s *sql.Selector) {
		p(s.Not())
	})
}
