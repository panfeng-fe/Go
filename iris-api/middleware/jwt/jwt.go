package jwt
//
//import (
//	"iris-api/config"
//	"iris-api/model"
//	"time"
//
//	jwt "github.com/dgrijalva/jwt-go"
//)
//
//var app = config.Setting.App
//
//// CreateToken ...
//func CreateToken(LoginInfo model.LoginInfo) (Token string) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"iat":      time.Now().Unix(),
//		"phone":    LoginInfo.Phone,
//		"password": LoginInfo.Password,
//		"exp":      time.Now().Add(time.Second * 30).Unix(), // 添加过期时间为30分钟
//	})
//	Token, _ = token.SignedString([]byte(app.TokenKey))
//	return Token
//}
//
//// CheckToken ...
//func CheckToken() {
//	// jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
//	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//	// 		return []byte(app.TokenKey), nil
//	// 	},
//	// 	SigningMethod: jwt.SigningMethodHS256,
//	// })
//
//}
