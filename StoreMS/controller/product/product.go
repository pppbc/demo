package product

import (
	"StoreMS/config"
	"StoreMS/model"
	"StoreMS/tools"
	"github.com/gin-gonic/gin"
)

func UploadHandler(r *gin.Context) {
	file, _ := r.FormFile("file")
	name := r.PostForm("name")

	filename := name + ".png"

	if err := r.SaveUploadedFile(file, config.GetLocalPath()+filename); err != nil {
		tools.JMFail(r, "upload  failed")
		return
	} else {
		tools.JMFail(r, "upload succeed")
	}
}

//获取商品信息
func GetProductsHandler(r *gin.Context) {
	//创建商品结构体对象
	p := &model.Product{}

	//获取数据库数据
	info, err := p.PrintAllProducts()

	//判断，响应
	if err == true {
		tools.JMOk(r, info, "get goodsList success")
	} else {
		tools.JMFail(r, "get goodsList failed")
	}
}
