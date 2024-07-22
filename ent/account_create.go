// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ai-gallery/ent/account"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ac *AccountCreate) SetCreateTime(t time.Time) *AccountCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreateTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *AccountCreate) SetUpdateTime(t time.Time) *AccountCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdateTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetDeleteTime sets the "delete_time" field.
func (ac *AccountCreate) SetDeleteTime(t time.Time) *AccountCreate {
	ac.mutation.SetDeleteTime(t)
	return ac
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDeleteTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetDeleteTime(*t)
	}
	return ac
}

// SetUsername sets the "username" field.
func (ac *AccountCreate) SetUsername(s string) *AccountCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUsername(s *string) *AccountCreate {
	if s != nil {
		ac.SetUsername(*s)
	}
	return ac
}

// SetPassword sets the "password" field.
func (ac *AccountCreate) SetPassword(s string) *AccountCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ac *AccountCreate) SetNillablePassword(s *string) *AccountCreate {
	if s != nil {
		ac.SetPassword(*s)
	}
	return ac
}

// SetNickname sets the "nickname" field.
func (ac *AccountCreate) SetNickname(s string) *AccountCreate {
	ac.mutation.SetNickname(s)
	return ac
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (ac *AccountCreate) SetNillableNickname(s *string) *AccountCreate {
	if s != nil {
		ac.SetNickname(*s)
	}
	return ac
}

// SetRole sets the "role" field.
func (ac *AccountCreate) SetRole(s string) *AccountCreate {
	ac.mutation.SetRole(s)
	return ac
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ac *AccountCreate) SetNillableRole(s *string) *AccountCreate {
	if s != nil {
		ac.SetRole(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *AccountCreate) SetStatus(i int) *AccountCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *AccountCreate) SetNillableStatus(i *int) *AccountCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := account.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := account.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Account.create_time"`)}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Account.update_time"`)}
	}
	if v, ok := ac.mutation.Username(); ok {
		if err := account.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Account.username": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Role(); ok {
		if err := account.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Account.role": %w`, err)}
		}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(account.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.SetField(account.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.DeleteTime(); ok {
		_spec.SetField(account.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.SetField(account.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := ac.mutation.Role(); ok {
		_spec.SetField(account.FieldRole, field.TypeString, value)
		_node.Role = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ac *AccountCreate) OnConflict(opts ...sql.ConflictOption) *AccountUpsertOne {
	ac.conflict = opts
	return &AccountUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AccountCreate) OnConflictColumns(columns ...string) *AccountUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertOne{
		create: ac,
	}
}

type (
	// AccountUpsertOne is the builder for "upsert"-ing
	//  one Account node.
	AccountUpsertOne struct {
		create *AccountCreate
	}

	// AccountUpsert is the "OnConflict" setter.
	AccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *AccountUpsert) SetUpdateTime(v time.Time) *AccountUpsert {
	u.Set(account.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUpdateTime() *AccountUpsert {
	u.SetExcluded(account.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *AccountUpsert) SetDeleteTime(v time.Time) *AccountUpsert {
	u.Set(account.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *AccountUpsert) UpdateDeleteTime() *AccountUpsert {
	u.SetExcluded(account.FieldDeleteTime)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *AccountUpsert) ClearDeleteTime() *AccountUpsert {
	u.SetNull(account.FieldDeleteTime)
	return u
}

// SetUsername sets the "username" field.
func (u *AccountUpsert) SetUsername(v string) *AccountUpsert {
	u.Set(account.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUsername() *AccountUpsert {
	u.SetExcluded(account.FieldUsername)
	return u
}

// ClearUsername clears the value of the "username" field.
func (u *AccountUpsert) ClearUsername() *AccountUpsert {
	u.SetNull(account.FieldUsername)
	return u
}

// SetPassword sets the "password" field.
func (u *AccountUpsert) SetPassword(v string) *AccountUpsert {
	u.Set(account.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AccountUpsert) UpdatePassword() *AccountUpsert {
	u.SetExcluded(account.FieldPassword)
	return u
}

// ClearPassword clears the value of the "password" field.
func (u *AccountUpsert) ClearPassword() *AccountUpsert {
	u.SetNull(account.FieldPassword)
	return u
}

// SetNickname sets the "nickname" field.
func (u *AccountUpsert) SetNickname(v string) *AccountUpsert {
	u.Set(account.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *AccountUpsert) UpdateNickname() *AccountUpsert {
	u.SetExcluded(account.FieldNickname)
	return u
}

// ClearNickname clears the value of the "nickname" field.
func (u *AccountUpsert) ClearNickname() *AccountUpsert {
	u.SetNull(account.FieldNickname)
	return u
}

// SetRole sets the "role" field.
func (u *AccountUpsert) SetRole(v string) *AccountUpsert {
	u.Set(account.FieldRole, v)
	return u
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *AccountUpsert) UpdateRole() *AccountUpsert {
	u.SetExcluded(account.FieldRole)
	return u
}

// ClearRole clears the value of the "role" field.
func (u *AccountUpsert) ClearRole() *AccountUpsert {
	u.SetNull(account.FieldRole)
	return u
}

// SetStatus sets the "status" field.
func (u *AccountUpsert) SetStatus(v int) *AccountUpsert {
	u.Set(account.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AccountUpsert) UpdateStatus() *AccountUpsert {
	u.SetExcluded(account.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *AccountUpsert) AddStatus(v int) *AccountUpsert {
	u.Add(account.FieldStatus, v)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *AccountUpsert) ClearStatus() *AccountUpsert {
	u.SetNull(account.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(account.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AccountUpsertOne) Ignore() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertOne) DoNothing() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreate.OnConflict
// documentation for more info.
func (u *AccountUpsertOne) Update(set func(*AccountUpsert)) *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *AccountUpsertOne) SetUpdateTime(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUpdateTime() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *AccountUpsertOne) SetDeleteTime(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateDeleteTime() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *AccountUpsertOne) ClearDeleteTime() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearDeleteTime()
	})
}

// SetUsername sets the "username" field.
func (u *AccountUpsertOne) SetUsername(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUsername() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *AccountUpsertOne) ClearUsername() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearUsername()
	})
}

// SetPassword sets the "password" field.
func (u *AccountUpsertOne) SetPassword(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdatePassword() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *AccountUpsertOne) ClearPassword() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearPassword()
	})
}

// SetNickname sets the "nickname" field.
func (u *AccountUpsertOne) SetNickname(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateNickname() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *AccountUpsertOne) ClearNickname() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearNickname()
	})
}

// SetRole sets the "role" field.
func (u *AccountUpsertOne) SetRole(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateRole() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateRole()
	})
}

