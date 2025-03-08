// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.sophist.static_analysis/biz/monitor"
)

func main() {
	h := server.Default(server.WithTracer(&monitor.ReqTracer))

	register(h)
	h.Spin()
}
