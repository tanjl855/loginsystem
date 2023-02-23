package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"loginsystem/Log"
	"loginsystem/models"
	"loginsystem/tools"
	"net/http"
)

// SetRespWriter 设置响应头部信息
func SetRespWriter(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //ajax跨域cookie丢失问题
	w.Header().Set("Content-Type", "application/json") //告诉前端用json格式
}

// Indexer 加载index页面
func Indexer(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		Log.ErrorLog.Printf("ParseFiles error:%v", err)
		_, err = w.Write([]byte("加载index.html出错"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		Log.ErrorLog.Printf("Execute error:%v", err)
		_, err = w.Write([]byte("加载index.html出错"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
}

// Register 注册接口
func Register(w http.ResponseWriter, r *http.Request) {
	SetRespWriter(w)
	user := models.Userinfo{}
	resp := models.Resp{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Log.ErrorLog.Printf("read body error:%v", err)
		_, err = w.Write([]byte("read body error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	Log.Info.Printf("read body success:%v", string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		Log.ErrorLog.Printf("Unmarshal error:%v", err)
		_, err = w.Write([]byte("Unmarshal error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	err = models.CreateUser(&user) //创建user到postgres
	if err != nil {
		resp = models.Resp{
			Code:    models.Failed, //注册失败201
			Data:    err.Error(),
			Message: "注册失败",
		}
		byteJson, err := json.Marshal(&resp)
		if err != nil {
			Log.ErrorLog.Printf("(注册失败)json marshal error:%v", err)
			_, err = w.Write([]byte("(注册失败)Marshal json error"))
			if err != nil {
				Log.ErrorLog.Printf("write error:%v", err)
				return
			}
			return
		}
		_, err = w.Write(byteJson)
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	resp = models.Resp{
		Code:    models.Success, //注册成功200
		Data:    nil,
		Message: "注册成功",
	}
	byteJson, err := json.Marshal(&resp)
	if err != nil {
		Log.ErrorLog.Printf("json marshal error:%v", err)
		_, err = w.Write([]byte("Marshal json error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	_, err = w.Write(byteJson)
	if err != nil {
		Log.ErrorLog.Printf("write error:%v", err)
		return
	}
}

// Login 登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	SetRespWriter(w)
	user := models.Userinfo{}
	resp := models.Resp{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Log.ErrorLog.Printf("read body error:%v", err)
		_, err = w.Write([]byte("read body error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	Log.Info.Printf("read body success:%v", string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		Log.ErrorLog.Printf("unmarshal json error:%v", err)
		_, err = w.Write([]byte("unmarshal json error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	err = models.CheckUser(&user)
	if err != nil {
		resp = models.Resp{
			Code:    models.Failed, //登陆失败
			Data:    err.Error(),
			Message: "登录失败",
		}
		byteJson, err := json.Marshal(&resp)
		if err != nil {
			Log.ErrorLog.Printf("marshal json error:%v", err)
			_, err = w.Write([]byte("marshal json error"))
			if err != nil {
				Log.ErrorLog.Printf("write error:%v", err)
				return
			}
			return
		}
		_, err = w.Write(byteJson)
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	err = models.GetUser(&user) //从postgres中获取个人用户信息
	if err != nil {
		_, err = w.Write([]byte("服务端出错"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	token := models.GenerateToken(user.Name, tools.Key) //生成token
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: false,
		Path:     "/", //cookie存储路径设置为根路径
	})
	resp = models.Resp{
		Code:    models.Success, //登陆成功
		Data:    user,
		Message: "登录成功",
	}
	byteJson, err := json.Marshal(&resp)
	if err != nil {
		Log.ErrorLog.Printf("marshal json error:%v", err)
		_, err = w.Write([]byte("marshal json error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	_, err = w.Write(byteJson)
	if err != nil {
		Log.ErrorLog.Printf("write error:%v", err)
		return
	}
}

// Update 更新接口
func Update(w http.ResponseWriter, r *http.Request) {
	SetRespWriter(w)
	resp := models.Resp{}
	user := models.Userinfo{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Log.ErrorLog.Printf("read body error:%v", err)
		_, err = w.Write([]byte("read body err"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	Log.Info.Printf("read body success:%v", string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		Log.ErrorLog.Printf("unmarshal error:%v", err)
		_, err = w.Write([]byte("Unmarshal error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	name := r.Context().Value("name")
	password, err := models.FindUser(name.(string)) //从postgres中查找用户名
	if err != nil {
		_, err = w.Write([]byte("没有权限"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		Log.ErrorLog.Printf("找不到要修改的User,error:%v", err)
		return
	}
	err = models.UpdateUser(name.(string), &user) //更新用户信息到postgres中
	if err != nil {
		resp = models.Resp{
			Code:    models.Failed,
			Data:    err.Error(),
			Message: "更新失败",
		}
		jsonByte, err := json.Marshal(&resp)
		if err != nil {
			_, err = w.Write([]byte("marshal json error"))
			if err != nil {
				Log.ErrorLog.Printf("write error:%v", err)
				return
			}
			Log.ErrorLog.Printf("marshal json error:%v", err)
			return
		}
		_, err = w.Write(jsonByte)
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		Log.ErrorLog.Printf("UpdateUser error:%v", err)
		return
	}
	if name != user.Name || password != tools.Md5Encrypt(user.Password) {
		resp = models.Resp{
			Code:    models.UpdateToken, //提示删除cookie 300
			Data:    user,
			Message: "更新成功",
		}
		jsonByte, err := json.Marshal(&resp)
		if err != nil {
			_, err = w.Write([]byte("marshal json error"))
			if err != nil {
				Log.ErrorLog.Printf("write error:%v", err)
				return
			}
			Log.ErrorLog.Printf("marshal json error:%v", err)
			return
		}
		_, err = w.Write(jsonByte)
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	resp = models.Resp{
		Code:    models.Success,
		Data:    user,
		Message: "更新成功",
	}
	jsonByte, err := json.Marshal(&resp)
	if err != nil {
		_, err = w.Write([]byte("marshal json error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		Log.ErrorLog.Printf("marshal json error:%v", err)
		return
	}
	_, err = w.Write(jsonByte)
	if err != nil {
		Log.ErrorLog.Printf("write error:%v", err)
		return
	}
}

// Delete 删除接口（管理员listed）
func Delete(w http.ResponseWriter, r *http.Request) {
	SetRespWriter(w)
	user := models.Userinfo{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Log.ErrorLog.Printf("read body error:%v", err)
		_, err = w.Write([]byte("read body error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	Log.Info.Printf("read body success:%v", string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		Log.ErrorLog.Printf("unmarshal json error:%v%v", err)
		_, err = w.Write([]byte("Unmarshal json error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	name := user.Name
	_, err = models.FindUser(name)
	if err != nil {
		Log.ErrorLog.Printf("找不到要删除的User,error:%v", err)
		_, err = w.Write([]byte("没有权限"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	resp := models.Resp{}
	err = models.DeleteUser(name) //从postgres中删除用户
	if err != nil {
		resp = models.Resp{
			Code:    models.Failed,
			Data:    err.Error(),
			Message: "删除失败",
		}
		jsonByte, err := json.Marshal(&resp)
		if err != nil {
			_, err = w.Write([]byte("marshal json error"))
			if err != nil {
				Log.ErrorLog.Printf("write error:%v", err)
				return
			}
			Log.ErrorLog.Printf("marshal json error:%v", err)
			return
		}
		_, err = w.Write(jsonByte)
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		return
	}
	resp = models.Resp{
		Code:    models.Success,
		Data:    nil,
		Message: "删除成功",
	}
	jsonByte, err := json.Marshal(&resp)
	if err != nil {
		_, err = w.Write([]byte("marshal json error"))
		if err != nil {
			Log.ErrorLog.Printf("write error:%v", err)
			return
		}
		Log.ErrorLog.Printf("marshal json error:%v", err)
		return
	}
	_, err = w.Write(jsonByte)
	if err != nil {
		Log.ErrorLog.Printf("write error:%v", err)
		return
	}
}

// Loaded 接口 渲染出个人信息页面
func Loaded(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/html/load.html")
	if err != nil {
		Log.ErrorLog.Printf("ParseFiles error:%v", err)
		_, err = w.Write([]byte("解析页面error"))
		if err != nil {
			Log.ErrorLog.Printf("ParseFiles error:%v", err)
			return
		}
		return
	}
	user := models.Userinfo{}
	if err != nil {
		Log.ErrorLog.Printf("get cookie error:%v", err)
		return
	}
	NameAny := r.Context().Value("name")
	user.Name = NameAny.(string)
	err = models.GetUser(&user) //从postgres中获取用户信息
	if err != nil {
		Log.ErrorLog.Printf("GetUser error:%v", err)
		return
	}
	err = tmpl.Execute(w, user)
	if err != nil {
		Log.ErrorLog.Printf("Execute error:%v", err)
		_, err = w.Write([]byte("渲染页面error"))
		if err != nil {
			Log.ErrorLog.Printf("Execute error:%v", err)
			return
		}
		return
	}
}

// Listed 接口 处理管理员页面信息
func Listed(w http.ResponseWriter, r *http.Request) {
	users := []models.Userinfo{}
	tmpl, err := template.ParseFiles("./static/html/list.html")
	if err != nil {
		Log.ErrorLog.Printf("ParseFiles error:%v", err)
		_, err = w.Write([]byte("解析页面error"))
		if err != nil {
			Log.ErrorLog.Printf("ParseFiles error:%v", err)
			return
		}
		return
	}
	users, err = models.GetAllUsers() //从postgres中获取所有用户信息
	if err != nil {
		Log.ErrorLog.Printf("GetAllUsers error:%v", err)
		return
	}
	err = tmpl.Execute(w, users)
	if err != nil {
		Log.ErrorLog.Printf("Execute error:%v", err)
		_, err = w.Write([]byte("渲染页面error"))
		if err != nil {
			Log.ErrorLog.Printf("Execute error:%v", err)
			return
		}
		return
	}
}
