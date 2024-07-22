package basic

import (
	"context"
	"io"
)

type LocalService struct {
	URL       string `json:"url"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
}

func NewLocalService() *LocalService {
	return &LocalService{}
}

func (c *LocalService) SetConfig(url, secretID, secretKey, bucket string) {
	c.URL = url
	c.SecretID = secretID
	c.SecretKey = secretKey
	c.Bucket = bucket
}

func (c *LocalService) Put(ctx context.Context, name string, r io.Reader, options ...any) (
	path string, err error,
) {
	return c.URL + "/" + name, nil
}
