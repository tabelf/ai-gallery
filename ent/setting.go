// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ai-gallery/ent/setting"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Setting is the model entity for the Setting schema.
type Setting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 配置名
	ConfigKey string `json:"config_key,omitempty"`
	// 配置值
	ConfigValue string `json:"config_value,omitempty"`
	// 说明
	Mark string `json:"mark,omitempty"`
	// 操作人ID
	OperateID int `json:"operate_id,omitempty"`
	// 操作人名称
	OperateName  string `json:"operate_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Setting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case setting.FieldID, setting.FieldOperateID:
			values[i] = new(sql.NullInt64)
		case setting.FieldConfigKey, setting.FieldConfigValue, setting.FieldMark, setting.FieldOperateName:
			values[i] = new(sql.NullString)
		case setting.FieldCreateTime, setting.FieldUpdateTime, setting.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Setting fields.
func (s *Setting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case setting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case setting.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case setting.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case setting.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				s.DeleteTime = new(time.Time)
				*s.DeleteTime = value.Time
			}
		case setting.FieldConfigKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config_key", values[i])
			} else if value.Valid {
				s.ConfigKey = value.String
			}
		case setting.FieldConfigValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config_value", values[i])
			} else if value.Valid {
				s.ConfigValue = value.String
			}
		case setting.FieldMark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mark", values[i])
			} else if value.Valid {
				s.Mark = value.String
			}
		case setting.FieldOperateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field operate_id", values[i])
			} else if value.Valid {
				s.OperateID = int(value.Int64)
			}
		case setting.FieldOperateName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operate_name", values[i])
			} else if value.Valid {
				s.OperateName = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Setting.
// This includes values selected through modifiers, order, etc.
func (s *Setting) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Setting.
// Note that you need to call Setting.Unwrap() before calling this method if this Setting
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Setting) Update() *SettingUpdateOne {
	return NewSettingClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Setting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Setting) Unwrap() *Setting {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Setting is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Setting) String() string {
	var builder strings.Builder
	builder.WriteString("Setting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := s.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("config_key=")
	builder.WriteString(s.ConfigKey)
	builder.WriteString(", ")
	builder.WriteString("config_value=")
	builder.WriteString(s.ConfigValue)
	builder.WriteString(", ")
	builder.WriteString("mark=")
	builder.WriteString(s.Mark)
	builder.WriteString(", ")
	builder.WriteString("operate_id=")
	builder.WriteString(fmt.Sprintf("%v", s.OperateID))
	builder.WriteString(", ")
	builder.WriteString("operate_name=")
	builder.WriteString(s.OperateName)
	builder.WriteByte(')')
	return builder.String()
}

// Settings is a parsable slice of Setting.
type Settings []*Setting