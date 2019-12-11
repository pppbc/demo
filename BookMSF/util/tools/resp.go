package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, 0, nil, msg)
}
func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 1, data, msg)
}
func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {

	w.Header().Set("Content-Type", "application/json")
	//设置成200状态
	w.WriteHeader(http.StatusOK)

	h := H{
		code,
		msg,
		data,
	}
	//结构体转换成json
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(ret)
}
