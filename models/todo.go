package models

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"loginsystem/Log"
	"loginsystem/conf"
	"loginsystem/tools"
	"strings"
	"time"
)

var DB *sql.DB //全局DB 用于postgres CRUD

func InitDB() (db *sql.DB) {
	//dsn:="postgres://postgres:postgres@172.20.10.4:5432/loginsystem?sslmode=disable"
	Hconf := conf.TanConfig
	dsn := "postgres://" + Hconf.App.Username + ":" + Hconf.App.Password + "@" + Hconf.App.Host + ":" + Hconf.App.Port + "/loginsystem?sslmode=disable" //数据库：loginsystem
	db, err := sql.Open("postgres", dsn)
	PanicErr(err)
	err = db.Ping()
	PanicErr(err)
	Log.Info.Printf("database连接成功, address:%v, port:%v!", Hconf.App.Host, Hconf.App.Port)
	return
}

func CheckUser(User *Userinfo) (err error) {
	if User == nil {
		return errors.New("user is nil")
	}
	var password string
	err = DB.QueryRow("select password from userinfo where name = $1 and deleted_at = $2", User.Name, ExistedNum).Scan(&password)
	if err != nil {
		Log.ErrorLog.Printf("query error:%v", err)
		return
	}
	password = strings.TrimSpace(password)
	if password == tools.Md5Encrypt(User.Password) {
		return nil
	}
	err = errors.New("账号或密码不正确")
	Log.ErrorLog.Printf("账号密码不正确,error:%v", err)
	return
}

func CreateUser(User *Userinfo) (err error) {
	if User == nil {
		return errors.New("user is nil")
	}
	rows, err := DB.Query("select name from userinfo where deleted_at = $1", ExistedNum) //user唯一
	if err != nil {
		Log.ErrorLog.Printf("Query error:%v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var Ename string //已存在的用户名
		err = rows.Scan(&Ename)
		if err != nil {
			Log.ErrorLog.Printf("Scan error:%v", err)
			return
		}
		if User.Name == Ename {
			err = errors.New("UserName已存在")
			Log.ErrorLog.Printf("UserName已存在:%v", err)
			return
		}
	}
	User.Name = strings.TrimSpace(User.Name)
	User.Password = strings.TrimSpace(User.Password)
	if User.Name == "" {
		return errors.New("userName is null")
	} else if User.Password == "" {
		return errors.New("password is null")
	} else if User.Phone == "" {
		return errors.New("phone is null")
	} else if User.Email == "" {
		return errors.New("email is null")
	}
	MD5Password := tools.Md5Encrypt(User.Password)
	_, err = DB.Exec("insert into userinfo(name,password,email,phone,created_at)values($1,$2,$3,$4,$5)", User.Name, MD5Password, User.Email, User.Phone, time.Now())
	if err != nil {
		Log.ErrorLog.Printf("database exec error: %v", err)
		return
	}
	return
}

func FindUser(UserName string) (password string, err error) {
	err = DB.QueryRow("select password from userinfo where name = $1 and deleted_at = $2", UserName, ExistedNum).Scan(&password)
	if err != nil {
		Log.ErrorLog.Printf("无法找到UserName,error:%v", err)
		return "", err //err == sql.ErrNoRows
	}
	return
}

func UpdateUser(oldName string, User *Userinfo) (err error) {
	if User == nil {
		return errors.New("user is nil")
	}
	tmpUser := Userinfo{} //tmpUser用户存旧的用户信息
	tmpUser.Name = oldName
	err = GetUser(&tmpUser)
	if err != nil {
		Log.ErrorLog.Printf("GetUser error:%v", err)
		return
	}
	//修改信息如果为空，等于原来的信息
	if User.Name == "" {
		User.Name = tmpUser.Name
	}
	if User.Email == "" {
		User.Email = tmpUser.Email
	}
	if User.Phone == "" {
		User.Phone = tmpUser.Phone
	}
	stmt, err := DB.Prepare("update userinfo set name=$1, password=$2,email=$3,phone=$4,updated_at=$5 where name=$6 and deleted_at=$7")
	if err != nil {
		Log.ErrorLog.Printf("Prepare error:%v", err)
		return
	}
	_, err = stmt.Exec(User.Name, tools.Md5Encrypt(User.Password), User.Email, User.Phone, time.Now(), oldName, ExistedNum)
	return
}

func DeleteUser(name string) (err error) {
	stmt, err := DB.Prepare("update userinfo set deleted_at = $1 where name = $2 and deleted_at = $3")
	if err != nil {
		Log.ErrorLog.Printf("Prepare error:%v", err)
		return
	}
	res, err := stmt.Exec(DeletedNum, name, ExistedNum)
	if err != nil {
		Log.ErrorLog.Printf("Exec error:%v", err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		Log.ErrorLog.Printf("RowsAffected error:%v", err)
		return
	}
	Log.Info.Printf("RowsAffected:%v", affected)
	return
}

func GetAllUsers() (users []Userinfo, err error) {
	User := Userinfo{}
	rows, err := DB.Query("select name, email, phone, created_at from userinfo where deleted_at = $1", ExistedNum)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&User.Name, &User.Email, &User.Phone, &User.CreatedAt)
		if err != nil {
			Log.ErrorLog.Printf("Scan error: %v", err)
			return
		}
		users = append(users, User)
	}
	return
}

func GetUser(User *Userinfo) (err error) {
	if User == nil {
		return errors.New("user is nil")
	}
	err = DB.QueryRow("select name, email, phone, created_at from userinfo where name = $1 and deleted_at = $2", User.Name, ExistedNum).Scan(&User.Name, &User.Email, &User.Phone, &User.CreatedAt)
	if err != nil {
		Log.ErrorLog.Printf("QueryRow error:%v", err)
		return
	}
	return
}

func CloseDB(db *sql.DB) {
	defer func(db *sql.DB) {
		err := db.Close()
		PanicErr(err)
	}(db)
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