// ClearRole clears the value of the "role" field.
func (u *AccountUpsertOne) ClearRole() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearRole()
	})
}

// SetStatus sets the "status" field.
func (u *AccountUpsertOne) SetStatus(v int) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *AccountUpsertOne) AddStatus(v int) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateStatus() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *AccountUpsertOne) ClearStatus() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *AccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccountUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccountUpsertBulk {
	acb.conflict = opts
	return &AccountUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflictColumns(columns ...string) *AccountUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertBulk{
		create: acb,
	}
}

// AccountUpsertBulk is the builder for "upsert"-ing
// a bulk of Account nodes.
type AccountUpsertBulk struct {
	create *AccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(account.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AccountUpsertBulk) Ignore() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertBulk) DoNothing() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreateBulk.OnConflict
// documentation for more info.
func (u *AccountUpsertBulk) Update(set func(*AccountUpsert)) *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *AccountUpsertBulk) SetUpdateTime(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUpdateTime() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *AccountUpsertBulk) SetDeleteTime(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateDeleteTime() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *AccountUpsertBulk) ClearDeleteTime() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearDeleteTime()
	})
}

// SetUsername sets the "username" field.
func (u *AccountUpsertBulk) SetUsername(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUsername() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *AccountUpsertBulk) ClearUsername() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearUsername()
	})
}

// SetPassword sets the "password" field.
func (u *AccountUpsertBulk) SetPassword(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdatePassword() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *AccountUpsertBulk) ClearPassword() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearPassword()
	})
}

// SetNickname sets the "nickname" field.
func (u *AccountUpsertBulk) SetNickname(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateNickname() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *AccountUpsertBulk) ClearNickname() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearNickname()
	})
}

// SetRole sets the "role" field.
func (u *AccountUpsertBulk) SetRole(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateRole() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateRole()
	})
}

// ClearRole clears the value of the "role" field.
func (u *AccountUpsertBulk) ClearRole() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearRole()
	})
}

// SetStatus sets the "status" field.
func (u *AccountUpsertBulk) SetStatus(v int) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *AccountUpsertBulk) AddStatus(v int) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateStatus() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *AccountUpsertBulk) ClearStatus() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *AccountUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}