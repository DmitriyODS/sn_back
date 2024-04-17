package models

import "fmt"

type Response struct {
	Ok   bool   `json:"ok"`
	Code int    `json:"code"`
	Err  string `json:"err"`
	Data any    `json:"data"`
}

func (r *Response) SetResponse(code int, data any, errValue any) *Response {
	r.Ok = errValue == nil
	r.Code = code
	r.Data = data

	if !r.Ok {
		r.Err = fmt.Sprintf("Ошибка: %v", errValue)
	}

	return r
}

func MakeResponse(code int, data any, errValue any) *Response {
	resp := &Response{
		Ok:   errValue == nil,
		Code: code,
		Err:  "",
		Data: data,
	}

	if !resp.Ok {
		resp.Err = fmt.Sprintf("Ошибка: %v", errValue)
	}

	return resp
}
