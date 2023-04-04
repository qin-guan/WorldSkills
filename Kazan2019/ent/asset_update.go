// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assetgroup"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assetphoto"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assettransferlog"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/departmentlocation"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/employee"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// AssetUpdate is the builder for updating Asset entities.
type AssetUpdate struct {
	config
	hooks    []Hook
	mutation *AssetMutation
}

// Where appends a list predicates to the AssetUpdate builder.
func (au *AssetUpdate) Where(ps ...predicate.Asset) *AssetUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAssetGroupID sets the "AssetGroupID" field.
func (au *AssetUpdate) SetAssetGroupID(i int) *AssetUpdate {
	au.mutation.SetAssetGroupID(i)
	return au
}

// SetEmployeeID sets the "EmployeeID" field.
func (au *AssetUpdate) SetEmployeeID(i int) *AssetUpdate {
	au.mutation.SetEmployeeID(i)
	return au
}

// SetDepartmentLocationID sets the "DepartmentLocationID" field.
func (au *AssetUpdate) SetDepartmentLocationID(i int) *AssetUpdate {
	au.mutation.SetDepartmentLocationID(i)
	return au
}

// SetAssetSN sets the "AssetSN" field.
func (au *AssetUpdate) SetAssetSN(i int) *AssetUpdate {
	au.mutation.ResetAssetSN()
	au.mutation.SetAssetSN(i)
	return au
}

// AddAssetSN adds i to the "AssetSN" field.
func (au *AssetUpdate) AddAssetSN(i int) *AssetUpdate {
	au.mutation.AddAssetSN(i)
	return au
}

// SetAssetName sets the "AssetName" field.
func (au *AssetUpdate) SetAssetName(s string) *AssetUpdate {
	au.mutation.SetAssetName(s)
	return au
}

// SetDescription sets the "Description" field.
func (au *AssetUpdate) SetDescription(s string) *AssetUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetWarrantyDate sets the "WarrantyDate" field.
func (au *AssetUpdate) SetWarrantyDate(t time.Time) *AssetUpdate {
	au.mutation.SetWarrantyDate(t)
	return au
}

// AddAssetPhotoIDs adds the "AssetPhoto" edge to the AssetPhoto entity by IDs.
func (au *AssetUpdate) AddAssetPhotoIDs(ids ...int) *AssetUpdate {
	au.mutation.AddAssetPhotoIDs(ids...)
	return au
}

// AddAssetPhoto adds the "AssetPhoto" edges to the AssetPhoto entity.
func (au *AssetUpdate) AddAssetPhoto(a ...*AssetPhoto) *AssetUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAssetPhotoIDs(ids...)
}

// AddAssetTransferLogIDs adds the "AssetTransferLog" edge to the AssetTransferLog entity by IDs.
func (au *AssetUpdate) AddAssetTransferLogIDs(ids ...int) *AssetUpdate {
	au.mutation.AddAssetTransferLogIDs(ids...)
	return au
}

// AddAssetTransferLog adds the "AssetTransferLog" edges to the AssetTransferLog entity.
func (au *AssetUpdate) AddAssetTransferLog(a ...*AssetTransferLog) *AssetUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAssetTransferLogIDs(ids...)
}

// SetDepartmentLocation sets the "DepartmentLocation" edge to the DepartmentLocation entity.
func (au *AssetUpdate) SetDepartmentLocation(d *DepartmentLocation) *AssetUpdate {
	return au.SetDepartmentLocationID(d.ID)
}

// SetEmployee sets the "Employee" edge to the Employee entity.
func (au *AssetUpdate) SetEmployee(e *Employee) *AssetUpdate {
	return au.SetEmployeeID(e.ID)
}

