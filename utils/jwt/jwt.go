package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("classroom_system") //jwt密钥

type Claims struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 定义过期时间
const ExpiredTime = time.Hour * 2

// 生成token
func GenerateToken(id int64, name string, password string) (string, error) {
	nowTime := time.Now()
	claims := Claims{
		id,
		name,
		password,
		jwt.StandardClaims{
			ExpiresAt: nowTime.Add(ExpiredTime).Unix(), //过期时间
			Issuer:    "classroom_system",              //签发人
		},
	}
	//使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的secret签名并获得完整的编码后的字符串token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 验证通过，返回用户信息
		// 从token中获取到Claims
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
