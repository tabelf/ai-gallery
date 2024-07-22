package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type BaseSchema struct {
	ent.Schema
}

func (BaseSchema) Mixin() []ent.Mixin {
	return []ent.Mixin{BaseMixin{}}
}

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	datetime := map[string]string{dialect.MySQL: "datetime"}
	return []ent.Field{
		field.Time("create_time").Default(time.Now).Immutable().SchemaType(datetime).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).SchemaType(datetime).UpdateDefault(time.Now).Comment("更新时间"),
		field.Time("delete_time").SchemaType(datetime).Optional().Nillable().Comment("删除时间"),
	}
}
