// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ai-gallery/ent/setting"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation *SettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (sc *SettingCreate) SetCreateTime(t time.Time) *SettingCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreateTime(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SettingCreate) SetUpdateTime(t time.Time) *SettingCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdateTime(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetDeleteTime sets the "delete_time" field.
func (sc *SettingCreate) SetDeleteTime(t time.Time) *SettingCreate {
	sc.mutation.SetDeleteTime(t)
	return sc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (sc *SettingCreate) SetNillableDeleteTime(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetDeleteTime(*t)
	}
	return sc
}

// SetConfigKey sets the "config_key" field.
func (sc *SettingCreate) SetConfigKey(s string) *SettingCreate {
	sc.mutation.SetConfigKey(s)
	return sc
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (sc *SettingCreate) SetNillableConfigKey(s *string) *SettingCreate {
	if s != nil {
		sc.SetConfigKey(*s)
	}
	return sc
}

// SetConfigValue sets the "config_value" field.
func (sc *SettingCreate) SetConfigValue(s string) *SettingCreate {
	sc.mutation.SetConfigValue(s)
	return sc
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (sc *SettingCreate) SetNillableConfigValue(s *string) *SettingCreate {
	if s != nil {
		sc.SetConfigValue(*s)
	}
	return sc
}

// SetMark sets the "mark" field.
func (sc *SettingCreate) SetMark(s string) *SettingCreate {
	sc.mutation.SetMark(s)
	return sc
}

// SetNillableMark sets the "mark" field if the given value is not nil.
func (sc *SettingCreate) SetNillableMark(s *string) *SettingCreate {
	if s != nil {
		sc.SetMark(*s)
	}
	return sc
}

// SetOperateID sets the "operate_id" field.
func (sc *SettingCreate) SetOperateID(i int) *SettingCreate {
	sc.mutation.SetOperateID(i)
	return sc
}

// SetNillableOperateID sets the "operate_id" field if the given value is not nil.
func (sc *SettingCreate) SetNillableOperateID(i *int) *SettingCreate {
	if i != nil {
		sc.SetOperateID(*i)
	}
	return sc
}

// SetOperateName sets the "operate_name" field.
func (sc *SettingCreate) SetOperateName(s string) *SettingCreate {
	sc.mutation.SetOperateName(s)
	return sc
}

// SetNillableOperateName sets the "operate_name" field if the given value is not nil.
func (sc *SettingCreate) SetNillableOperateName(s *string) *SettingCreate {
	if s != nil {
		sc.SetOperateName(*s)
	}
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := setting.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := setting.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Setting.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Setting.update_time"`)}
	}
	if v, ok := sc.mutation.ConfigKey(); ok {
		if err := setting.ConfigKeyValidator(v); err != nil {
			return &ValidationError{Name: "config_key", err: fmt.Errorf(`ent: validator failed for field "Setting.config_key": %w`, err)}
		}
	}
	if v, ok := sc.mutation.ConfigValue(); ok {
		if err := setting.ConfigValueValidator(v); err != nil {
			return &ValidationError{Name: "config_value", err: fmt.Errorf(`ent: validator failed for field "Setting.config_value": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Mark(); ok {
		if err := setting.MarkValidator(v); err != nil {
			return &ValidationError{Name: "mark", err: fmt.Errorf(`ent: validator failed for field "Setting.mark": %w`, err)}
		}
	}
	if v, ok := sc.mutation.OperateName(); ok {
		if err := setting.OperateNameValidator(v); err != nil {
			return &ValidationError{Name: "operate_name", err: fmt.Errorf(`ent: validator failed for field "Setting.operate_name": %w`, err)}
		}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(setting.Table, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(setting.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.DeleteTime(); ok {
		_spec.SetField(setting.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := sc.mutation.ConfigKey(); ok {
		_spec.SetField(setting.FieldConfigKey, field.TypeString, value)
		_node.ConfigKey = value
	}
	if value, ok := sc.mutation.ConfigValue(); ok {
		_spec.SetField(setting.FieldConfigValue, field.TypeString, value)
		_node.ConfigValue = value
	}
	if value, ok := sc.mutation.Mark(); ok {
		_spec.SetField(setting.FieldMark, field.TypeString, value)
		_node.Mark = value
	}
	if value, ok := sc.mutation.OperateID(); ok {
		_spec.SetField(setting.FieldOperateID, field.TypeInt, value)
		_node.OperateID = value
	}
	if value, ok := sc.mutation.OperateName(); ok {
		_spec.SetField(setting.FieldOperateName, field.TypeString, value)
		_node.OperateName = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sc *SettingCreate) OnConflict(opts ...sql.ConflictOption) *SettingUpsertOne {
	sc.conflict = opts
	return &SettingUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SettingCreate) OnConflictColumns(columns ...string) *SettingUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertOne{
		create: sc,
	}
}

type (
	// SettingUpsertOne is the builder for "upsert"-ing
	//  one Setting node.
	SettingUpsertOne struct {
		create *SettingCreate
	}

	// SettingUpsert is the "OnConflict" setter.
	SettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *SettingUpsert) SetUpdateTime(v time.Time) *SettingUpsert {
	u.Set(setting.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SettingUpsert) UpdateUpdateTime() *SettingUpsert {
	u.SetExcluded(setting.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *SettingUpsert) SetDeleteTime(v time.Time) *SettingUpsert {
	u.Set(setting.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *SettingUpsert) UpdateDeleteTime() *SettingUpsert {
	u.SetExcluded(setting.FieldDeleteTime)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *SettingUpsert) ClearDeleteTime() *SettingUpsert {
	u.SetNull(setting.FieldDeleteTime)
	return u
}

// SetConfigKey sets the "config_key" field.
func (u *SettingUpsert) SetConfigKey(v string) *SettingUpsert {
	u.Set(setting.FieldConfigKey, v)
	return u
}

// UpdateConfigKey sets the "config_key" field to the value that was provided on create.
func (u *SettingUpsert) UpdateConfigKey() *SettingUpsert {
	u.SetExcluded(setting.FieldConfigKey)
	return u
}

// ClearConfigKey clears the value of the "config_key" field.
func (u *SettingUpsert) ClearConfigKey() *SettingUpsert {
	u.SetNull(setting.FieldConfigKey)
	return u
}

// SetConfigValue sets the "config_value" field.
func (u *SettingUpsert) SetConfigValue(v string) *SettingUpsert {
	u.Set(setting.FieldConfigValue, v)
	return u
}

// UpdateConfigValue sets the "config_value" field to the value that was provided on create.
func (u *SettingUpsert) UpdateConfigValue() *SettingUpsert {
	u.SetExcluded(setting.FieldConfigValue)
	return u
}

// ClearConfigValue clears the value of the "config_value" field.
func (u *SettingUpsert) ClearConfigValue() *SettingUpsert {
	u.SetNull(setting.FieldConfigValue)
	return u
}

// SetMark sets the "mark" field.
func (u *SettingUpsert) SetMark(v string) *SettingUpsert {
	u.Set(setting.FieldMark, v)
	return u
}

// UpdateMark sets the "mark" field to the value that was provided on create.
func (u *SettingUpsert) UpdateMark() *SettingUpsert {
	u.SetExcluded(setting.FieldMark)
	return u
}

// ClearMark clears the value of the "mark" field.
func (u *SettingUpsert) ClearMark() *SettingUpsert {
	u.SetNull(setting.FieldMark)
	return u
}

// SetOperateID sets the "operate_id" field.
func (u *SettingUpsert) SetOperateID(v int) *SettingUpsert {
	u.Set(setting.FieldOperateID, v)
	return u
}

// UpdateOperateID sets the "operate_id" field to the value that was provided on create.
func (u *SettingUpsert) UpdateOperateID() *SettingUpsert {
	u.SetExcluded(setting.FieldOperateID)
	return u
}

// AddOperateID adds v to the "operate_id" field.
func (u *SettingUpsert) AddOperateID(v int) *SettingUpsert {
	u.Add(setting.FieldOperateID, v)
	return u
}

// ClearOperateID clears the value of the "operate_id" field.
func (u *SettingUpsert) ClearOperateID() *SettingUpsert {
	u.SetNull(setting.FieldOperateID)
	return u
}

// SetOperateName sets the "operate_name" field.
func (u *SettingUpsert) SetOperateName(v string) *SettingUpsert {
	u.Set(setting.FieldOperateName, v)
	return u
}

// UpdateOperateName sets the "operate_name" field to the value that was provided on create.
func (u *SettingUpsert) UpdateOperateName() *SettingUpsert {
	u.SetExcluded(setting.FieldOperateName)
	return u
}

// ClearOperateName clears the value of the "operate_name" field.
func (u *SettingUpsert) ClearOperateName() *SettingUpsert {
	u.SetNull(setting.FieldOperateName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingUpsertOne) UpdateNewValues() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(setting.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SettingUpsertOne) Ignore() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertOne) DoNothing() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreate.OnConflict
// documentation for more info.
func (u *SettingUpsertOne) Update(set func(*SettingUpsert)) *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SettingUpsertOne) SetUpdateTime(v time.Time) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateUpdateTime() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *SettingUpsertOne) SetDeleteTime(v time.Time) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateDeleteTime() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *SettingUpsertOne) ClearDeleteTime() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearDeleteTime()
	})
}

// SetConfigKey sets the "config_key" field.
func (u *SettingUpsertOne) SetConfigKey(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetConfigKey(v)
	})
}

// UpdateConfigKey sets the "config_key" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateConfigKey() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateConfigKey()
	})
}

// ClearConfigKey clears the value of the "config_key" field.
func (u *SettingUpsertOne) ClearConfigKey() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearConfigKey()
	})
}

// SetConfigValue sets the "config_value" field.
func (u *SettingUpsertOne) SetConfigValue(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetConfigValue(v)
	})
}

// UpdateConfigValue sets the "config_value" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateConfigValue() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateConfigValue()
	})
}

// ClearConfigValue clears the value of the "config_value" field.
func (u *SettingUpsertOne) ClearConfigValue() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearConfigValue()
	})
}

// SetMark sets the "mark" field.
func (u *SettingUpsertOne) SetMark(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetMark(v)
	})
}

// UpdateMark sets the "mark" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateMark() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateMark()
	})
}

// ClearMark clears the value of the "mark" field.
func (u *SettingUpsertOne) ClearMark() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearMark()
	})
}

// SetOperateID sets the "operate_id" field.
func (u *SettingUpsertOne) SetOperateID(v int) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetOperateID(v)
	})
}

// AddOperateID adds v to the "operate_id" field.
func (u *SettingUpsertOne) AddOperateID(v int) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.AddOperateID(v)
	})
}

// UpdateOperateID sets the "operate_id" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateOperateID() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateOperateID()
	})
}

// ClearOperateID clears the value of the "operate_id" field.
func (u *SettingUpsertOne) ClearOperateID() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearOperateID()
	})
}

// SetOperateName sets the "operate_name" field.
func (u *SettingUpsertOne) SetOperateName(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetOperateName(v)
	})
}

// UpdateOperateName sets the "operate_name" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateOperateName() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateOperateName()
	})
}

// ClearOperateName clears the value of the "operate_name" field.
func (u *SettingUpsertOne) ClearOperateName() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearOperateName()
	})
}

// Exec executes the query.
func (u *SettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	err      error
	builders []*SettingCreate
	conflict []sql.ConflictOption
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingUpsertBulk {
	scb.conflict = opts
	return &SettingUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflictColumns(columns ...string) *SettingUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertBulk{
		create: scb,
	}
}

// SettingUpsertBulk is the builder for "upsert"-ing
// a bulk of Setting nodes.
type SettingUpsertBulk struct {
	create *SettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SettingUpsertBulk) UpdateNewValues() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(setting.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SettingUpsertBulk) Ignore() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertBulk) DoNothing() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreateBulk.OnConflict
// documentation for more info.
func (u *SettingUpsertBulk) Update(set func(*SettingUpsert)) *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SettingUpsertBulk) SetUpdateTime(v time.Time) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateUpdateTime() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *SettingUpsertBulk) SetDeleteTime(v time.Time) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateDeleteTime() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *SettingUpsertBulk) ClearDeleteTime() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearDeleteTime()
	})
}

// SetConfigKey sets the "config_key" field.
func (u *SettingUpsertBulk) SetConfigKey(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetConfigKey(v)
	})
}

// UpdateConfigKey sets the "config_key" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateConfigKey() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateConfigKey()
	})
}

// ClearConfigKey clears the value of the "config_key" field.
func (u *SettingUpsertBulk) ClearConfigKey() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearConfigKey()
	})
}

// SetConfigValue sets the "config_value" field.
func (u *SettingUpsertBulk) SetConfigValue(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetConfigValue(v)
	})
}

// UpdateConfigValue sets the "config_value" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateConfigValue() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateConfigValue()
	})
}

// ClearConfigValue clears the value of the "config_value" field.
func (u *SettingUpsertBulk) ClearConfigValue() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearConfigValue()
	})
}

// SetMark sets the "mark" field.
func (u *SettingUpsertBulk) SetMark(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetMark(v)
	})
}

// UpdateMark sets the "mark" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateMark() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateMark()
	})
}

// ClearMark clears the value of the "mark" field.
func (u *SettingUpsertBulk) ClearMark() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearMark()
	})
}

// SetOperateID sets the "operate_id" field.
func (u *SettingUpsertBulk) SetOperateID(v int) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetOperateID(v)
	})
}

// AddOperateID adds v to the "operate_id" field.
func (u *SettingUpsertBulk) AddOperateID(v int) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.AddOperateID(v)
	})
}

// UpdateOperateID sets the "operate_id" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateOperateID() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateOperateID()
	})
}

// ClearOperateID clears the value of the "operate_id" field.
func (u *SettingUpsertBulk) ClearOperateID() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearOperateID()
	})
}

// SetOperateName sets the "operate_name" field.
func (u *SettingUpsertBulk) SetOperateName(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetOperateName(v)
	})
}

// UpdateOperateName sets the "operate_name" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateOperateName() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateOperateName()
	})
}

// ClearOperateName clears the value of the "operate_name" field.
func (u *SettingUpsertBulk) ClearOperateName() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearOperateName()
	})
}

// Exec executes the query.
func (u *SettingUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
