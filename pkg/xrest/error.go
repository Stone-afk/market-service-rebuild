package xrest

import "market-service/pkg/errs"

var ErrNoResponse = errs.ErrNoResponse
var ErrUnauthorized = errs.ErrUnauthorized

var (
	systemErrorResult = Result{
		Code: errs.SystemError.Code,
		Msg:  errs.SystemError.Msg,
	}
)
