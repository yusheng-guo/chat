package common

import "github.com/golang-jwt/jwt"

// Claims 需求 包含需要通过 jwt 传输的数据
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
