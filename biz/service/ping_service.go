package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.sophist.static_analysis/biz/model/index"
	"github.sophist.static_analysis/biz/monitor"
)

func Ping(ctx context.Context, req index.PingReq) (map[string]string, error) {
	hlog.CtxInfof(ctx, "logId %v Ping...", ctx.Value(monitor.TraceTag))
	return map[string]string{"ping": "pong"}, nil
}
