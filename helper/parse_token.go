package helper

import (
	"edu-management-system/schema"
	"github.com/dgrijalva/jwt-go"
)

/** 解析token
 * @description 使用私钥解码
 * @since 2023/1/215:53
 * @param token JWT字符串
 * @param encryptToken 私钥
 * @return 解码的用户信息或解码错误信息
 *  */

func ParseToken(token string, encryptToken []byte) (*schema.Claim, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &schema.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return encryptToken, nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*schema.Claim); ok && tokenClaim.Valid {
			return claims, err
		}
	}

	return nil, err
}
