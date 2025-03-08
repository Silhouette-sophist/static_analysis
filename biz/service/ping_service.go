package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.sophist.static_analysis/biz/model/index"
)

func Ping(ctx context.Context, req index.PingReq) (map[string]string, error) {
	hlog.CtxInfof(ctx, "Ping...")
	return map[string]string{"ping": "pong"}, nil
}
