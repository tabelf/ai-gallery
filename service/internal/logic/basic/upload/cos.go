package basic

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosService struct {
	URL       string `json:"url"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
}

func NewCosService() *CosService {
	return &CosService{}
}

func (c *CosService) SetConfig(url, secretID, secretKey, bucket string) {
	c.URL = url
	c.SecretID = secretID
	c.SecretKey = secretKey
	c.Bucket = bucket
}

func (c *CosService) Put(ctx context.Context, name string, r io.Reader, options ...any) (
	path string, err error,
) {
	u, _ := url.Parse(c.URL)
	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 10 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  c.SecretID,
			SecretKey: c.SecretKey,
		},
	})
	if _, err = client.Object.Put(ctx, name, r, nil); err != nil {
		return "", err
	}
	return c.URL + "/" + name, nil
}
