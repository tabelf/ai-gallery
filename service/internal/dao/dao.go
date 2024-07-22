package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"ai-gallery/ent"
	"ai-gallery/service/internal/config"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	EntClient *ent.Client
)

func NewDB(conf config.EntConf) (*ent.Client, error) {
	db, err := sql.Open(conf.Driver, conf.URL)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	EntClient = ent.NewClient(ent.Driver(entsql.OpenDB(conf.Driver, db)))

	return EntClient, nil
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if r := tx.Rollback(); r != nil {
			err = errors.New(fmt.Sprintf("transaction rollback, err: %+v", r))
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return errors.New(fmt.Sprintf("transaction commint, err: %+v", err))
	}
	return nil
}