// SetAssetGroup sets the "AssetGroup" edge to the AssetGroup entity.
func (au *AssetUpdate) SetAssetGroup(a *AssetGroup) *AssetUpdate {
	return au.SetAssetGroupID(a.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (au *AssetUpdate) Mutation() *AssetMutation {
	return au.mutation
}

// ClearAssetPhoto clears all "AssetPhoto" edges to the AssetPhoto entity.
func (au *AssetUpdate) ClearAssetPhoto() *AssetUpdate {
	au.mutation.ClearAssetPhoto()
	return au
}

// RemoveAssetPhotoIDs removes the "AssetPhoto" edge to AssetPhoto entities by IDs.
func (au *AssetUpdate) RemoveAssetPhotoIDs(ids ...int) *AssetUpdate {
	au.mutation.RemoveAssetPhotoIDs(ids...)
	return au
}

// RemoveAssetPhoto removes "AssetPhoto" edges to AssetPhoto entities.
func (au *AssetUpdate) RemoveAssetPhoto(a ...*AssetPhoto) *AssetUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAssetPhotoIDs(ids...)
}

// ClearAssetTransferLog clears all "AssetTransferLog" edges to the AssetTransferLog entity.
func (au *AssetUpdate) ClearAssetTransferLog() *AssetUpdate {
	au.mutation.ClearAssetTransferLog()
	return au
}

// RemoveAssetTransferLogIDs removes the "AssetTransferLog" edge to AssetTransferLog entities by IDs.
func (au *AssetUpdate) RemoveAssetTransferLogIDs(ids ...int) *AssetUpdate {
	au.mutation.RemoveAssetTransferLogIDs(ids...)
	return au
}

// RemoveAssetTransferLog removes "AssetTransferLog" edges to AssetTransferLog entities.
func (au *AssetUpdate) RemoveAssetTransferLog(a ...*AssetTransferLog) *AssetUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAssetTransferLogIDs(ids...)
}

// ClearDepartmentLocation clears the "DepartmentLocation" edge to the DepartmentLocation entity.
func (au *AssetUpdate) ClearDepartmentLocation() *AssetUpdate {
	au.mutation.ClearDepartmentLocation()
	return au
}

// ClearEmployee clears the "Employee" edge to the Employee entity.
func (au *AssetUpdate) ClearEmployee() *AssetUpdate {
	au.mutation.ClearEmployee()
	return au
}

// ClearAssetGroup clears the "AssetGroup" edge to the AssetGroup entity.
func (au *AssetUpdate) ClearAssetGroup() *AssetUpdate {
	au.mutation.ClearAssetGroup()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AssetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AssetMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AssetUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AssetUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AssetUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AssetUpdate) check() error {
	if v, ok := au.mutation.AssetSN(); ok {
		if err := asset.AssetSNValidator(v); err != nil {
			return &ValidationError{Name: "AssetSN", err: fmt.Errorf(`ent: validator failed for field "Asset.AssetSN": %w`, err)}
		}
	}
	if _, ok := au.mutation.DepartmentLocationID(); au.mutation.DepartmentLocationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.DepartmentLocation"`)
	}
	if _, ok := au.mutation.EmployeeID(); au.mutation.EmployeeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.Employee"`)
	}
	if _, ok := au.mutation.AssetGroupID(); au.mutation.AssetGroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.AssetGroup"`)
	}
	return nil
}

