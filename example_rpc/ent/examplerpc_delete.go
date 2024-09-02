// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sanagiig/fox-admin-core/example_rpc/ent/examplerpc"
	"github.com/Sanagiig/fox-admin-core/example_rpc/ent/predicate"
)

// ExampleRpcDelete is the builder for deleting a ExampleRpc entity.
type ExampleRpcDelete struct {
	config
	hooks    []Hook
	mutation *ExampleRpcMutation
}

// Where appends a list predicates to the ExampleRpcDelete builder.
func (erd *ExampleRpcDelete) Where(ps ...predicate.ExampleRpc) *ExampleRpcDelete {
	erd.mutation.Where(ps...)
	return erd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (erd *ExampleRpcDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, erd.sqlExec, erd.mutation, erd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (erd *ExampleRpcDelete) ExecX(ctx context.Context) int {
	n, err := erd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (erd *ExampleRpcDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(examplerpc.Table, sqlgraph.NewFieldSpec(examplerpc.FieldID, field.TypeInt))
	if ps := erd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, erd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	erd.mutation.done = true
	return affected, err
}

// ExampleRpcDeleteOne is the builder for deleting a single ExampleRpc entity.
type ExampleRpcDeleteOne struct {
	erd *ExampleRpcDelete
}

// Where appends a list predicates to the ExampleRpcDelete builder.
func (erdo *ExampleRpcDeleteOne) Where(ps ...predicate.ExampleRpc) *ExampleRpcDeleteOne {
	erdo.erd.mutation.Where(ps...)
	return erdo
}

// Exec executes the deletion query.
func (erdo *ExampleRpcDeleteOne) Exec(ctx context.Context) error {
	n, err := erdo.erd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{examplerpc.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (erdo *ExampleRpcDeleteOne) ExecX(ctx context.Context) {
	if err := erdo.Exec(ctx); err != nil {
		panic(err)
	}
}
