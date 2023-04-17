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
	"github.com/DanielKirkwood/youchooseserver/ent/predicate"
	"github.com/DanielKirkwood/youchooseserver/ent/user"
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

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetOtp sets the "otp" field.
func (uu *UserUpdate) SetOtp(s string) *UserUpdate {
	uu.mutation.SetOtp(s)
	return uu
}

// SetNillableOtp sets the "otp" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOtp(s *string) *UserUpdate {
	if s != nil {
		uu.SetOtp(*s)
	}
	return uu
}

// ClearOtp clears the value of the "otp" field.
func (uu *UserUpdate) ClearOtp() *UserUpdate {
	uu.mutation.ClearOtp()
	return uu
}

// SetOtpExpiresAt sets the "otp_expires_at" field.
func (uu *UserUpdate) SetOtpExpiresAt(t time.Time) *UserUpdate {
	uu.mutation.SetOtpExpiresAt(t)
	return uu
}

// SetNillableOtpExpiresAt sets the "otp_expires_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOtpExpiresAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetOtpExpiresAt(*t)
	}
	return uu
}

// ClearOtpExpiresAt clears the value of the "otp_expires_at" field.
func (uu *UserUpdate) ClearOtpExpiresAt() *UserUpdate {
	uu.mutation.ClearOtpExpiresAt()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
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

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Otp(); ok {
		_spec.SetField(user.FieldOtp, field.TypeString, value)
	}
	if uu.mutation.OtpCleared() {
		_spec.ClearField(user.FieldOtp, field.TypeString)
	}
	if value, ok := uu.mutation.OtpExpiresAt(); ok {
		_spec.SetField(user.FieldOtpExpiresAt, field.TypeTime, value)
	}
	if uu.mutation.OtpExpiresAtCleared() {
		_spec.ClearField(user.FieldOtpExpiresAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetOtp sets the "otp" field.
func (uuo *UserUpdateOne) SetOtp(s string) *UserUpdateOne {
	uuo.mutation.SetOtp(s)
	return uuo
}

// SetNillableOtp sets the "otp" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOtp(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOtp(*s)
	}
	return uuo
}

// ClearOtp clears the value of the "otp" field.
func (uuo *UserUpdateOne) ClearOtp() *UserUpdateOne {
	uuo.mutation.ClearOtp()
	return uuo
}

// SetOtpExpiresAt sets the "otp_expires_at" field.
func (uuo *UserUpdateOne) SetOtpExpiresAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetOtpExpiresAt(t)
	return uuo
}

// SetNillableOtpExpiresAt sets the "otp_expires_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOtpExpiresAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetOtpExpiresAt(*t)
	}
	return uuo
}

// ClearOtpExpiresAt clears the value of the "otp_expires_at" field.
func (uuo *UserUpdateOne) ClearOtpExpiresAt() *UserUpdateOne {
	uuo.mutation.ClearOtpExpiresAt()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
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

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
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
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Otp(); ok {
		_spec.SetField(user.FieldOtp, field.TypeString, value)
	}
	if uuo.mutation.OtpCleared() {
		_spec.ClearField(user.FieldOtp, field.TypeString)
	}
	if value, ok := uuo.mutation.OtpExpiresAt(); ok {
		_spec.SetField(user.FieldOtpExpiresAt, field.TypeTime, value)
	}
	if uuo.mutation.OtpExpiresAtCleared() {
		_spec.ClearField(user.FieldOtpExpiresAt, field.TypeTime)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
