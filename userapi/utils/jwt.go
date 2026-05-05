package utils

import (
    "errors"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var secretKey = []byte("your-secret-key")  // 签名密钥

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims  // 内置字段，包含过期时间等
}

// 生成 token
func GenerateToken(userID int) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

// 验证 token
func ParseToken(tokenStr string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }
    claims, ok := token.Claims.(*Claims)
    if !ok {
        return nil, errors.New("token 无效")
    }
    return claims, nil
}