package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrParamValidate = Gerror(100, "请求参数验证错误")
)

func Gerror(code int, message string) error {
	return status.Error(codes.Code(code), message)
}
