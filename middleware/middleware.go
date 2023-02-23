package middleware

import (
	"context"
	"loginsystem/Log"
	"loginsystem/models"
	"net/http"
)

// LoginMiddleware 日志中间件
func LoginMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info.Println(r.URL.Path[:])
		next.ServeHTTP(w, r)
	}
}

// AuthMiddleware 中间件验证token
func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if next == nil {
			next = http.DefaultServeMux
		}
		cookie, err := r.Cookie("token")
		if cookie == nil || err != nil {
			Log.ErrorLog.Printf("没有权限:%v", err)
			models.Error(w, http.StatusUnauthorized)
			return
		}
		name := models.ParseToken(cookie.Value)
		var ctx context.Context
		ctx = context.WithValue(context.Background(), "name", name) //保存上下文信息
		r = r.WithContext(ctx)
		_, err = models.FindUser(name) //这里找一下数据库，管理员删除后，数据库没数据，过不了验证
		if err != nil {
			Log.ErrorLog.Printf("没有权限:%v", err)
			models.Error(w, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func AuthList(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if next == nil {
			next = http.DefaultServeMux
		}
		cookie, err := r.Cookie("token")
		if cookie == nil || err != nil {
			Log.ErrorLog.Printf("没有权限:%v", err)
			models.Error(w, http.StatusUnauthorized)
			return
		}
		name := models.ParseToken(cookie.Value)
		_, err = models.FindUser(name)
		if err != nil || name != "root" {
			Log.ErrorLog.Printf("没有权限:%v", err)
			models.Error(w, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
