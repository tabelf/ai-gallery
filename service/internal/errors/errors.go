package errors

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

var (
	ErrServiceInternal = func(ctx context.Context) *ErrorMessage {
		return NewError(10000, "Service internal error", "服务内部错误")
	}
	ErrHeaderNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10001, "The request header information does not exist", "请求头信息不存在")
	}
	ErrHeaderParsing = func(ctx context.Context) *ErrorMessage {
		return NewError(10002, "Request header information parsing failed", "请求头信息解析失败")
	}
	ErrPermissionProvider = func(ctx context.Context) *ErrorMessage {
		return NewError(10003, "No valid identity permission provided", "未提供有效身份许可")
	}
	ErrTokenExpired = func(ctx context.Context) *ErrorMessage {
		return NewError(10004, "Token has expired", "令牌已失效")
	}
	ErrTokenValid = func(ctx context.Context) *ErrorMessage {
		return NewError(10005, "No valid token provided", "未提供有效令牌")
	}
)

var (
	ErrTaskImageBase64Parse = func(ctx context.Context) *ErrorMessage {
		return NewError(10100, "Image base64 parsing error", "图片base64解析错误")
	}
	ErrUploadServiceNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10101, "Upload service does not exist", "上传服务不存在")
	}
)

var (
	ErrAccountNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10200, "Account does not exist", "账号不存在")
	}
	ErrAccountDisabled = func(ctx context.Context) *ErrorMessage {
		return NewError(10201, "Account is unavailable", "账号不可用")
	}
	ErrAccountPwd = func(ctx context.Context) *ErrorMessage {
		return NewError(10202, "wrong password", "密码错误")
	}
	ErrAccountTokenGen = func(ctx context.Context) *ErrorMessage {
		return NewError(10203, "token invalid data", "无效 token 数据")
	}
	ErrAccountPermission = func(ctx context.Context) *ErrorMessage {
		return NewError(10204, "Account has no permissions", "账号无权限")
	}
	ErrAccountTokenNotFound = func(ctx context.Context) *ErrorMessage {
		return NewError(10205, "token does not exist", "token不存在")
	}
	ErrAccountUsernameExist = func(ctx context.Context) *ErrorMessage {
		return NewError(10206, "Username already exists", "用户名已存在")
	}
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Desc    string `json:"desc"`
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewError(code int, message, desc string) *ErrorMessage {
	return &ErrorMessage{
		Code:    code,
		Message: message,
		Desc:    desc,
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
		respErr.Desc = m.Desc
	}
	httpx.WriteJson(w, http.StatusBadRequest, respErr)
}
