// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ai-gallery/ent/account"
	"ai-gallery/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdateTime sets the "update_time" field.
func (au *AccountUpdate) SetUpdateTime(t time.Time) *AccountUpdate {
	au.mutation.SetUpdateTime(t)
	return au
}

// SetDeleteTime sets the "delete_time" field.
func (au *AccountUpdate) SetDeleteTime(t time.Time) *AccountUpdate {
	au.mutation.SetDeleteTime(t)
	return au
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeleteTime(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetDeleteTime(*t)
	}
	return au
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (au *AccountUpdate) ClearDeleteTime() *AccountUpdate {
	au.mutation.ClearDeleteTime()
	return au
}

// SetUsername sets the "username" field.
func (au *AccountUpdate) SetUsername(s string) *AccountUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUsername(s *string) *AccountUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// ClearUsername clears the value of the "username" field.
func (au *AccountUpdate) ClearUsername() *AccountUpdate {
	au.mutation.ClearUsername()
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePassword(s *string) *AccountUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// ClearPassword clears the value of the "password" field.
func (au *AccountUpdate) ClearPassword() *AccountUpdate {
	au.mutation.ClearPassword()
	return au
}

// SetNickname sets the "nickname" field.
func (au *AccountUpdate) SetNickname(s string) *AccountUpdate {
	au.mutation.SetNickname(s)
	return au
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (au *AccountUpdate) SetNillableNickname(s *string) *AccountUpdate {
	if s != nil {
		au.SetNickname(*s)
	}
	return au
}

// ClearNickname clears the value of the "nickname" field.
func (au *AccountUpdate) ClearNickname() *AccountUpdate {
	au.mutation.ClearNickname()
	return au
}

// SetRole sets the "role" field.
func (au *AccountUpdate) SetRole(s string) *AccountUpdate {
	au.mutation.SetRole(s)
	return au
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (au *AccountUpdate) SetNillableRole(s *string) *AccountUpdate {
	if s != nil {
		au.SetRole(*s)
	}
	return au
}

// ClearRole clears the value of the "role" field.
func (au *AccountUpdate) ClearRole() *AccountUpdate {
	au.mutation.ClearRole()
	return au
}

// SetStatus sets the "status" field.
func (au *AccountUpdate) SetStatus(i int) *AccountUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(i)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *AccountUpdate) SetNillableStatus(i *int) *AccountUpdate {
	if i != nil {
		au.SetStatus(*i)
	}
	return au
}

// AddStatus adds i to the "status" field.
func (au *AccountUpdate) AddStatus(i int) *AccountUpdate {
	au.mutation.AddStatus(i)
	return au
}

// ClearStatus clears the value of the "status" field.
func (au *AccountUpdate) ClearStatus() *AccountUpdate {
	au.mutation.ClearStatus()
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdateTime(); !ok {
		v := account.UpdateDefaultUpdateTime()
		au.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Username(); ok {
		if err := account.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Account.username": %w`, err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if v, ok := au.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if v, ok := au.mutation.Role(); ok {
		if err := account.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Account.role": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdateTime(); ok {
		_spec.SetField(account.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := au.mutation.DeleteTime(); ok {
		_spec.SetField(account.FieldDeleteTime, field.TypeTime, value)
	}
	if au.mutation.DeleteTimeCleared() {
		_spec.ClearField(account.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(account.FieldUsername, field.TypeString, value)
	}
	if au.mutation.UsernameCleared() {
		_spec.ClearField(account.FieldUsername, field.TypeString)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if au.mutation.PasswordCleared() {
		_spec.ClearField(account.FieldPassword, field.TypeString)
	}
	if value, ok := au.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
	}
	if au.mutation.NicknameCleared() {
		_spec.ClearField(account.FieldNickname, field.TypeString)
	}
	if value, ok := au.mutation.Role(); ok {
		_spec.SetField(account.FieldRole, field.TypeString, value)
	}
	if au.mutation.RoleCleared() {
		_spec.ClearField(account.FieldRole, field.TypeString)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.AddField(account.FieldStatus, field.TypeInt, value)
	}
	if au.mutation.StatusCleared() {
		_spec.ClearField(account.FieldStatus, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdateTime sets the "update_time" field.
func (auo *AccountUpdateOne) SetUpdateTime(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdateTime(t)
	return auo
}

// SetDeleteTime sets the "delete_time" field.
func (auo *AccountUpdateOne) SetDeleteTime(t time.Time) *AccountUpdateOne {
	auo.mutation.SetDeleteTime(t)
	return auo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeleteTime(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetDeleteTime(*t)
	}
	return auo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (auo *AccountUpdateOne) ClearDeleteTime() *AccountUpdateOne {
	auo.mutation.ClearDeleteTime()
	return auo
}

// SetUsername sets the "username" field.
func (auo *AccountUpdateOne) SetUsername(s string) *AccountUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUsername(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// ClearUsername clears the value of the "username" field.
func (auo *AccountUpdateOne) ClearUsername() *AccountUpdateOne {
	auo.mutation.ClearUsername()
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePassword(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// ClearPassword clears the value of the "password" field.
func (auo *AccountUpdateOne) ClearPassword() *AccountUpdateOne {
	auo.mutation.ClearPassword()
	return auo
}

// SetNickname sets the "nickname" field.
func (auo *AccountUpdateOne) SetNickname(s string) *AccountUpdateOne {
	auo.mutation.SetNickname(s)
	return auo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableNickname(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetNickname(*s)
	}
	return auo
}

// ClearNickname clears the value of the "nickname" field.
func (auo *AccountUpdateOne) ClearNickname() *AccountUpdateOne {
	auo.mutation.ClearNickname()
	return auo
}

// SetRole sets the "role" field.
func (auo *AccountUpdateOne) SetRole(s string) *AccountUpdateOne {
	auo.mutation.SetRole(s)
	return auo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableRole(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetRole(*s)
	}
	return auo
}

// ClearRole clears the value of the "role" field.
func (auo *AccountUpdateOne) ClearRole() *AccountUpdateOne {
	auo.mutation.ClearRole()
	return auo
}

// SetStatus sets the "status" field.
func (auo *AccountUpdateOne) SetStatus(i int) *AccountUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(i)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableStatus(i *int) *AccountUpdateOne {
	if i != nil {
		auo.SetStatus(*i)
	}
	return auo
}

// AddStatus adds i to the "status" field.
func (auo *AccountUpdateOne) AddStatus(i int) *AccountUpdateOne {
	auo.mutation.AddStatus(i)
	return auo
}

// ClearStatus clears the value of the "status" field.
func (auo *AccountUpdateOne) ClearStatus() *AccountUpdateOne {
	auo.mutation.ClearStatus()
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdateTime(); !ok {
		v := account.UpdateDefaultUpdateTime()
		auo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Username(); ok {
		if err := account.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Account.username": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Role(); ok {
		if err := account.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Account.role": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
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
	if value, ok := auo.mutation.UpdateTime(); ok {
		_spec.SetField(account.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.DeleteTime(); ok {
		_spec.SetField(account.FieldDeleteTime, field.TypeTime, value)
	}
	if auo.mutation.DeleteTimeCleared() {
		_spec.ClearField(account.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(account.FieldUsername, field.TypeString, value)
	}
	if auo.mutation.UsernameCleared() {
		_spec.ClearField(account.FieldUsername, field.TypeString)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if auo.mutation.PasswordCleared() {
		_spec.ClearField(account.FieldPassword, field.TypeString)
	}
	if value, ok := auo.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
	}
	if auo.mutation.NicknameCleared() {
		_spec.ClearField(account.FieldNickname, field.TypeString)
	}
	if value, ok := auo.mutation.Role(); ok {
		_spec.SetField(account.FieldRole, field.TypeString, value)
	}
	if auo.mutation.RoleCleared() {
		_spec.ClearField(account.FieldRole, field.TypeString)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.AddField(account.FieldStatus, field.TypeInt, value)
	}
	if auo.mutation.StatusCleared() {
		_spec.ClearField(account.FieldStatus, field.TypeInt)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
