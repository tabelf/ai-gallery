package basic

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssService struct {
	URL       string `json:"url"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
}

func NewOssService() *OssService {
	return &OssService{}
}

func (c *OssService) SetConfig(url, secretID, secretKey, bucket string) {
	c.URL = url
	c.SecretID = secretID
	c.SecretKey = secretKey
	c.Bucket = bucket
}

func (c *OssService) Put(ctx context.Context, name string, r io.Reader, options ...any) (
	path string, err error,
) {
	client, err := oss.New(c.URL, c.SecretID, c.SecretKey)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(c.Bucket)
	if err != nil {
		return "", err
	}
	if err = bucket.PutObject(name, r); err != nil {
		return "", err
	}
	region := strings.TrimPrefix(strings.TrimPrefix(c.URL, "http://"), "https://")
	// 构造完整的URL
	url := fmt.Sprintf("https://%s.%s/%s", c.Bucket, region, name)
	return url, err
}
