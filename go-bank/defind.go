package main

const (
	ResponseOk   = 0
	ResponseFail = 1
)

// Api 统一响应体
type ApiResult[T any] struct {
	Code    uint8  `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// 响应失败
func Failed[T any](msg string) *ApiResult[T] {
	return &ApiResult[T]{Code: ResponseFail, Message: msg}
}

// 响应成功
func Success[T any](data T) *ApiResult[T] {
	return &ApiResult[T]{Code: ResponseOk, Message: "successful", Data: data}
}
