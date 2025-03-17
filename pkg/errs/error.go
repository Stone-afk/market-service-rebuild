package errs

import "errors"

var ErrNoResponse = errors.New("不需要返回 response")
var ErrUnauthorized = errors.New("未授权")
var ErrSessionKeyNotFound = errors.New("session 中没找到对应的 key")
