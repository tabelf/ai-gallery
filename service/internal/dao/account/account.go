package account

import (
	"context"
	"time"

	"ai-gallery/ent"
	entaccount "ai-gallery/ent/account"
	"ai-gallery/service/internal/dao"

	"github.com/duke-git/lancet/v2/strutil"
)

type PO struct {
	ID         int        `json:"id"`
	Username   string     `json:"username"`
	Nickname   string     `json:"nickname"`
	Password   string     `json:"password"`
	Role       string     `json:"role"`
	Status     int        `json:"status"`
	DeleteTime *time.Time `json:"delete_time"`
}

func GetByID(ctx context.Context, id int) (*ent.Account, error) {
	return dao.EntClient.Account.Query().Where(
		entaccount.ID(id),
		entaccount.DeleteTimeIsNil(),
	).First(ctx)
}

func GetByUsername(ctx context.Context, username string) (*ent.Account, error) {
	return dao.EntClient.Account.Query().Where(
		entaccount.Username(username),
		entaccount.DeleteTimeIsNil(),
	).First(ctx)
}

func Create(ctx context.Context, po *PO) error {
	return dao.EntClient.Account.Create().
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		SetUsername(po.Username).
		SetNickname(po.Nickname).
		SetPassword(po.Password).
		SetStatus(po.Status).
		SetRole(po.Role).
		Exec(ctx)
}

func Update(ctx context.Context, po *PO) error {
	update := dao.EntClient.Account.Update()
	if strutil.IsNotBlank(po.Username) {
		update.SetUsername(po.Username)
	}
	if strutil.IsNotBlank(po.Nickname) {
		update.SetNickname(po.Nickname)
	}
	if strutil.IsNotBlank(po.Role) {
		update.SetRole(po.Role)
	}
	if strutil.IsNotBlank(po.Password) {
		update.SetPassword(po.Password)
	}
	if po.Status != -1 {
		update.SetStatus(po.Status)
	}
	if po.DeleteTime != nil {
		update.SetNillableDeleteTime(po.DeleteTime)
	}
	return update.
		Where(entaccount.ID(po.ID), entaccount.DeleteTimeIsNil()).
		Exec(ctx)
}
