package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-template/pkg/utils"
	"go-template/setting"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(setting.Conf.JwtConfig.Secret),
	}
}

var mySecret = []byte("夏天夏天悄悄过去")

// GenToken 生成JWT
func (j *JWT) GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	nowTime := time.Now()
	expireTime, err := utils.ParseDuration(setting.Conf.JwtConfig.ExpireTime)
	if err != nil {
		return "", errors.New("expire time is invalid")
	}
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(expireTime)), // 过期时间
			Issuer:    setting.Conf.Name,                           // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
