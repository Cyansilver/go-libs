package server

import (
	errp "github.com/cyansilver/go-libs/err"
)

// Result presents the http response
type Result struct {
	Code int32                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"message"`
}

func DefaultResult() *Result {
	return &Result{Code: 0, Msg: "Success", Data: make(map[string]interface{})}
}

func (r *Result) AddData(key string, data interface{}) {
	r.Data[key] = data
}

func (r *Result) SetError(er *errp.Error) {
	r.Code = er.Code
	r.Msg = er.Msg
}
