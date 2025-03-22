package xrest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// BS 的意思是，传入的业务逻辑方法可以接受 req 和 sess 两个参数
func BS[Req, Resp any](fn func(ctx *gin.Context, req Req) (Resp, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		// Bind 方法本身会返回 400 的错误
		if err := ctx.Bind(&req); err != nil {
			slog.Debug("绑定参数失败", slog.Any("err", err))
			return
		}
		var res Resp
		res, err := fn(ctx, req)

		if errors.Is(err, ErrNoResponse) {
			slog.Debug("不需要响应", slog.Any("err", err))
			return
		}
		if errors.Is(err, ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.PureJSON(http.StatusInternalServerError, systemErrorResult)
			return
		}
		ctx.PureJSON(http.StatusOK, Result{Data: res})
	}
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
