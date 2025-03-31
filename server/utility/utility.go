package utility

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword 使用 bcrypt 加密密码
func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword 比较密码是否匹配
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GetCurrentTime 获取当前时间
func GetCurrentTime() *gtime.Time {
	return gtime.Now()
}

// GetCurrentTimestamp 获取当前时间戳（秒）
func GetCurrentTimestamp() int64 {
	return gtime.Timestamp()
}

// FormatTime 格式化时间为字符串，默认格式：Y-m-d H:i:s
func FormatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func StrToFormatTime(timeStr string) string {
	// 解析时间2025-03-25 02:07:37 +0800 CST
	t, err := gtime.StrToTime(timeStr)
	if err != nil {
		fmt.Println("解析时间出错:", err)
		return ""
	}

	// 格式化为2006-01-02 15:04:05格式
	formatted := t.Format("Y-m-d H:i:s")
	//fmt.Println("格式化后的时间:", formatted) // 输出: 2025-03-25 02:07:37
	if formatted == "0001-01-01 00:00:00" {
		return ""
	}
	return formatted
}

// FormatFileSize 格式化文件大小
func FormatFileSize(fileSize int64) string {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize))
	} else if fileSize < 1024*1024 {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/1024)
	} else if fileSize < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/(1024*1024*1024))
	}
}

// SafeGo 安全的调用协程，遇到错误时输出错误日志而不是抛出panic
func SafeGo(ctx context.Context, f func(ctx context.Context), lv ...int) {
	g.Go(ctx, f, func(ctx context.Context, err error) {
		var level = glog.LEVEL_ERRO
		if len(lv) > 0 {
			level = lv[0]
		}
		Logf(level, ctx, "SafeGo exec failed:%+v", err)
	})
}

func Logf(level int, ctx context.Context, format string, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debugf(ctx, format, v...)
	case glog.LEVEL_INFO:
		g.Log().Infof(ctx, format, v...)
	case glog.LEVEL_NOTI:
		g.Log().Noticef(ctx, format, v...)
	case glog.LEVEL_WARN:
		g.Log().Warningf(ctx, format, v...)
	case glog.LEVEL_ERRO:
		g.Log().Errorf(ctx, format, v...)
	case glog.LEVEL_CRIT:
		g.Log().Criticalf(ctx, format, v...)
	case glog.LEVEL_PANI:
		g.Log().Panicf(ctx, format, v...)
	case glog.LEVEL_FATA:
		g.Log().Fatalf(ctx, format, v...)
	default:
		g.Log().Errorf(ctx, format, v...)
	}
}
