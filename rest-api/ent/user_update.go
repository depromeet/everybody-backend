// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/device"
	"github.com/depromeet/everybody-backend/rest-api/ent/notificationconfig"
	"github.com/depromeet/everybody-backend/rest-api/ent/picture"
	"github.com/depromeet/everybody-backend/rest-api/ent/predicate"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetHeight sets the "height" field.
func (uu *UserUpdate) SetHeight(i int) *UserUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(i)
	return uu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHeight(i *int) *UserUpdate {
	if i != nil {
		uu.SetHeight(*i)
	}
	return uu
}

// AddHeight adds i to the "height" field.
func (uu *UserUpdate) AddHeight(i int) *UserUpdate {
	uu.mutation.AddHeight(i)
	return uu
}

// ClearHeight clears the value of the "height" field.
func (uu *UserUpdate) ClearHeight() *UserUpdate {
	uu.mutation.ClearHeight()
	return uu
}

// SetWeight sets the "weight" field.
func (uu *UserUpdate) SetWeight(i int) *UserUpdate {
	uu.mutation.ResetWeight()
	uu.mutation.SetWeight(i)
	return uu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWeight(i *int) *UserUpdate {
	if i != nil {
		uu.SetWeight(*i)
	}
	return uu
}

// AddWeight adds i to the "weight" field.
func (uu *UserUpdate) AddWeight(i int) *UserUpdate {
	uu.mutation.AddWeight(i)
	return uu
}

// ClearWeight clears the value of the "weight" field.
func (uu *UserUpdate) ClearWeight() *UserUpdate {
	uu.mutation.ClearWeight()
	return uu
}

// SetType sets the "type" field.
func (uu *UserUpdate) SetType(u user.Type) *UserUpdate {
	uu.mutation.SetType(u)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// AddDeviceIDs adds the "device" edge to the Device entity by IDs.
func (uu *UserUpdate) AddDeviceIDs(ids ...int) *UserUpdate {
	uu.mutation.AddDeviceIDs(ids...)
	return uu
}

// AddDevice adds the "device" edges to the Device entity.
func (uu *UserUpdate) AddDevice(d ...*Device) *UserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.AddDeviceIDs(ids...)
}

// AddNotificationConfigIDs adds the "notification_config" edge to the NotificationConfig entity by IDs.
func (uu *UserUpdate) AddNotificationConfigIDs(ids ...int) *UserUpdate {
	uu.mutation.AddNotificationConfigIDs(ids...)
	return uu
}

// AddNotificationConfig adds the "notification_config" edges to the NotificationConfig entity.
func (uu *UserUpdate) AddNotificationConfig(n ...*NotificationConfig) *UserUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uu.AddNotificationConfigIDs(ids...)
}

// AddAlbumIDs adds the "album" edge to the Album entity by IDs.
func (uu *UserUpdate) AddAlbumIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAlbumIDs(ids...)
	return uu
}

// AddAlbum adds the "album" edges to the Album entity.
func (uu *UserUpdate) AddAlbum(a ...*Album) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAlbumIDs(ids...)
}

// AddPictureIDs adds the "picture" edge to the Picture entity by IDs.
func (uu *UserUpdate) AddPictureIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPictureIDs(ids...)
	return uu
}

// AddPicture adds the "picture" edges to the Picture entity.
func (uu *UserUpdate) AddPicture(p ...*Picture) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPictureIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearDevice clears all "device" edges to the Device entity.
func (uu *UserUpdate) ClearDevice() *UserUpdate {
	uu.mutation.ClearDevice()
	return uu
}

// RemoveDeviceIDs removes the "device" edge to Device entities by IDs.
func (uu *UserUpdate) RemoveDeviceIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveDeviceIDs(ids...)
	return uu
}

// RemoveDevice removes "device" edges to Device entities.
func (uu *UserUpdate) RemoveDevice(d ...*Device) *UserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.RemoveDeviceIDs(ids...)
}

// ClearNotificationConfig clears all "notification_config" edges to the NotificationConfig entity.
func (uu *UserUpdate) ClearNotificationConfig() *UserUpdate {
	uu.mutation.ClearNotificationConfig()
	return uu
}

// RemoveNotificationConfigIDs removes the "notification_config" edge to NotificationConfig entities by IDs.
func (uu *UserUpdate) RemoveNotificationConfigIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveNotificationConfigIDs(ids...)
	return uu
}

