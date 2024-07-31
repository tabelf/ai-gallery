package task

import (
	"ai-gallery/service/internal/dao"
	"context"
	"time"

	"ai-gallery/ent"
	enttask "ai-gallery/ent/task"
	"entgo.io/ent/dialect/sql"
	"github.com/duke-git/lancet/v2/strutil"
)

type PO struct {
	Category     string    `json:"category"`
	Prompt       string    `json:"prompt"`
	AuthorIDs    []int     `json:"author_ids"`
	Size         string    `json:"size"`
	SdModelName  string    `json:"sd_model_name"`
	Status       string    `json:"status"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Sorted       string    `json:"sorted"`
	HasRefImage  *int      `json:"has_ref_image"`
	HasExcellent *int      `json:"has_excellent"`
}

func GetByIDs(ctx context.Context, ids []int) ([]*ent.Task, error) {
	return dao.EntClient.Task.Query().Where(
		enttask.IDIn(ids...),
		enttask.DeleteTimeIsNil(),
	).All(ctx)
}

func GetByID(ctx context.Context, id int) (*ent.Task, error) {
	return dao.EntClient.Task.Query().Where(
		enttask.ID(id),
		enttask.DeleteTimeIsNil(),
	).First(ctx)
}

func GetQueryByPO(po *PO) *ent.TaskQuery {
	taskQuery := dao.EntClient.Task.Query().Where(enttask.DeleteTimeIsNil())
	if strutil.IsNotBlank(po.Category) {
		taskQuery.Where(enttask.Category(po.Category))
	}
	if strutil.IsNotBlank(po.Prompt) {
		taskQuery.Where(func(s *sql.Selector) {
			s.Where(sql.Like(enttask.FieldPrompt, "%"+po.Prompt+"%"))
		})
	}
	if len(po.AuthorIDs) > 0 {
		taskQuery.Where(enttask.AuthorIDIn(po.AuthorIDs...))
	}
	if strutil.IsNotBlank(po.Size) {
		taskQuery.Where(enttask.ImgSize(po.Size))
	}
	if strutil.IsNotBlank(po.SdModelName) {
		taskQuery.Where(enttask.SdModelName(po.SdModelName))
	}
	if po.HasRefImage != nil {
		if *po.HasRefImage == 0 {
			taskQuery.Where(enttask.RefImagesIsNil())
		}
		if *po.HasRefImage == 1 {
			taskQuery.Where(enttask.RefImagesNotNil())
		}
	}
	if strutil.IsNotBlank(po.Status) {
		taskQuery.Where(enttask.Status(po.Status))
	}
	if po.HasExcellent != nil {
		if *po.HasExcellent == 1 {
			taskQuery.Where(enttask.HasExcellent(true))
		}
		if *po.HasExcellent == 0 {
			taskQuery.Where(enttask.HasExcellent(false))
		}
	}
	if !po.StartTime.IsZero() {
		taskQuery.Where(enttask.CreateTimeGTE(po.StartTime))
	}
	if !po.EndTime.IsZero() {
		taskQuery.Where(enttask.CreateTimeLTE(po.EndTime))
	}
	if po.Sorted == "asc" {
		taskQuery.Order(ent.Asc(enttask.FieldCreateTime))
	}
	if po.Sorted == "desc" {
		taskQuery.Order(ent.Desc(enttask.FieldCreateTime))
	}
	return taskQuery
}

func UpdateAuthorInfo(ctx context.Context, authorID int, authorName string) error {
	return dao.EntClient.Task.Update().
		SetAuthorName(authorName).
		Where(enttask.AuthorID(authorID)).
		Exec(ctx)
}
