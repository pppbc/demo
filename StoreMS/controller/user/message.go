package user

import (
	"StoreMS/tools"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// func TimeToStr(time.Time t)string{
// 		//2006-01-02 15:04:05是golang诞生的日期
// 		formatNow := t.Format("YYYYMMDDHHMMSS")
// 		return formatNow
// }

func PostMessage(c *gin.Context) {
	phone := c.Request.FormValue("phone")
	//验证码在后台生成，同时将生成的验证码返回给前台
	//code := c.Request.FormValue("code")

	if tools.JudgePhone(phone) == false {
		tools.JMFail(c, "参数错误")
		return
	}

	currentTime := time.Now()
	name := "LT-wangwang"
	seed := fmt.Sprintf("%4d%02d%02d%2d%2d%2d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	key := MD5ToString(MD5ToString("bnzcsdw1") + seed)
	dest := phone
	mesCode := GenValidateCode(6)
	content := "【数字积分交易所】你好：你的验证码为：" + mesCode

	url := fmt.Sprintf("name=%s&seed=%s&key=%s&dest=%s&content=%s", name, seed, key, dest, content)
	resp, err := http.Post("http://223.111.192.46:8080/eums/sms/utf8/send.do?",
		"application/x-www-form-urlencoded",
		strings.NewReader(url))
	fmt.Println(url)
	if err != nil {
		log.Println(err)
		tools.JMFail(c, "发送错误")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		tools.JMFail(c, "信息错误")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"obj":    string(body),
		"code":   mesCode,
	})
}

func MD5ToString(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	//string.Builder是go 1.10之后新添加的内容
	var sb bytes.Buffer
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
