package monitor

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
)

/**
// Tracer is executed at the start and finish of an HTTP.
type Tracer interface {
	Start(ctx context.Context, c *app.RequestContext) context.Context
	Finish(ctx context.Context, c *app.RequestContext)
}
*/

var ReqTracer = RequestTracer{}

type RequestTracer struct {
}

func (receiver *RequestTracer) Start(ctx context.Context, c *app.RequestContext) context.Context {
	// do nothing
	return ctx
}

// Finish https://www.cloudwego.io/zh/docs/hertz/tutorials/framework-exten/monitor/
func (receiver *RequestTracer) Finish(ctx context.Context, c *app.RequestContext) {
	ti := c.GetTraceInfo()
	rpcStart := ti.Stats().GetEvent(stats.HTTPStart)
	rpcFinish := ti.Stats().GetEvent(stats.HTTPFinish)
	cost := rpcFinish.Time().Sub(rpcStart.Time()).Milliseconds()
	path := c.Path()
	hlog.CtxInfof(ctx, "request traceId %v path %s finished cost %d ms", c.Value(TraceTag), path, cost)
}