func (au *AssetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(asset.Table, asset.Columns, sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.AssetSN(); ok {
		_spec.SetField(asset.FieldAssetSN, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedAssetSN(); ok {
		_spec.AddField(asset.FieldAssetSN, field.TypeInt, value)
	}
	if value, ok := au.mutation.AssetName(); ok {
		_spec.SetField(asset.FieldAssetName, field.TypeString, value)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(asset.FieldDescription, field.TypeString, value)
	}
	if value, ok := au.mutation.WarrantyDate(); ok {
		_spec.SetField(asset.FieldWarrantyDate, field.TypeTime, value)
	}
	if au.mutation.AssetPhotoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetPhotoTable,
			Columns: []string{asset.AssetPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAssetPhotoIDs(); len(nodes) > 0 && !au.mutation.AssetPhotoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetPhotoTable,
			Columns: []string{asset.AssetPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AssetPhotoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetPhotoTable,
			Columns: []string{asset.AssetPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AssetTransferLogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetTransferLogTable,
			Columns: []string{asset.AssetTransferLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAssetTransferLogIDs(); len(nodes) > 0 && !au.mutation.AssetTransferLogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetTransferLogTable,
			Columns: []string{asset.AssetTransferLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AssetTransferLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetTransferLogTable,
			Columns: []string{asset.AssetTransferLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DepartmentLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.DepartmentLocationTable,
			Columns: []string{asset.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DepartmentLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.DepartmentLocationTable,
			Columns: []string{asset.DepartmentLocationColumn},
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
	if au.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.EmployeeTable,
			Columns: []string{asset.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.EmployeeTable,
			Columns: []string{asset.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AssetGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.AssetGroupTable,
			Columns: []string{asset.AssetGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AssetGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.AssetGroupTable,
			Columns: []string{asset.AssetGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AssetUpdateOne is the builder for updating a single Asset entity.
type AssetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetMutation
}

// SetAssetGroupID sets the "AssetGroupID" field.
func (auo *AssetUpdateOne) SetAssetGroupID(i int) *AssetUpdateOne {
	auo.mutation.SetAssetGroupID(i)
	return auo
}

// SetEmployeeID sets the "EmployeeID" field.
func (auo *AssetUpdateOne) SetEmployeeID(i int) *AssetUpdateOne {
	auo.mutation.SetEmployeeID(i)
	return auo
}

// SetDepartmentLocationID sets the "DepartmentLocationID" field.
func (auo *AssetUpdateOne) SetDepartmentLocationID(i int) *AssetUpdateOne {
	auo.mutation.SetDepartmentLocationID(i)
	return auo
}

// SetAssetSN sets the "AssetSN" field.
func (auo *AssetUpdateOne) SetAssetSN(i int) *AssetUpdateOne {
	auo.mutation.ResetAssetSN()
	auo.mutation.SetAssetSN(i)
	return auo
}

// AddAssetSN adds i to the "AssetSN" field.
func (auo *AssetUpdateOne) AddAssetSN(i int) *AssetUpdateOne {
	auo.mutation.AddAssetSN(i)
	return auo
}

// SetAssetName sets the "AssetName" field.
func (auo *AssetUpdateOne) SetAssetName(s string) *AssetUpdateOne {
	auo.mutation.SetAssetName(s)
	return auo
}

// SetDescription sets the "Description" field.
func (auo *AssetUpdateOne) SetDescription(s string) *AssetUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetWarrantyDate sets the "WarrantyDate" field.
func (auo *AssetUpdateOne) SetWarrantyDate(t time.Time) *AssetUpdateOne {
	auo.mutation.SetWarrantyDate(t)
	return auo
}

// AddAssetPhotoIDs adds the "AssetPhoto" edge to the AssetPhoto entity by IDs.
func (auo *AssetUpdateOne) AddAssetPhotoIDs(ids ...int) *AssetUpdateOne {
	auo.mutation.AddAssetPhotoIDs(ids...)
	return auo
}

// AddAssetPhoto adds the "AssetPhoto" edges to the AssetPhoto entity.
func (auo *AssetUpdateOne) AddAssetPhoto(a ...*AssetPhoto) *AssetUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAssetPhotoIDs(ids...)
}

// AddAssetTransferLogIDs adds the "AssetTransferLog" edge to the AssetTransferLog entity by IDs.
func (auo *AssetUpdateOne) AddAssetTransferLogIDs(ids ...int) *AssetUpdateOne {
	auo.mutation.AddAssetTransferLogIDs(ids...)
	return auo
}

// AddAssetTransferLog adds the "AssetTransferLog" edges to the AssetTransferLog entity.
func (auo *AssetUpdateOne) AddAssetTransferLog(a ...*AssetTransferLog) *AssetUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAssetTransferLogIDs(ids...)
}

// SetDepartmentLocation sets the "DepartmentLocation" edge to the DepartmentLocation entity.
func (auo *AssetUpdateOne) SetDepartmentLocation(d *DepartmentLocation) *AssetUpdateOne {
	return auo.SetDepartmentLocationID(d.ID)
}

// SetEmployee sets the "Employee" edge to the Employee entity.
func (auo *AssetUpdateOne) SetEmployee(e *Employee) *AssetUpdateOne {
	return auo.SetEmployeeID(e.ID)
}

// SetAssetGroup sets the "AssetGroup" edge to the AssetGroup entity.
func (auo *AssetUpdateOne) SetAssetGroup(a *AssetGroup) *AssetUpdateOne {
	return auo.SetAssetGroupID(a.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (auo *AssetUpdateOne) Mutation() *AssetMutation {
	return auo.mutation
}

// ClearAssetPhoto clears all "AssetPhoto" edges to the AssetPhoto entity.
func (auo *AssetUpdateOne) ClearAssetPhoto() *AssetUpdateOne {
	auo.mutation.ClearAssetPhoto()
	return auo
}

// RemoveAssetPhotoIDs removes the "AssetPhoto" edge to AssetPhoto entities by IDs.
func (auo *AssetUpdateOne) RemoveAssetPhotoIDs(ids ...int) *AssetUpdateOne {
	auo.mutation.RemoveAssetPhotoIDs(ids...)
	return auo
}

// RemoveAssetPhoto removes "AssetPhoto" edges to AssetPhoto entities.
func (auo *AssetUpdateOne) RemoveAssetPhoto(a ...*AssetPhoto) *AssetUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAssetPhotoIDs(ids...)
}

// ClearAssetTransferLog clears all "AssetTransferLog" edges to the AssetTransferLog entity.
func (auo *AssetUpdateOne) ClearAssetTransferLog() *AssetUpdateOne {
	auo.mutation.ClearAssetTransferLog()
	return auo
}

// RemoveAssetTransferLogIDs removes the "AssetTransferLog" edge to AssetTransferLog entities by IDs.
func (auo *AssetUpdateOne) RemoveAssetTransferLogIDs(ids ...int) *AssetUpdateOne {
	auo.mutation.RemoveAssetTransferLogIDs(ids...)
	return auo
}

// RemoveAssetTransferLog removes "AssetTransferLog" edges to AssetTransferLog entities.
func (auo *AssetUpdateOne) RemoveAssetTransferLog(a ...*AssetTransferLog) *AssetUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAssetTransferLogIDs(ids...)
}

// ClearDepartmentLocation clears the "DepartmentLocation" edge to the DepartmentLocation entity.
func (auo *AssetUpdateOne) ClearDepartmentLocation() *AssetUpdateOne {
	auo.mutation.ClearDepartmentLocation()
	return auo
}

// ClearEmployee clears the "Employee" edge to the Employee entity.
func (auo *AssetUpdateOne) ClearEmployee() *AssetUpdateOne {
	auo.mutation.ClearEmployee()
	return auo
}

// ClearAssetGroup clears the "AssetGroup" edge to the AssetGroup entity.
func (auo *AssetUpdateOne) ClearAssetGroup() *AssetUpdateOne {
	auo.mutation.ClearAssetGroup()
	return auo
}

// Where appends a list predicates to the AssetUpdate builder.
func (auo *AssetUpdateOne) Where(ps ...predicate.Asset) *AssetUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AssetUpdateOne) Select(field string, fields ...string) *AssetUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Asset entity.
func (auo *AssetUpdateOne) Save(ctx context.Context) (*Asset, error) {
	return withHooks[*Asset, AssetMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AssetUpdateOne) SaveX(ctx context.Context) *Asset {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AssetUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AssetUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AssetUpdateOne) check() error {
	if v, ok := auo.mutation.AssetSN(); ok {
		if err := asset.AssetSNValidator(v); err != nil {
			return &ValidationError{Name: "AssetSN", err: fmt.Errorf(`ent: validator failed for field "Asset.AssetSN": %w`, err)}
		}
	}
	if _, ok := auo.mutation.DepartmentLocationID(); auo.mutation.DepartmentLocationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.DepartmentLocation"`)
	}
	if _, ok := auo.mutation.EmployeeID(); auo.mutation.EmployeeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.Employee"`)
	}
	if _, ok := auo.mutation.AssetGroupID(); auo.mutation.AssetGroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.AssetGroup"`)
	}
	return nil
}

func (auo *AssetUpdateOne) sqlSave(ctx context.Context) (_node *Asset, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(asset.Table, asset.Columns, sqlgraph.NewFieldSpec(asset.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Asset.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asset.FieldID)
		for _, f := range fields {
			if !asset.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != asset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.AssetSN(); ok {
		_spec.SetField(asset.FieldAssetSN, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedAssetSN(); ok {
		_spec.AddField(asset.FieldAssetSN, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AssetName(); ok {
		_spec.SetField(asset.FieldAssetName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(asset.FieldDescription, field.TypeString, value)
	}
	if value, ok := auo.mutation.WarrantyDate(); ok {
		_spec.SetField(asset.FieldWarrantyDate, field.TypeTime, value)
	}
	if auo.mutation.AssetPhotoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetPhotoTable,
			Columns: []string{asset.AssetPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAssetPhotoIDs(); len(nodes) > 0 && !auo.mutation.AssetPhotoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetPhotoTable,
			Columns: []string{asset.AssetPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AssetPhotoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetPhotoTable,
			Columns: []string{asset.AssetPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetphoto.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AssetTransferLogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetTransferLogTable,
			Columns: []string{asset.AssetTransferLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAssetTransferLogIDs(); len(nodes) > 0 && !auo.mutation.AssetTransferLogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetTransferLogTable,
			Columns: []string{asset.AssetTransferLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AssetTransferLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.AssetTransferLogTable,
			Columns: []string{asset.AssetTransferLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DepartmentLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.DepartmentLocationTable,
			Columns: []string{asset.DepartmentLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DepartmentLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.DepartmentLocationTable,
			Columns: []string{asset.DepartmentLocationColumn},
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
	if auo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.EmployeeTable,
			Columns: []string{asset.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.EmployeeTable,
			Columns: []string{asset.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AssetGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.AssetGroupTable,
			Columns: []string{asset.AssetGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AssetGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.AssetGroupTable,
			Columns: []string{asset.AssetGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(assetgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Asset{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
