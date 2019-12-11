package tools

import (
	"github.com/gin-gonic/gin"
)

type Info struct {
	Status int         `json:"status"`
	Desc   string      `json:"desc"`
	Data   interface{} `json:"data,omitempty"`
}

func JMFail(r *gin.Context, msg string) {
	JsonMessage(r, 0, nil, msg)
}

func JMOk(r *gin.Context, data interface{}, msg string) {
	JsonMessage(r, 1, data, msg)
}

func JsonMessage(r *gin.Context, status int, data interface{}, desc string) {
	h := Info{
		status,
		desc,
		data,
	}
	r.JSON(200, h)
}
