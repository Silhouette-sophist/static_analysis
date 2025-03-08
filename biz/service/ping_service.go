package service

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.sophist.static_analysis/biz/model/index"
)

func Ping(ctx context.Context, req index.PingReq) (map[string]string, error) {
	logger.CtxInfof(ctx, "Ping...")
	return map[string]string{"ping": "pong"}, nil
}
