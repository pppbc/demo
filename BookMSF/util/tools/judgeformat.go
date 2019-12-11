package tools

import "regexp"

//判断e-mail格式
func JudgeEmail(email string) bool {
	//匹配样式
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	//写入匹配库
	reg := regexp.MustCompile(pattern)
	//将传入的值比较
	return reg.MatchString(email)
}

//判断电话号码格式
func JudgePhone(phone string) bool {
	pattern := "^((13[0-9])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}

//判断密码格式
func JudgePassword(password string) bool {
	pattern := "^[a-z0-9_-]{6,18}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}

func JudgeNotNull(str string) bool {
	if str == "" {
		return false
	}
	return true
}
