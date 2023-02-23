package models

import (
	"loginsystem/tools"
	"net/http"
)

const (
	Success = 200 //接口调用成功
	Failed  = 201 //接口调用失败

	UpdateToken = 300 //需要删除cookie信息
	ExistedNum  = "0" //0未删除
	DeletedNum  = "1" //1已删除
)

type Resp struct {
	Code    int         `json:"code"`    //接口调用响应码
	Data    interface{} `json:"data"`    //接口调用返回数据
	Message string      `json:"message"` //接口调用返回信息
}

type Userinfo struct {
	UserId    string `json:"user_id"`    //用户id 自增 主键索引
	Name      string `json:"name"`       //用户名 已存在用户中唯一 普通索引
	Password  string `json:"password"`   //用户密码
	Email     string `json:"email"`      //邮箱
	Phone     string `json:"phone"`      //手机号
	CreatedAt string `json:"created_at"` //创建时间
	UpdatedAt string `json:"updated_at"` //更新时间
	DeletedAt string `json:"deleted_at"` //删除状态 实现软删除
	Status    string `json:"status"`     //状态（区分用户身份）
}

// GenerateToken 生成Token值
func GenerateToken(userName string, key string) string {
	return tools.WithCode(userName) + key
}

// ParseToken 解析token"
func ParseToken(token string) string {
	length := len(tools.Key)
	return string([]byte(token)[5 : len(token)-length])
}

func Error(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
