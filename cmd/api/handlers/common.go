package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
)

func handleError(err error, message string, c *app.RequestContext) {
	hlog.Info(message, err)
	consts.SendResponse(c, struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
	}{consts.ErrCode, message})
}
