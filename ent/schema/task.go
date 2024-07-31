package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Task struct {
	BaseSchema
}

func (Task) Fields() []ent.Field {
	datetime := map[string]string{dialect.MySQL: "datetime"}
	return []ent.Field{
		field.Text("prompt").Optional().Comment("提示词"),
		field.Text("negative_prompt").Optional().Comment("反向提示词"),
		field.String("category").MaxLen(20).Optional().Comment("类型: txt2img, img2img"),
		field.Float32("weight").Optional().Comment("长度"),
		field.Float32("height").Optional().Comment("宽度"),
		field.String("img_size").MaxLen(20).Optional().Comment("图片比例"),
		field.String("seed").MaxLen(20).Optional().Comment("随机数"),
		field.String("sampler_name").MaxLen(50).Optional().Comment("采样器"),
		field.Int("steps").Optional().Comment("采样步数"),
		field.Int("cfg_scale").Optional().Comment("引导系数"),
		field.Int("batch_size").Optional().Comment("批次"),
		field.Int("total").Optional().Comment("总量"),
		field.String("sd_model_name").MaxLen(100).Optional().Comment("模型名称"),
		field.String("sd_model_hash").MaxLen(50).Optional().Comment("模型hash值"),
		field.String("sd_vae_name").MaxLen(100).Optional().Comment("VAE名称"),
		field.String("sd_vae_hash").MaxLen(50).Optional().Comment("VAE hash值"),
		field.Time("job_timestamp").SchemaType(datetime).Optional().Comment("任务时间"),
		field.String("version").MaxLen(20).Optional().Comment("SD 版本号"),
		field.String("grid_image_url").MaxLen(255).Optional().Comment("网格图片地址"),
		field.String("status").MaxLen(20).Optional().Comment("状态：INIT, PROCESS, COMPLETE, EXCEPTION"),
		field.Int("author_id").Optional().Comment("作者ID"),
		field.String("author_name").MaxLen(30).Optional().Comment("作者名称"),
		field.JSON("ref_images", []string{}).Optional().Comment("参考图片"),
		field.String("store").Optional().Comment("服务存储名称"),
		field.Int("count").Optional().Comment("实际作品数量"),
		field.Bool("has_excellent").Optional().Comment("是否存在优秀作品"),
		field.Text("extra").Optional().Comment("扩展信息"),
	}
}

func (Task) Edges() []ent.Edge {
	return nil
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("author_id"),
	}
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "core_task"},
	}
}
