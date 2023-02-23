package routers

import (
	"loginsystem/Log"
	"loginsystem/handler"
	"loginsystem/middleware"
	"net/http"
)

func RunRouters() {
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("static/html")))) //加载静态资源
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	http.HandleFunc("/", middleware.LoginMiddleware(http.HandlerFunc(handler.Indexer)))                                     //首页
	http.HandleFunc("/public/register", middleware.LoginMiddleware(http.HandlerFunc(handler.Register)))                     //注册
	http.HandleFunc("/public/login", middleware.LoginMiddleware(http.HandlerFunc(handler.Login)))                           //登录
	http.HandleFunc("/adm/update", middleware.LoginMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.Update)))) //更新
	http.HandleFunc("/adm/delete", middleware.LoginMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.Delete)))) //删除
	http.HandleFunc("/adm/loaded", middleware.LoginMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.Loaded)))) //个人页面
	http.HandleFunc("/adm/listed", middleware.LoginMiddleware(middleware.AuthList(http.HandlerFunc(handler.Listed))))       //管理员页面
	Log.Info.Println("RunRouters设置完成...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		Log.ErrorLog.Printf("ListenAndServe error:%v", err)
		return
	}
}
