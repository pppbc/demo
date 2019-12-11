package tools

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//查找字符串指定部分
func StrFinder(str string, begin int, num int) string {
	strRune := []rune(str)

	var strRuneNew []rune
	for i := begin; i < begin+num; i++ {
		strRuneNew = append(strRuneNew, strRune[i])
	}
	return string(strRuneNew)
}

//生成4位随机数
func RandNum() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

//生成订单编号

func CreateOrderNo(userid int) int64 {
	formatNow := time.Now().Format("2006-01-02 15:04:05")
	u := fmt.Sprintf("%2d", userid)
	strNew := StrFinder(formatNow, 2, 2) + StrFinder(formatNow, 5, 2) + StrFinder(formatNow, 8, 2) + StrFinder(formatNow, 11, 2) + u + RandNum()

	a, err := strconv.ParseInt(strNew, 10, 64)
	if err != nil {
		return -1
	}
	return a
}
