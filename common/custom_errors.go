package common

import (
	"encoding/json"
)

type grpcError struct {
	Code int
	Msg string
}

func (e *grpcError) Error() string {
	// 返回一个带code和msg的json格式的错误信息
	// 以便客户端解析是何种错误
	coded, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(coded)
}

func NewError(code int, msg string) *grpcError {
	return &grpcError{
		Code: code,
		Msg:  msg,
	}
}
