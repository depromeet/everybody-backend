// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/predicate"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

// NotificationConfigUpdate is the builder for updating NotificationConfig entities.
type NotificationConfigUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationConfigMutation
}

// Where appends a list predicates to the NotificationConfigUpdate builder.
func (ncu *NotificationConfigUpdate) Where(ps ...predicate.NotificationConfig) *NotificationConfigUpdate {
	ncu.mutation.Where(ps...)
	return ncu
}

// SetInterval sets the "interval" field.
func (ncu *NotificationConfigUpdate) SetInterval(i int) *NotificationConfigUpdate {
	ncu.mutation.ResetInterval()
	ncu.mutation.SetInterval(i)
	return ncu
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (ncu *NotificationConfigUpdate) SetNillableInterval(i *int) *NotificationConfigUpdate {
	if i != nil {
		ncu.SetInterval(*i)
	}
	return ncu
}

// AddInterval adds i to the "interval" field.
func (ncu *NotificationConfigUpdate) AddInterval(i int) *NotificationConfigUpdate {
	ncu.mutation.AddInterval(i)
	return ncu
}

// ClearInterval clears the value of the "interval" field.
func (ncu *NotificationConfigUpdate) ClearInterval() *NotificationConfigUpdate {
	ncu.mutation.ClearInterval()
	return ncu
}

// SetLastNotifiedAt sets the "last_notified_at" field.
func (ncu *NotificationConfigUpdate) SetLastNotifiedAt(t time.Time) *NotificationConfigUpdate {
	ncu.mutation.SetLastNotifiedAt(t)
	return ncu
}

// SetNillableLastNotifiedAt sets the "last_notified_at" field if the given value is not nil.
func (ncu *NotificationConfigUpdate) SetNillableLastNotifiedAt(t *time.Time) *NotificationConfigUpdate {
	if t != nil {
		ncu.SetLastNotifiedAt(*t)
	}
	return ncu
}

// ClearLastNotifiedAt clears the value of the "last_notified_at" field.
func (ncu *NotificationConfigUpdate) ClearLastNotifiedAt() *NotificationConfigUpdate {
	ncu.mutation.ClearLastNotifiedAt()
	return ncu
}

// SetIsActivated sets the "is_activated" field.
func (ncu *NotificationConfigUpdate) SetIsActivated(b bool) *NotificationConfigUpdate {
	ncu.mutation.SetIsActivated(b)
	return ncu
}

// SetNillableIsActivated sets the "is_activated" field if the given value is not nil.
func (ncu *NotificationConfigUpdate) SetNillableIsActivated(b *bool) *NotificationConfigUpdate {
	if b != nil {
		ncu.SetIsActivated(*b)
	}
	return ncu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ncu *NotificationConfigUpdate) SetUserID(id int) *NotificationConfigUpdate {
	ncu.mutation.SetUserID(id)
	return ncu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ncu *NotificationConfigUpdate) SetNillableUserID(id *int) *NotificationConfigUpdate {
	if id != nil {
		ncu = ncu.SetUserID(*id)
	}
	return ncu
}

// SetUser sets the "user" edge to the User entity.
func (ncu *NotificationConfigUpdate) SetUser(u *User) *NotificationConfigUpdate {
	return ncu.SetUserID(u.ID)
}

// Mutation returns the NotificationConfigMutation object of the builder.
func (ncu *NotificationConfigUpdate) Mutation() *NotificationConfigMutation {
	return ncu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ncu *NotificationConfigUpdate) ClearUser() *NotificationConfigUpdate {
	ncu.mutation.ClearUser()
	return ncu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ncu *NotificationConfigUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ncu.hooks) == 0 {
		affected, err = ncu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncu.mutation = mutation
			affected, err = ncu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ncu.hooks) - 1; i >= 0; i-- {
			if ncu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncu *NotificationConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := ncu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ncu *NotificationConfigUpdate) Exec(ctx context.Context) error {
	_, err := ncu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncu *NotificationConfigUpdate) ExecX(ctx context.Context) {
	if err := ncu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ncu *NotificationConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notificationconfig.Table,
			Columns: notificationconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notificationconfig.FieldID,
			},
		},
	}
	if ps := ncu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncu.mutation.Interval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: notificationconfig.FieldInterval,
		})
	}
	if value, ok := ncu.mutation.AddedInterval(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: notificationconfig.FieldInterval,
		})
	}
	if ncu.mutation.IntervalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: notificationconfig.FieldInterval,
		})
	}
	if value, ok := ncu.mutation.LastNotifiedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notificationconfig.FieldLastNotifiedAt,
		})
	}
	if ncu.mutation.LastNotifiedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: notificationconfig.FieldLastNotifiedAt,
		})
	}
	if value, ok := ncu.mutation.IsActivated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notificationconfig.FieldIsActivated,
		})
	}
	if ncu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationconfig.UserTable,
			Columns: []string{notificationconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ncu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationconfig.UserTable,
			Columns: []string{notificationconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ncu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NotificationConfigUpdateOne is the builder for updating a single NotificationConfig entity.
type NotificationConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationConfigMutation
}

// SetInterval sets the "interval" field.
func (ncuo *NotificationConfigUpdateOne) SetInterval(i int) *NotificationConfigUpdateOne {
	ncuo.mutation.ResetInterval()
	ncuo.mutation.SetInterval(i)
	return ncuo
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (ncuo *NotificationConfigUpdateOne) SetNillableInterval(i *int) *NotificationConfigUpdateOne {
	if i != nil {
		ncuo.SetInterval(*i)
	}
	return ncuo
}

// AddInterval adds i to the "interval" field.
func (ncuo *NotificationConfigUpdateOne) AddInterval(i int) *NotificationConfigUpdateOne {
	ncuo.mutation.AddInterval(i)
	return ncuo
}

// ClearInterval clears the value of the "interval" field.
func (ncuo *NotificationConfigUpdateOne) ClearInterval() *NotificationConfigUpdateOne {
	ncuo.mutation.ClearInterval()
	return ncuo
}

// SetLastNotifiedAt sets the "last_notified_at" field.
func (ncuo *NotificationConfigUpdateOne) SetLastNotifiedAt(t time.Time) *NotificationConfigUpdateOne {
	ncuo.mutation.SetLastNotifiedAt(t)
	return ncuo
}

// SetNillableLastNotifiedAt sets the "last_notified_at" field if the given value is not nil.
func (ncuo *NotificationConfigUpdateOne) SetNillableLastNotifiedAt(t *time.Time) *NotificationConfigUpdateOne {
	if t != nil {
		ncuo.SetLastNotifiedAt(*t)
	}
	return ncuo
}

// ClearLastNotifiedAt clears the value of the "last_notified_at" field.
func (ncuo *NotificationConfigUpdateOne) ClearLastNotifiedAt() *NotificationConfigUpdateOne {
	ncuo.mutation.ClearLastNotifiedAt()
	return ncuo
}

// SetIsActivated sets the "is_activated" field.
func (ncuo *NotificationConfigUpdateOne) SetIsActivated(b bool) *NotificationConfigUpdateOne {
	ncuo.mutation.SetIsActivated(b)
	return ncuo
}

// SetNillableIsActivated sets the "is_activated" field if the given value is not nil.
func (ncuo *NotificationConfigUpdateOne) SetNillableIsActivated(b *bool) *NotificationConfigUpdateOne {
	if b != nil {
		ncuo.SetIsActivated(*b)
	}
	return ncuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ncuo *NotificationConfigUpdateOne) SetUserID(id int) *NotificationConfigUpdateOne {
	ncuo.mutation.SetUserID(id)
	return ncuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ncuo *NotificationConfigUpdateOne) SetNillableUserID(id *int) *NotificationConfigUpdateOne {
	if id != nil {
		ncuo = ncuo.SetUserID(*id)
	}
	return ncuo
}

// SetUser sets the "user" edge to the User entity.
func (ncuo *NotificationConfigUpdateOne) SetUser(u *User) *NotificationConfigUpdateOne {
	return ncuo.SetUserID(u.ID)
}

// Mutation returns the NotificationConfigMutation object of the builder.
func (ncuo *NotificationConfigUpdateOne) Mutation() *NotificationConfigMutation {
	return ncuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ncuo *NotificationConfigUpdateOne) ClearUser() *NotificationConfigUpdateOne {
	ncuo.mutation.ClearUser()
	return ncuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ncuo *NotificationConfigUpdateOne) Select(field string, fields ...string) *NotificationConfigUpdateOne {
	ncuo.fields = append([]string{field}, fields...)
	return ncuo
}

// Save executes the query and returns the updated NotificationConfig entity.
func (ncuo *NotificationConfigUpdateOne) Save(ctx context.Context) (*NotificationConfig, error) {
	var (
		err  error
		node *NotificationConfig
	)
	if len(ncuo.hooks) == 0 {
		node, err = ncuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncuo.mutation = mutation
			node, err = ncuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ncuo.hooks) - 1; i >= 0; i-- {
			if ncuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncuo *NotificationConfigUpdateOne) SaveX(ctx context.Context) *NotificationConfig {
	node, err := ncuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ncuo *NotificationConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := ncuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncuo *NotificationConfigUpdateOne) ExecX(ctx context.Context) {
	if err := ncuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ncuo *NotificationConfigUpdateOne) sqlSave(ctx context.Context) (_node *NotificationConfig, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notificationconfig.Table,
			Columns: notificationconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notificationconfig.FieldID,
			},
		},
	}
	id, ok := ncuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing NotificationConfig.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ncuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationconfig.FieldID)
		for _, f := range fields {
			if !notificationconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notificationconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ncuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncuo.mutation.Interval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: notificationconfig.FieldInterval,
		})
	}
	if value, ok := ncuo.mutation.AddedInterval(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: notificationconfig.FieldInterval,
		})
	}
	if ncuo.mutation.IntervalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: notificationconfig.FieldInterval,
		})
	}
	if value, ok := ncuo.mutation.LastNotifiedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notificationconfig.FieldLastNotifiedAt,
		})
	}
	if ncuo.mutation.LastNotifiedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: notificationconfig.FieldLastNotifiedAt,
		})
	}
	if value, ok := ncuo.mutation.IsActivated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notificationconfig.FieldIsActivated,
		})
	}
	if ncuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationconfig.UserTable,
			Columns: []string{notificationconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ncuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notificationconfig.UserTable,
			Columns: []string{notificationconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotificationConfig{config: ncuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ncuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
