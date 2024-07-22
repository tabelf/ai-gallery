// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ai-gallery/ent/predicate"
	"ai-gallery/ent/setting"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingUpdate is the builder for updating Setting entities.
type SettingUpdate struct {
	config
	hooks    []Hook
	mutation *SettingMutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (su *SettingUpdate) Where(ps ...predicate.Setting) *SettingUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SettingUpdate) SetUpdateTime(t time.Time) *SettingUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetDeleteTime sets the "delete_time" field.
func (su *SettingUpdate) SetDeleteTime(t time.Time) *SettingUpdate {
	su.mutation.SetDeleteTime(t)
	return su
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (su *SettingUpdate) SetNillableDeleteTime(t *time.Time) *SettingUpdate {
	if t != nil {
		su.SetDeleteTime(*t)
	}
	return su
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (su *SettingUpdate) ClearDeleteTime() *SettingUpdate {
	su.mutation.ClearDeleteTime()
	return su
}

// SetConfigKey sets the "config_key" field.
func (su *SettingUpdate) SetConfigKey(s string) *SettingUpdate {
	su.mutation.SetConfigKey(s)
	return su
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (su *SettingUpdate) SetNillableConfigKey(s *string) *SettingUpdate {
	if s != nil {
		su.SetConfigKey(*s)
	}
	return su
}

// ClearConfigKey clears the value of the "config_key" field.
func (su *SettingUpdate) ClearConfigKey() *SettingUpdate {
	su.mutation.ClearConfigKey()
	return su
}

// SetConfigValue sets the "config_value" field.
func (su *SettingUpdate) SetConfigValue(s string) *SettingUpdate {
	su.mutation.SetConfigValue(s)
	return su
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (su *SettingUpdate) SetNillableConfigValue(s *string) *SettingUpdate {
	if s != nil {
		su.SetConfigValue(*s)
	}
	return su
}

// ClearConfigValue clears the value of the "config_value" field.
func (su *SettingUpdate) ClearConfigValue() *SettingUpdate {
	su.mutation.ClearConfigValue()
	return su
}

// SetMark sets the "mark" field.
func (su *SettingUpdate) SetMark(s string) *SettingUpdate {
	su.mutation.SetMark(s)
	return su
}

// SetNillableMark sets the "mark" field if the given value is not nil.
func (su *SettingUpdate) SetNillableMark(s *string) *SettingUpdate {
	if s != nil {
		su.SetMark(*s)
	}
	return su
}

// ClearMark clears the value of the "mark" field.
func (su *SettingUpdate) ClearMark() *SettingUpdate {
	su.mutation.ClearMark()
	return su
}

// SetOperateID sets the "operate_id" field.
func (su *SettingUpdate) SetOperateID(i int) *SettingUpdate {
	su.mutation.ResetOperateID()
	su.mutation.SetOperateID(i)
	return su
}

// SetNillableOperateID sets the "operate_id" field if the given value is not nil.
func (su *SettingUpdate) SetNillableOperateID(i *int) *SettingUpdate {
	if i != nil {
		su.SetOperateID(*i)
	}
	return su
}

// AddOperateID adds i to the "operate_id" field.
func (su *SettingUpdate) AddOperateID(i int) *SettingUpdate {
	su.mutation.AddOperateID(i)
	return su
}

// ClearOperateID clears the value of the "operate_id" field.
func (su *SettingUpdate) ClearOperateID() *SettingUpdate {
	su.mutation.ClearOperateID()
	return su
}

// SetOperateName sets the "operate_name" field.
func (su *SettingUpdate) SetOperateName(s string) *SettingUpdate {
	su.mutation.SetOperateName(s)
	return su
}

// SetNillableOperateName sets the "operate_name" field if the given value is not nil.
func (su *SettingUpdate) SetNillableOperateName(s *string) *SettingUpdate {
	if s != nil {
		su.SetOperateName(*s)
	}
	return su
}

// ClearOperateName clears the value of the "operate_name" field.
func (su *SettingUpdate) ClearOperateName() *SettingUpdate {
	su.mutation.ClearOperateName()
	return su
}

// Mutation returns the SettingMutation object of the builder.
func (su *SettingUpdate) Mutation() *SettingMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SettingUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := setting.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SettingUpdate) check() error {
	if v, ok := su.mutation.ConfigKey(); ok {
		if err := setting.ConfigKeyValidator(v); err != nil {
			return &ValidationError{Name: "config_key", err: fmt.Errorf(`ent: validator failed for field "Setting.config_key": %w`, err)}
		}
	}
	if v, ok := su.mutation.ConfigValue(); ok {
		if err := setting.ConfigValueValidator(v); err != nil {
			return &ValidationError{Name: "config_value", err: fmt.Errorf(`ent: validator failed for field "Setting.config_value": %w`, err)}
		}
	}
	if v, ok := su.mutation.Mark(); ok {
		if err := setting.MarkValidator(v); err != nil {
			return &ValidationError{Name: "mark", err: fmt.Errorf(`ent: validator failed for field "Setting.mark": %w`, err)}
		}
	}
	if v, ok := su.mutation.OperateName(); ok {
		if err := setting.OperateNameValidator(v); err != nil {
			return &ValidationError{Name: "operate_name", err: fmt.Errorf(`ent: validator failed for field "Setting.operate_name": %w`, err)}
		}
	}
	return nil
}

func (su *SettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.DeleteTime(); ok {
		_spec.SetField(setting.FieldDeleteTime, field.TypeTime, value)
	}
	if su.mutation.DeleteTimeCleared() {
		_spec.ClearField(setting.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := su.mutation.ConfigKey(); ok {
		_spec.SetField(setting.FieldConfigKey, field.TypeString, value)
	}
	if su.mutation.ConfigKeyCleared() {
		_spec.ClearField(setting.FieldConfigKey, field.TypeString)
	}
	if value, ok := su.mutation.ConfigValue(); ok {
		_spec.SetField(setting.FieldConfigValue, field.TypeString, value)
	}
	if su.mutation.ConfigValueCleared() {
		_spec.ClearField(setting.FieldConfigValue, field.TypeString)
	}
	if value, ok := su.mutation.Mark(); ok {
		_spec.SetField(setting.FieldMark, field.TypeString, value)
	}
	if su.mutation.MarkCleared() {
		_spec.ClearField(setting.FieldMark, field.TypeString)
	}
	if value, ok := su.mutation.OperateID(); ok {
		_spec.SetField(setting.FieldOperateID, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedOperateID(); ok {
		_spec.AddField(setting.FieldOperateID, field.TypeInt, value)
	}
	if su.mutation.OperateIDCleared() {
		_spec.ClearField(setting.FieldOperateID, field.TypeInt)
	}
	if value, ok := su.mutation.OperateName(); ok {
		_spec.SetField(setting.FieldOperateName, field.TypeString, value)
	}
	if su.mutation.OperateNameCleared() {
		_spec.ClearField(setting.FieldOperateName, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SettingUpdateOne is the builder for updating a single Setting entity.
type SettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SettingMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *SettingUpdateOne) SetUpdateTime(t time.Time) *SettingUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetDeleteTime sets the "delete_time" field.
func (suo *SettingUpdateOne) SetDeleteTime(t time.Time) *SettingUpdateOne {
	suo.mutation.SetDeleteTime(t)
	return suo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableDeleteTime(t *time.Time) *SettingUpdateOne {
	if t != nil {
		suo.SetDeleteTime(*t)
	}
	return suo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (suo *SettingUpdateOne) ClearDeleteTime() *SettingUpdateOne {
	suo.mutation.ClearDeleteTime()
	return suo
}

// SetConfigKey sets the "config_key" field.
func (suo *SettingUpdateOne) SetConfigKey(s string) *SettingUpdateOne {
	suo.mutation.SetConfigKey(s)
	return suo
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableConfigKey(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetConfigKey(*s)
	}
	return suo
}

// ClearConfigKey clears the value of the "config_key" field.
func (suo *SettingUpdateOne) ClearConfigKey() *SettingUpdateOne {
	suo.mutation.ClearConfigKey()
	return suo
}

// SetConfigValue sets the "config_value" field.
func (suo *SettingUpdateOne) SetConfigValue(s string) *SettingUpdateOne {
	suo.mutation.SetConfigValue(s)
	return suo
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableConfigValue(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetConfigValue(*s)
	}
	return suo
}

// ClearConfigValue clears the value of the "config_value" field.
func (suo *SettingUpdateOne) ClearConfigValue() *SettingUpdateOne {
	suo.mutation.ClearConfigValue()
	return suo
}

// SetMark sets the "mark" field.
func (suo *SettingUpdateOne) SetMark(s string) *SettingUpdateOne {
	suo.mutation.SetMark(s)
	return suo
}

// SetNillableMark sets the "mark" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableMark(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetMark(*s)
	}
	return suo
}

// ClearMark clears the value of the "mark" field.
func (suo *SettingUpdateOne) ClearMark() *SettingUpdateOne {
	suo.mutation.ClearMark()
	return suo
}

// SetOperateID sets the "operate_id" field.
func (suo *SettingUpdateOne) SetOperateID(i int) *SettingUpdateOne {
	suo.mutation.ResetOperateID()
	suo.mutation.SetOperateID(i)
	return suo
}

// SetNillableOperateID sets the "operate_id" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableOperateID(i *int) *SettingUpdateOne {
	if i != nil {
		suo.SetOperateID(*i)
	}
	return suo
}

// AddOperateID adds i to the "operate_id" field.
func (suo *SettingUpdateOne) AddOperateID(i int) *SettingUpdateOne {
	suo.mutation.AddOperateID(i)
	return suo
}

// ClearOperateID clears the value of the "operate_id" field.
func (suo *SettingUpdateOne) ClearOperateID() *SettingUpdateOne {
	suo.mutation.ClearOperateID()
	return suo
}

// SetOperateName sets the "operate_name" field.
func (suo *SettingUpdateOne) SetOperateName(s string) *SettingUpdateOne {
	suo.mutation.SetOperateName(s)
	return suo
}

// SetNillableOperateName sets the "operate_name" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableOperateName(s *string) *SettingUpdateOne {
	if s != nil {
		suo.SetOperateName(*s)
	}
	return suo
}

// ClearOperateName clears the value of the "operate_name" field.
func (suo *SettingUpdateOne) ClearOperateName() *SettingUpdateOne {
	suo.mutation.ClearOperateName()
	return suo
}

// Mutation returns the SettingMutation object of the builder.
func (suo *SettingUpdateOne) Mutation() *SettingMutation {
	return suo.mutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (suo *SettingUpdateOne) Where(ps ...predicate.Setting) *SettingUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingUpdateOne) Select(field string, fields ...string) *SettingUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Setting entity.
func (suo *SettingUpdateOne) Save(ctx context.Context) (*Setting, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingUpdateOne) SaveX(ctx context.Context) *Setting {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SettingUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := setting.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SettingUpdateOne) check() error {
	if v, ok := suo.mutation.ConfigKey(); ok {
		if err := setting.ConfigKeyValidator(v); err != nil {
			return &ValidationError{Name: "config_key", err: fmt.Errorf(`ent: validator failed for field "Setting.config_key": %w`, err)}
		}
	}
	if v, ok := suo.mutation.ConfigValue(); ok {
		if err := setting.ConfigValueValidator(v); err != nil {
			return &ValidationError{Name: "config_value", err: fmt.Errorf(`ent: validator failed for field "Setting.config_value": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Mark(); ok {
		if err := setting.MarkValidator(v); err != nil {
			return &ValidationError{Name: "mark", err: fmt.Errorf(`ent: validator failed for field "Setting.mark": %w`, err)}
		}
	}
	if v, ok := suo.mutation.OperateName(); ok {
		if err := setting.OperateNameValidator(v); err != nil {
			return &ValidationError{Name: "operate_name", err: fmt.Errorf(`ent: validator failed for field "Setting.operate_name": %w`, err)}
		}
	}
	return nil
}

func (suo *SettingUpdateOne) sqlSave(ctx context.Context) (_node *Setting, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Setting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, setting.FieldID)
		for _, f := range fields {
			if !setting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != setting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.DeleteTime(); ok {
		_spec.SetField(setting.FieldDeleteTime, field.TypeTime, value)
	}
	if suo.mutation.DeleteTimeCleared() {
		_spec.ClearField(setting.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := suo.mutation.ConfigKey(); ok {
		_spec.SetField(setting.FieldConfigKey, field.TypeString, value)
	}
	if suo.mutation.ConfigKeyCleared() {
		_spec.ClearField(setting.FieldConfigKey, field.TypeString)
	}
	if value, ok := suo.mutation.ConfigValue(); ok {
		_spec.SetField(setting.FieldConfigValue, field.TypeString, value)
	}
	if suo.mutation.ConfigValueCleared() {
		_spec.ClearField(setting.FieldConfigValue, field.TypeString)
	}
	if value, ok := suo.mutation.Mark(); ok {
		_spec.SetField(setting.FieldMark, field.TypeString, value)
	}
	if suo.mutation.MarkCleared() {
		_spec.ClearField(setting.FieldMark, field.TypeString)
	}
	if value, ok := suo.mutation.OperateID(); ok {
		_spec.SetField(setting.FieldOperateID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedOperateID(); ok {
		_spec.AddField(setting.FieldOperateID, field.TypeInt, value)
	}
	if suo.mutation.OperateIDCleared() {
		_spec.ClearField(setting.FieldOperateID, field.TypeInt)
	}
	if value, ok := suo.mutation.OperateName(); ok {
		_spec.SetField(setting.FieldOperateName, field.TypeString, value)
	}
	if suo.mutation.OperateNameCleared() {
		_spec.ClearField(setting.FieldOperateName, field.TypeString)
	}
	_node = &Setting{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
