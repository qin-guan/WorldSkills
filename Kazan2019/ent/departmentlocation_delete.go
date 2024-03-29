// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/departmentlocation"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// DepartmentLocationDelete is the builder for deleting a DepartmentLocation entity.
type DepartmentLocationDelete struct {
	config
	hooks    []Hook
	mutation *DepartmentLocationMutation
}

// Where appends a list predicates to the DepartmentLocationDelete builder.
func (dld *DepartmentLocationDelete) Where(ps ...predicate.DepartmentLocation) *DepartmentLocationDelete {
	dld.mutation.Where(ps...)
	return dld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dld *DepartmentLocationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DepartmentLocationMutation](ctx, dld.sqlExec, dld.mutation, dld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dld *DepartmentLocationDelete) ExecX(ctx context.Context) int {
	n, err := dld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dld *DepartmentLocationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(departmentlocation.Table, sqlgraph.NewFieldSpec(departmentlocation.FieldID, field.TypeInt))
	if ps := dld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dld.mutation.done = true
	return affected, err
}

// DepartmentLocationDeleteOne is the builder for deleting a single DepartmentLocation entity.
type DepartmentLocationDeleteOne struct {
	dld *DepartmentLocationDelete
}

// Where appends a list predicates to the DepartmentLocationDelete builder.
func (dldo *DepartmentLocationDeleteOne) Where(ps ...predicate.DepartmentLocation) *DepartmentLocationDeleteOne {
	dldo.dld.mutation.Where(ps...)
	return dldo
}

// Exec executes the deletion query.
func (dldo *DepartmentLocationDeleteOne) Exec(ctx context.Context) error {
	n, err := dldo.dld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{departmentlocation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dldo *DepartmentLocationDeleteOne) ExecX(ctx context.Context) {
	if err := dldo.Exec(ctx); err != nil {
		panic(err)
	}
}
