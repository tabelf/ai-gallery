package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Setting struct {
	BaseSchema
}

func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("config_key").MaxLen(30).Optional().Comment("配置名"),
		field.String("config_value").MaxLen(255).Optional().Comment("配置值"),
		field.String("mark").MaxLen(255).Optional().Comment("说明"),
		field.Int("operate_id").Optional().Comment("操作人ID"),
		field.String("operate_name").MaxLen(50).Optional().Comment("操作人名称"),
	}
}

func (Setting) Edges() []ent.Edge {
	return nil
}

func (Setting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "core_setting"},
	}
}
