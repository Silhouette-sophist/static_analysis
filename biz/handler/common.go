package handler

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.sophist.static_analysis/biz/model/base"
)

func SuccessBase() *base.Base {
	return &base.Base{
		Code:   consts.StatusOK,
		Messga: consts.StatusMessage(consts.StatusOK),
	}
}