// RemoveNotificationConfig removes "notification_config" edges to NotificationConfig entities.
func (uu *UserUpdate) RemoveNotificationConfig(n ...*NotificationConfig) *UserUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uu.RemoveNotificationConfigIDs(ids...)
}

// ClearAlbum clears all "album" edges to the Album entity.
func (uu *UserUpdate) ClearAlbum() *UserUpdate {
	uu.mutation.ClearAlbum()
	return uu
}

// RemoveAlbumIDs removes the "album" edge to Album entities by IDs.
func (uu *UserUpdate) RemoveAlbumIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAlbumIDs(ids...)
	return uu
}

// RemoveAlbum removes "album" edges to Album entities.
func (uu *UserUpdate) RemoveAlbum(a ...*Album) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAlbumIDs(ids...)
}

// ClearPicture clears all "picture" edges to the Picture entity.
func (uu *UserUpdate) ClearPicture() *UserUpdate {
	uu.mutation.ClearPicture()
	return uu
}

// RemovePictureIDs removes the "picture" edge to Picture entities by IDs.
func (uu *UserUpdate) RemovePictureIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePictureIDs(ids...)
	return uu
}

// RemovePicture removes "picture" edges to Picture entities.
func (uu *UserUpdate) RemovePicture(p ...*Picture) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePictureIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.GetType(); ok {
		if err := user.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHeight,
		})
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHeight,
		})
	}
	if uu.mutation.HeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldHeight,
		})
	}
	if value, ok := uu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if uu.mutation.WeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldType,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if uu.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedDeviceIDs(); len(nodes) > 0 && !uu.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.NotificationConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotificationConfigTable,
			Columns: []string{user.NotificationConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: notificationconfig.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedNotificationConfigIDs(); len(nodes) > 0 && !uu.mutation.NotificationConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotificationConfigTable,
			Columns: []string{user.NotificationConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: notificationconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.NotificationConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotificationConfigTable,
			Columns: []string{user.NotificationConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: notificationconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlbumTable,
			Columns: []string{user.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAlbumIDs(); len(nodes) > 0 && !uu.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlbumTable,
			Columns: []string{user.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AlbumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlbumTable,
			Columns: []string{user.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PictureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PictureTable,
			Columns: []string{user.PictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPictureIDs(); len(nodes) > 0 && !uu.mutation.PictureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PictureTable,
			Columns: []string{user.PictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PictureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PictureTable,
			Columns: []string{user.PictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UserUpdateOne) SetHeight(i int) *UserUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(i)
	return uuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHeight(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetHeight(*i)
	}
	return uuo
}

// AddHeight adds i to the "height" field.
func (uuo *UserUpdateOne) AddHeight(i int) *UserUpdateOne {
	uuo.mutation.AddHeight(i)
	return uuo
}

// ClearHeight clears the value of the "height" field.
func (uuo *UserUpdateOne) ClearHeight() *UserUpdateOne {
	uuo.mutation.ClearHeight()
	return uuo
}

// SetWeight sets the "weight" field.
func (uuo *UserUpdateOne) SetWeight(i int) *UserUpdateOne {
	uuo.mutation.ResetWeight()
	uuo.mutation.SetWeight(i)
	return uuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWeight(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetWeight(*i)
	}
	return uuo
}

// AddWeight adds i to the "weight" field.
func (uuo *UserUpdateOne) AddWeight(i int) *UserUpdateOne {
	uuo.mutation.AddWeight(i)
	return uuo
}

// ClearWeight clears the value of the "weight" field.
func (uuo *UserUpdateOne) ClearWeight() *UserUpdateOne {
	uuo.mutation.ClearWeight()
	return uuo
}

// SetType sets the "type" field.
func (uuo *UserUpdateOne) SetType(u user.Type) *UserUpdateOne {
	uuo.mutation.SetType(u)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// AddDeviceIDs adds the "device" edge to the Device entity by IDs.
func (uuo *UserUpdateOne) AddDeviceIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddDeviceIDs(ids...)
	return uuo
}

// AddDevice adds the "device" edges to the Device entity.
func (uuo *UserUpdateOne) AddDevice(d ...*Device) *UserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.AddDeviceIDs(ids...)
}

// AddNotificationConfigIDs adds the "notification_config" edge to the NotificationConfig entity by IDs.
func (uuo *UserUpdateOne) AddNotificationConfigIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddNotificationConfigIDs(ids...)
	return uuo
}

// AddNotificationConfig adds the "notification_config" edges to the NotificationConfig entity.
func (uuo *UserUpdateOne) AddNotificationConfig(n ...*NotificationConfig) *UserUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uuo.AddNotificationConfigIDs(ids...)
}

// AddAlbumIDs adds the "album" edge to the Album entity by IDs.
func (uuo *UserUpdateOne) AddAlbumIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAlbumIDs(ids...)
	return uuo
}

// AddAlbum adds the "album" edges to the Album entity.
func (uuo *UserUpdateOne) AddAlbum(a ...*Album) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAlbumIDs(ids...)
}

// AddPictureIDs adds the "picture" edge to the Picture entity by IDs.
func (uuo *UserUpdateOne) AddPictureIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPictureIDs(ids...)
	return uuo
}

// AddPicture adds the "picture" edges to the Picture entity.
func (uuo *UserUpdateOne) AddPicture(p ...*Picture) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPictureIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearDevice clears all "device" edges to the Device entity.
func (uuo *UserUpdateOne) ClearDevice() *UserUpdateOne {
	uuo.mutation.ClearDevice()
	return uuo
}

// RemoveDeviceIDs removes the "device" edge to Device entities by IDs.
func (uuo *UserUpdateOne) RemoveDeviceIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveDeviceIDs(ids...)
	return uuo
}

