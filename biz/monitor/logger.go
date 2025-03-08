package monitor

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

const (
	TraceTag = "log_id_tag"
)

func LogIDMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		hlog.CtxInfof(ctx, "LogIDMiddleware...")
		// 生成唯一的 logid
		logID := uuid.New().String()
		// 将 logid 存储到请求上下文中
		c.Set(TraceTag, logID)
		logContext := context.WithValue(ctx, TraceTag, logID)
		// 继续处理请求
		c.Next(logContext)
	}
}
