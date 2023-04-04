// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/asset"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/employee"
)

// EmployeeCreate is the builder for creating a Employee entity.
type EmployeeCreate struct {
	config
	mutation *EmployeeMutation
	hooks    []Hook
}

// SetFirstName sets the "FirstName" field.
func (ec *EmployeeCreate) SetFirstName(s string) *EmployeeCreate {
	ec.mutation.SetFirstName(s)
	return ec
}

// SetLastName sets the "LastName" field.
func (ec *EmployeeCreate) SetLastName(s string) *EmployeeCreate {
	ec.mutation.SetLastName(s)
	return ec
}

// SetPhone sets the "Phone" field.
func (ec *EmployeeCreate) SetPhone(s string) *EmployeeCreate {
	ec.mutation.SetPhone(s)
	return ec
}

// SetID sets the "id" field.
func (ec *EmployeeCreate) SetID(i int) *EmployeeCreate {
	ec.mutation.SetID(i)
	return ec
}

// AddAssetIDs adds the "Asset" edge to the Asset entity by IDs.
func (ec *EmployeeCreate) AddAssetIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddAssetIDs(ids...)
	return ec
}

// AddAsset adds the "Asset" edges to the Asset entity.
func (ec *EmployeeCreate) AddAsset(a ...*Asset) *EmployeeCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ec.AddAssetIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (ec *EmployeeCreate) Mutation() *EmployeeMutation {
	return ec.mutation
}

// Save creates the Employee in the database.
func (ec *EmployeeCreate) Save(ctx context.Context) (*Employee, error) {
	return withHooks[*Employee, EmployeeMutation](ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeCreate) SaveX(ctx context.Context) *Employee {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EmployeeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EmployeeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EmployeeCreate) check() error {
	if _, ok := ec.mutation.FirstName(); !ok {
		return &ValidationError{Name: "FirstName", err: errors.New(`ent: missing required field "Employee.FirstName"`)}
	}
	if _, ok := ec.mutation.LastName(); !ok {
		return &ValidationError{Name: "LastName", err: errors.New(`ent: missing required field "Employee.LastName"`)}
	}
	if _, ok := ec.mutation.Phone(); !ok {
		return &ValidationError{Name: "Phone", err: errors.New(`ent: missing required field "Employee.Phone"`)}
	}
	return nil
}

func (ec *EmployeeCreate) sqlSave(ctx context.Context) (*Employee, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EmployeeCreate) createSpec() (*Employee, *sqlgraph.CreateSpec) {
	var (
		_node = &Employee{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(employee.Table, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := ec.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := ec.mutation.Phone(); ok {
		_spec.SetField(employee.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if nodes := ec.mutation.AssetIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmployeeCreateBulk is the builder for creating many Employee entities in bulk.
type EmployeeCreateBulk struct {
	config
	builders []*EmployeeCreate
}

// Save creates the Employee entities in the database.
func (ecb *EmployeeCreateBulk) Save(ctx context.Context) ([]*Employee, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Employee, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmployeeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EmployeeCreateBulk) SaveX(ctx context.Context) []*Employee {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EmployeeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EmployeeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