// RemoveDevice removes "device" edges to Device entities.
func (uuo *UserUpdateOne) RemoveDevice(d ...*Device) *UserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.RemoveDeviceIDs(ids...)
}

// ClearNotificationConfig clears all "notification_config" edges to the NotificationConfig entity.
func (uuo *UserUpdateOne) ClearNotificationConfig() *UserUpdateOne {
	uuo.mutation.ClearNotificationConfig()
	return uuo
}

// RemoveNotificationConfigIDs removes the "notification_config" edge to NotificationConfig entities by IDs.
func (uuo *UserUpdateOne) RemoveNotificationConfigIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveNotificationConfigIDs(ids...)
	return uuo
}

// RemoveNotificationConfig removes "notification_config" edges to NotificationConfig entities.
func (uuo *UserUpdateOne) RemoveNotificationConfig(n ...*NotificationConfig) *UserUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uuo.RemoveNotificationConfigIDs(ids...)
}

// ClearAlbum clears all "album" edges to the Album entity.
func (uuo *UserUpdateOne) ClearAlbum() *UserUpdateOne {
	uuo.mutation.ClearAlbum()
	return uuo
}

// RemoveAlbumIDs removes the "album" edge to Album entities by IDs.
func (uuo *UserUpdateOne) RemoveAlbumIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAlbumIDs(ids...)
	return uuo
}

// RemoveAlbum removes "album" edges to Album entities.
func (uuo *UserUpdateOne) RemoveAlbum(a ...*Album) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAlbumIDs(ids...)
}

// ClearPicture clears all "picture" edges to the Picture entity.
func (uuo *UserUpdateOne) ClearPicture() *UserUpdateOne {
	uuo.mutation.ClearPicture()
	return uuo
}

// RemovePictureIDs removes the "picture" edge to Picture entities by IDs.
func (uuo *UserUpdateOne) RemovePictureIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePictureIDs(ids...)
	return uuo
}

// RemovePicture removes "picture" edges to Picture entities.
func (uuo *UserUpdateOne) RemovePicture(p ...*Picture) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePictureIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.GetType(); ok {
		if err := user.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHeight,
		})
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHeight,
		})
	}
	if uuo.mutation.HeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldHeight,
		})
	}
	if value, ok := uuo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uuo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if uuo.mutation.WeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldType,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if uuo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedDeviceIDs(); len(nodes) > 0 && !uuo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.NotificationConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotificationConfigTable,
			Columns: []string{user.NotificationConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: notificationconfig.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedNotificationConfigIDs(); len(nodes) > 0 && !uuo.mutation.NotificationConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotificationConfigTable,
			Columns: []string{user.NotificationConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: notificationconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.NotificationConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotificationConfigTable,
			Columns: []string{user.NotificationConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: notificationconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlbumTable,
			Columns: []string{user.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAlbumIDs(); len(nodes) > 0 && !uuo.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlbumTable,
			Columns: []string{user.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AlbumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlbumTable,
			Columns: []string{user.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PictureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PictureTable,
			Columns: []string{user.PictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPictureIDs(); len(nodes) > 0 && !uuo.mutation.PictureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PictureTable,
			Columns: []string{user.PictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PictureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PictureTable,
			Columns: []string{user.PictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
