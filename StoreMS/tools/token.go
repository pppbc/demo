package tools

import (
	"StoreMS/database"
	"StoreMS/model"
	"StoreMS/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const SecretKey = "efdegv"

type JWTCustomClaims struct {
	jwt.StandardClaims
	Uid  int    `json:"uid"`
	UDID string `json:"udid"`
}

//⽣成 token

func CreateToken(SecretKey []byte, issuer string, Uid int, UDID string) (tokenString string, err error) {
	claims := &JWTCustomClaims{
		jwt.StandardClaims{ExpiresAt: int64(time.Now().Add(time.Hour * 360).Unix()), Issuer: issuer},
		Uid,
		UDID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

//解析 token

func ParseToken(tokenSrt string, SecretKey []byte) (jwt.Claims, error) {
	var token *jwt.Token
	token, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	} else {
		if token.Valid {
			claims := token.Claims
			return claims, nil
		}
	}
	return nil, nil
}

//验证
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1:检查token参数
		token := c.Request.FormValue("token")
		if token == "" {
			JMFail(c, "缺少token参数")
			c.Abort()
			return
		}
		//2 :解析token参数
		claims, err := ParseToken(token, []byte(SecretKey))
		if nil != err {
			JMFail(c, "token异常，重新登陆")
			c.Abort()
			return
		}
		//
		//第⼀步对比是否过时

		//第二步对应与redis的是否一致，都通过了就可以继续了
		userID := int(claims.(jwt.MapClaims)["uid"].(float64))

		result, err := redis.DoRedisCommand("hget token_list token:user_id:"+strconv.Itoa(userID), "string")
		//如果redis没有token，则从数据库获取，重新存
		var info model.User
		if result == nil {
			database.DB.Model(&model.User{}).Where("id = ?", userID).First(&info)
			redis.DoRedisCommand("hset token_list token:user_id:"+strconv.Itoa(info.ID)+" "+info.Token, "int")
			result, err = redis.DoRedisCommand("hget token_list token:user_id:"+strconv.Itoa(userID), "string")
		}
		//和redis里的token比较
		if result != token {
			JMFail(c, "token异常，重新登陆")
			c.Abort()
			return
		}
		c.Next()
	}
}
