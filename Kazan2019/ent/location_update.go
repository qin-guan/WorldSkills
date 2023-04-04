// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/departmentlocation"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/location"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "Name" field.
func (lu *LocationUpdate) SetName(s string) *LocationUpdate {
	lu.mutation.SetName(s)
	return lu
}

// AddDepartmentLocationIDs adds the "DepartmentLocation" edge to the DepartmentLocation entity by IDs.
func (lu *LocationUpdate) AddDepartmentLocationIDs(ids ...int) *LocationUpdate {
	lu.mutation.AddDepartmentLocationIDs(ids...)
	return lu
}

// AddDepartmentLocation adds the "DepartmentLocation" edges to the DepartmentLocation entity.
func (lu *LocationUpdate) AddDepartmentLocation(d ...*DepartmentLocation) *LocationUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return lu.AddDepartmentLocationIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// ClearDepartmentLocation clears all "DepartmentLocation" edges to the DepartmentLocation entity.
func (lu *LocationUpdate) ClearDepartmentLocation() *LocationUpdate {
	lu.mutation.ClearDepartmentLocation()
	return lu
}

// RemoveDepartmentLocationIDs removes the "DepartmentLocation" edge to DepartmentLocation entities by IDs.
func (lu *LocationUpdate) RemoveDepartmentLocationIDs(ids ...int) *LocationUpdate {
	lu.mutation.RemoveDepartmentLocationIDs(ids...)
	return lu
}

// RemoveDepartmentLocation removes "DepartmentLocation" edges to DepartmentLocation entities.
func (lu *LocationUpdate) RemoveDepartmentLocation(d ...*DepartmentLocation) *LocationUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return lu.RemoveDepartmentLocationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LocationMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if lu.mutation.DepartmentLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.DepartmentLocationTable,
			Columns: []string{location.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedDepartmentLocationIDs(); len(nodes) > 0 && !lu.mutation.DepartmentLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.DepartmentLocationTable,
			Columns: []string{location.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.DepartmentLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.DepartmentLocationTable,
			Columns: []string{location.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetName sets the "Name" field.
func (luo *LocationUpdateOne) SetName(s string) *LocationUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// AddDepartmentLocationIDs adds the "DepartmentLocation" edge to the DepartmentLocation entity by IDs.
func (luo *LocationUpdateOne) AddDepartmentLocationIDs(ids ...int) *LocationUpdateOne {
	luo.mutation.AddDepartmentLocationIDs(ids...)
	return luo
}

// AddDepartmentLocation adds the "DepartmentLocation" edges to the DepartmentLocation entity.
func (luo *LocationUpdateOne) AddDepartmentLocation(d ...*DepartmentLocation) *LocationUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return luo.AddDepartmentLocationIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// ClearDepartmentLocation clears all "DepartmentLocation" edges to the DepartmentLocation entity.
func (luo *LocationUpdateOne) ClearDepartmentLocation() *LocationUpdateOne {
	luo.mutation.ClearDepartmentLocation()
	return luo
}

// RemoveDepartmentLocationIDs removes the "DepartmentLocation" edge to DepartmentLocation entities by IDs.
func (luo *LocationUpdateOne) RemoveDepartmentLocationIDs(ids ...int) *LocationUpdateOne {
	luo.mutation.RemoveDepartmentLocationIDs(ids...)
	return luo
}

// RemoveDepartmentLocation removes "DepartmentLocation" edges to DepartmentLocation entities.
func (luo *LocationUpdateOne) RemoveDepartmentLocation(d ...*DepartmentLocation) *LocationUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return luo.RemoveDepartmentLocationIDs(ids...)
}

// Where appends a list predicates to the LocationUpdate builder.
func (luo *LocationUpdateOne) Where(ps ...predicate.Location) *LocationUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	return withHooks[*Location, LocationMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if luo.mutation.DepartmentLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.DepartmentLocationTable,
			Columns: []string{location.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedDepartmentLocationIDs(); len(nodes) > 0 && !luo.mutation.DepartmentLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.DepartmentLocationTable,
			Columns: []string{location.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.DepartmentLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.DepartmentLocationTable,
			Columns: []string{location.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}