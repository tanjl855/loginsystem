package tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

var nowDate = "20221120"
var Key = nowDate + "tanJlGitdayangGcom"

func Md5Encrypt(data string) string {
	md5Ctx := md5.New()          //md5 init
	md5Ctx.Write([]byte(data))   //md5 updata
	cipherStr := md5Ctx.Sum(nil) //md5 final
	return hex.EncodeToString(cipherStr)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// WithCode 对密码进行加密
func WithCode(p string) string {
	return RandStringBytes(5) + p
}
