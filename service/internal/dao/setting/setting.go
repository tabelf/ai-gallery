package setting

import (
	"context"
	"time"

	"ai-gallery/ent"
	entsetting "ai-gallery/ent/setting"
	"ai-gallery/pkg/errgroup"
	"ai-gallery/service/internal/dao"
	"ai-gallery/service/internal/model"
)

type PO struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	SecureID     string `json:"secure_id"`
	SecureKey    string `json:"secure_key"`
	StoreBucket  string `json:"store_bucket"`
	InitPwd      string `json:"init_pwd"`
	OperateID    int    `json:"operate_id"`
	OperateName  string `json:"operate_name"`
}

func GetPO(ctx context.Context) (*PO, error) {
	var (
		storeName    = ""
		storeAddress = ""
		secureID     = ""
		secureKey    = ""
		storeBucket  = ""
		initPwd      = ""
		g            = errgroup.WithContext(ctx)
	)
	g.Go(func(ctx context.Context) (err error) {
		storeName, err = GetByConfigKey(ctx, model.SettingEnums.StoreName.Code)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		storeAddress, err = GetByConfigKey(ctx, model.SettingEnums.StoreAddress.Code)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		secureID, err = GetByConfigKey(ctx, model.SettingEnums.SecureID.Code)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		secureKey, err = GetByConfigKey(ctx, model.SettingEnums.SecureKey.Code)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		storeBucket, err = GetByConfigKey(ctx, model.SettingEnums.StoreBucket.Code)
		return err
	})
	g.Go(func(ctx context.Context) (err error) {
		initPwd, err = GetByConfigKey(ctx, model.SettingEnums.InitPwd.Code)
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return &PO{
		StoreName:    storeName,
		StoreAddress: storeAddress,
		SecureID:     secureID,
		SecureKey:    secureKey,
		StoreBucket:  storeBucket,
		InitPwd:      initPwd,
	}, nil
}

func GetByConfigKey(ctx context.Context, configKey string) (string, error) {
	setting, err := dao.EntClient.Setting.Query().
		Where(entsetting.ConfigKey(configKey)).
		First(ctx)
	if err != nil {
		return "", err
	}
	return setting.ConfigValue, nil
}

func GetInitPwd(ctx context.Context) (string, error) {
	initPwd, err := GetByConfigKey(ctx, model.SettingEnums.InitPwd.Code)
	if err != nil {
		return "", err
	}
	return initPwd, err
}

func Create(ctx context.Context, po *PO) error {
	return dao.WithTx(ctx, dao.EntClient, func(tx *ent.Tx) error {
		if err := create(ctx, tx, model.SettingEnums.StoreName, po.StoreName); err != nil {
			return err
		}
		if err := create(ctx, tx, model.SettingEnums.StoreAddress, po.StoreAddress); err != nil {
			return err
		}
		if err := create(ctx, tx, model.SettingEnums.SecureID, po.SecureID); err != nil {
			return err
		}
		if err := create(ctx, tx, model.SettingEnums.SecureKey, po.SecureKey); err != nil {
			return err
		}
		if err := create(ctx, tx, model.SettingEnums.StoreBucket, po.StoreBucket); err != nil {
			return err
		}
		if err := create(ctx, tx, model.SettingEnums.InitPwd, po.InitPwd); err != nil {
			return err
		}
		return nil
	})
}

func create(ctx context.Context, tx *ent.Tx, enum *model.Enum, value string) error {
	exist, err := tx.Setting.Query().
		Where(entsetting.ConfigKey(enum.Code)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exist {
		if err = tx.Setting.Create().
			SetCreateTime(time.Now()).
			SetUpdateTime(time.Now()).
			SetConfigKey(enum.Code).
			SetConfigValue(value).
			SetMark(enum.Name).
			SetOperateID(-1).
			SetOperateName("system").
			Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func Update(ctx context.Context, po *PO) error {
	return dao.WithTx(ctx, dao.EntClient, func(tx *ent.Tx) error {
		if err := tx.Setting.Update().
			SetConfigValue(po.StoreName).
			SetOperateID(po.OperateID).
			SetOperateName(po.OperateName).
			SetUpdateTime(time.Now()).
			Where(entsetting.ConfigKey(model.SettingEnums.StoreName.Code)).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Setting.Update().
			SetConfigValue(po.StoreAddress).
			SetOperateID(po.OperateID).
			SetOperateName(po.OperateName).
			SetUpdateTime(time.Now()).
			Where(entsetting.ConfigKey(model.SettingEnums.StoreAddress.Code)).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Setting.Update().
			SetConfigValue(po.SecureID).
			SetOperateID(po.OperateID).
			SetOperateName(po.OperateName).
			SetUpdateTime(time.Now()).
			Where(entsetting.ConfigKey(model.SettingEnums.SecureID.Code)).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Setting.Update().
			SetConfigValue(po.SecureKey).
			SetOperateID(po.OperateID).
			SetOperateName(po.OperateName).
			SetUpdateTime(time.Now()).
			Where(entsetting.ConfigKey(model.SettingEnums.SecureKey.Code)).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Setting.Update().
			SetConfigValue(po.StoreBucket).
			SetOperateID(po.OperateID).
			SetOperateName(po.OperateName).
			SetUpdateTime(time.Now()).
			Where(entsetting.ConfigKey(model.SettingEnums.StoreBucket.Code)).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Setting.Update().
			SetConfigValue(po.InitPwd).
			SetOperateID(po.OperateID).
			SetOperateName(po.OperateName).
			SetUpdateTime(time.Now()).
			Where(entsetting.ConfigKey(model.SettingEnums.InitPwd.Code)).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}
