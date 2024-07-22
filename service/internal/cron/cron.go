package cron

import (
	"context"
	"time"

	"ai-gallery/service/internal/config"
	"ai-gallery/service/internal/cron/jobs"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logc"
)

type JobFunc func(ctx context.Context, config config.Config)

type Job struct {
	Name string  `json:"name"`
	Spec string  `json:"spec"`
	Func JobFunc `json:"func"`
}

var cronJobs = []Job{
	// 每 10 分钟执行一次
	{Spec: "@every 10m", Name: "CheckTaskStatus", Func: jobs.CheckTaskStatus},
}

func CreateCronJob(config config.Config) (*cron.Cron, error) {
	c := cron.New(cron.WithChain(cron.Recover(NewLogger())))
	for _, j := range cronJobs {
		_, err := c.AddFunc(j.Spec, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()
			j.Func(ctx, config)
		})
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Logger 自定义Cron.Logger.
type Logger struct {
	ctx context.Context
}

func NewLogger() Logger {
	return Logger{ctx: context.Background()}
}

// Info level为debug.
func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	logc.Infof(l.ctx, "%s %+v", msg, keysAndValues)
}

// Error level为error.
func (l Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	logc.Errorf(l.ctx, "%w %s %+v", err, msg, keysAndValues)
}
