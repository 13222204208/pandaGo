package utility

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

// 定义 JWT 密钥
var (
	jwtSecret         = []byte("panda-admin-jwt-secret") // 默认密钥
	accessExpiration  = 24 * time.Hour                   // 访问令牌有效期24小时
	refreshExpiration = 168 * time.Hour                  // 刷新令牌有效期7天(168小时)
)

// CustomClaims 自定义JWT声明结构
type CustomClaims struct {
	UserId int64  `json:"userId"`
	Type   string `json:"type"` // access 或 refresh
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userId int64, tokenType string) (string, error) {
	// 初始化JWT配置
	InitJwtConfig()

	// 确定过期时间
	var expiration time.Duration
	if tokenType == "access" {
		expiration = accessExpiration
	} else {
		expiration = refreshExpiration
	}

	// 创建声明
	claims := CustomClaims{
		UserId: userId,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "RapidGo",
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌有效性
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetUserId(ctx context.Context) int64 {
	return g.RequestFromCtx(ctx).GetCtxVar("userId").Int64()
}

// GetTokenExpires 获取令牌过期时间戳
func GetTokenExpires(tokenType string) int64 {
	var expiration time.Duration
	if tokenType == "access" {
		expiration = accessExpiration
	} else {
		expiration = refreshExpiration
	}

	return time.Now().Add(expiration).Unix()
}

// RefreshToken 刷新访问令牌
func RefreshToken(refreshToken string) (string, int64, error) {
	// 解析刷新令牌
	claims, err := ParseToken(refreshToken)
	fmt.Println("claims:", claims.Type)
	if err != nil {
		return "", 0, err
	}

	// 验证令牌类型
	if claims.Type != "refresh" {
		return "", 0, fmt.Errorf("invalid token type")
	}

	// 生成新的访问令牌
	accessToken, err := GenerateToken(claims.UserId, "access")
	if err != nil {
		return "", 0, err
	}

	// 获取新令牌的过期时间
	expires := GetTokenExpires("access")

	return accessToken, expires, nil
}

// GetUserIdFromToken 从令牌中获取用户ID
func GetUserIdFromToken(tokenString string) (int64, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserId, nil
}

// InitJwtConfig 初始化JWT配置
func InitJwtConfig() {
	// 创建一个背景上下文
	ctx := context.Background()

	// 从配置文件读取JWT密钥
	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	fmt.Println("secret:", secret)
	if secret != "" {
		jwtSecret = []byte(secret)
	}

	// 读取访问令牌过期时间（小时）
	if accessExp := g.Cfg().MustGet(ctx, "jwt.accessExpiration").Int(); accessExp > 0 {
		accessExpiration = time.Duration(accessExp) * time.Hour
	}

	// 读取刷新令牌过期时间（小时）
	if refreshExp := g.Cfg().MustGet(ctx, "jwt.refreshExpiration").Int(); refreshExp > 0 {
		refreshExpiration = time.Duration(refreshExp) * time.Hour
	}
	fmt.Println("我的jwtSecret:", string(jwtSecret))
}
