package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TaskDetail struct {
	BaseSchema
}

func (TaskDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("task_id").Optional().Comment("任务ID"),
		field.String("image_url").MaxLen(255).Optional().Comment("图片地址"),
		field.Bool("has_excellent").Optional().Comment("是否为优秀作品"),
	}
}

func (TaskDetail) Edges() []ent.Edge {
	return nil
}

func (TaskDetail) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id"),
	}
}

func (TaskDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "core_task_detail"},
	}
}
