package tools

import (
	"github.com/gin-gonic/gin"
)

type Infos struct {
	Status int         `json:"status"`
	Detail  string      `json:"detail"`
	Info interface{} `json:"info,omitempty"`
}

func JMFail(r *gin.Context, msg string) {
	JsonMessage(r, 0, nil, msg)
}

func JMOk(r *gin.Context, data interface{}, msg string) {
	JsonMessage(r, 1, data, msg)
}

func JsonMessage(r *gin.Context, code1 int, data1 interface{}, msg1 string) {
	h := Infos{
		code1,
		msg1,
		data1,
	}
	r.JSON(200, h)
}
