// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/employee"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetFirstName sets the "FirstName" field.
func (eu *EmployeeUpdate) SetFirstName(s string) *EmployeeUpdate {
	eu.mutation.SetFirstName(s)
	return eu
}

// SetLastName sets the "LastName" field.
func (eu *EmployeeUpdate) SetLastName(s string) *EmployeeUpdate {
	eu.mutation.SetLastName(s)
	return eu
}

// SetPhone sets the "Phone" field.
func (eu *EmployeeUpdate) SetPhone(s string) *EmployeeUpdate {
	eu.mutation.SetPhone(s)
	return eu
}

// AddAssetIDs adds the "Asset" edge to the Asset entity by IDs.
func (eu *EmployeeUpdate) AddAssetIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.AddAssetIDs(ids...)
	return eu
}

// AddAsset adds the "Asset" edges to the Asset entity.
func (eu *EmployeeUpdate) AddAsset(a ...*Asset) *EmployeeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.AddAssetIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// ClearAsset clears all "Asset" edges to the Asset entity.
func (eu *EmployeeUpdate) ClearAsset() *EmployeeUpdate {
	eu.mutation.ClearAsset()
	return eu
}

// RemoveAssetIDs removes the "Asset" edge to Asset entities by IDs.
func (eu *EmployeeUpdate) RemoveAssetIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.RemoveAssetIDs(ids...)
	return eu
}

// RemoveAsset removes "Asset" edges to Asset entities.
func (eu *EmployeeUpdate) RemoveAsset(a ...*Asset) *EmployeeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.RemoveAssetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EmployeeMutation](ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
	}
	if value, ok := eu.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Phone(); ok {
		_spec.SetField(employee.FieldPhone, field.TypeString, value)
	}
	if eu.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AssetTable,
			Columns: []string{employee.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedAssetIDs(); len(nodes) > 0 && !eu.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AssetTable,
			Columns: []string{employee.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AssetTable,
			Columns: []string{employee.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetFirstName sets the "FirstName" field.
func (euo *EmployeeUpdateOne) SetFirstName(s string) *EmployeeUpdateOne {
	euo.mutation.SetFirstName(s)
	return euo
}

// SetLastName sets the "LastName" field.
func (euo *EmployeeUpdateOne) SetLastName(s string) *EmployeeUpdateOne {
	euo.mutation.SetLastName(s)
	return euo
}

// SetPhone sets the "Phone" field.
func (euo *EmployeeUpdateOne) SetPhone(s string) *EmployeeUpdateOne {
	euo.mutation.SetPhone(s)
	return euo
}

// AddAssetIDs adds the "Asset" edge to the Asset entity by IDs.
func (euo *EmployeeUpdateOne) AddAssetIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.AddAssetIDs(ids...)
	return euo
}

// AddAsset adds the "Asset" edges to the Asset entity.
func (euo *EmployeeUpdateOne) AddAsset(a ...*Asset) *EmployeeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.AddAssetIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// ClearAsset clears all "Asset" edges to the Asset entity.
func (euo *EmployeeUpdateOne) ClearAsset() *EmployeeUpdateOne {
	euo.mutation.ClearAsset()
	return euo
}

// RemoveAssetIDs removes the "Asset" edge to Asset entities by IDs.
func (euo *EmployeeUpdateOne) RemoveAssetIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.RemoveAssetIDs(ids...)
	return euo
}

// RemoveAsset removes "Asset" edges to Asset entities.
func (euo *EmployeeUpdateOne) RemoveAsset(a ...*Asset) *EmployeeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.RemoveAssetIDs(ids...)
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (euo *EmployeeUpdateOne) Where(ps ...predicate.Employee) *EmployeeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmployeeUpdateOne) Select(field string, fields ...string) *EmployeeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Employee entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	return withHooks[*Employee, EmployeeMutation](ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (_node *Employee, err error) {
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Employee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for _, f := range fields {
			if !employee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
	}
	if value, ok := euo.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Phone(); ok {
		_spec.SetField(employee.FieldPhone, field.TypeString, value)
	}
	if euo.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AssetTable,
			Columns: []string{employee.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedAssetIDs(); len(nodes) > 0 && !euo.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AssetTable,
			Columns: []string{employee.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.AssetTable,
			Columns: []string{employee.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Employee{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
