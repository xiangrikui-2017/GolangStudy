package result

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

type PageResponse struct {
	Response
	Page     int `json:"page"`      // 当前页码
	PageSize int `json:"page_size"` // 每页数量
	Total    int `json:"total"`     // 总数
}

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func Success(data interface{}) *Response {
	return &Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
}

func PageSuccess(data interface{}, page int, pageSize int, total int) *PageResponse {
	return &PageResponse{
		Response: Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: data,
		},
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}

func Error(msg string) *Response {
	return &Response{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: nil,
	}
}
