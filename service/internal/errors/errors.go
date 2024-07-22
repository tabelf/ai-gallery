package errors

import (
	"context"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

var (
	ErrServiceInternal = func(ctx context.Context) *ErrorMessage {
		return NewError(10000, "服务内部错误")
	}
)

var (
	ErrTaskImageBase64Parse = func(ctx context.Context) *ErrorMessage {
		return NewError(10100, "图片base64解析错误")
	}
	ErrUploadServiceNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10101, "上传服务不存在")
	}
)

var (
	ErrAccountNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10200, "账号不存在")
	}
	ErrAccountDisabled = func(ctx context.Context) *ErrorMessage {
		return NewError(10201, "账号不可用")
	}
	ErrAccountPwd = func(ctx context.Context) *ErrorMessage {
		return NewError(10202, "密码错误")
	}
	ErrAccountTokenGen = func(ctx context.Context) *ErrorMessage {
		return NewError(10203, "无效 token 数据")
	}
	ErrAccountPermission = func(ctx context.Context) *ErrorMessage {
		return NewError(10204, "账号无权限")
	}
	ErrAccountTokenNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10205, "token不存在")
	}
	ErrAccountUsernameExist = func(ctx context.Context) *ErrorMessage {
		return NewError(10206, "用户名已存在")
	}
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorMessage) Error() string {
	return e.Message
}

func NewError(code int, message string) *ErrorMessage {
	return &ErrorMessage{
		Code:    code,
		Message: message,
	}
}

func HandleResponseError(w http.ResponseWriter, err error) {
	respErr := &ErrorMessage{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
	var m *ErrorMessage
	if ok := errors.As(err, &m); ok {
		respErr.Code = m.Code
		respErr.Message = m.Message
	}
	httpx.WriteJson(w, http.StatusBadRequest, respErr)
}
