// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/assettransferlog"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// AssetTransferLogDelete is the builder for deleting a AssetTransferLog entity.
type AssetTransferLogDelete struct {
	config
	hooks    []Hook
	mutation *AssetTransferLogMutation
}

// Where appends a list predicates to the AssetTransferLogDelete builder.
func (atld *AssetTransferLogDelete) Where(ps ...predicate.AssetTransferLog) *AssetTransferLogDelete {
	atld.mutation.Where(ps...)
	return atld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atld *AssetTransferLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AssetTransferLogMutation](ctx, atld.sqlExec, atld.mutation, atld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (atld *AssetTransferLogDelete) ExecX(ctx context.Context) int {
	n, err := atld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atld *AssetTransferLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(assettransferlog.Table, sqlgraph.NewFieldSpec(assettransferlog.FieldID, field.TypeInt))
	if ps := atld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	atld.mutation.done = true
	return affected, err
}

// AssetTransferLogDeleteOne is the builder for deleting a single AssetTransferLog entity.
type AssetTransferLogDeleteOne struct {
	atld *AssetTransferLogDelete
}

// Where appends a list predicates to the AssetTransferLogDelete builder.
func (atldo *AssetTransferLogDeleteOne) Where(ps ...predicate.AssetTransferLog) *AssetTransferLogDeleteOne {
	atldo.atld.mutation.Where(ps...)
	return atldo
}

// Exec executes the deletion query.
func (atldo *AssetTransferLogDeleteOne) Exec(ctx context.Context) error {
	n, err := atldo.atld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{assettransferlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atldo *AssetTransferLogDeleteOne) ExecX(ctx context.Context) {
	if err := atldo.Exec(ctx); err != nil {
		panic(err)
	}
}