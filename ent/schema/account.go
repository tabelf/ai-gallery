package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Account struct {
	BaseSchema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(30).Optional().Comment("用户名"),
		field.String("password").MaxLen(255).Optional().Comment("密码"),
		field.String("nickname").MaxLen(30).Optional().Comment("昵称"),
		field.String("role").MaxLen(20).Optional().Comment("角色: ADMIN, USER"),
		field.Int("status").Optional().Comment("状态: 0 禁用, 1 可用"),
	}
}

func (Account) Edges() []ent.Edge {
	return nil
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "core_account"},
	}
}
