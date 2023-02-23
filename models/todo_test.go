package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"loginsystem/tools"
	"testing"
	"time"
)

func TestCheckUser(t *testing.T) {
	users := []Userinfo{
		{
			Name:     "tjl",
			Password: "tjl",
		},
		{
			Name:     "root",
			Password: "root",
		},
		{
			Name:     "lzy",
			Password: "lzy",
		},
		{
			Name:     "u1",
			Password: "u1",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when a stub database connection", err)
	}
	defer db.Close()
	for _, user := range users {
		rows := sqlmock.NewRows([]string{"password"}).AddRow(tools.Md5Encrypt(user.Password))
		mock.ExpectQuery("select password from userinfo").WithArgs(user.Name, "0").WillReturnRows(rows)
		DB = db
		if err = CheckUser(&user); err != nil {
			t.Errorf("error was not expected while check user:%s", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}
}

func TestCreateUser(t *testing.T) {
	users := []Userinfo{
		{
			Name:      "lzy",
			Password:  "lzy",
			Email:     "lzy@example.com",
			Phone:     "19872773116",
			CreatedAt: time.Now().String(),
		},
		{
			Name:      "u1",
			Password:  "u1",
			Email:     "u1@example.com",
			Phone:     "19872773116",
			CreatedAt: time.Now().String(),
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when a stub databse connection", err)
	}
	defer db.Close()
	var rows *sqlmock.Rows
	rows = sqlmock.NewRows([]string{"name"}).AddRow("root")
	for _, user := range users {
		mock.ExpectQuery("select name from userinfo").WithArgs("0").WillReturnRows(rows)
		mock.ExpectExec("insert into userinfo").
			WithArgs(user.Name, tools.Md5Encrypt(user.Password), user.Email, user.Phone, time.Now()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		DB = db
		if err = CreateUser(&user); err != nil {
			t.Errorf("error was not expected while create user:%s", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("expected were unfulfilled expectations:%s", err)
		}
		rows = sqlmock.NewRows([]string{"name"}).AddRow(user.Name)
	}
}

func TestDeleteUser(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {
	User := Userinfo{
		Name:     "tjl",
		Password: "tjl",
		Email:    "tjl@tjl.com",
		Phone:    "18200871880",
	}
	OldUser := Userinfo{
		Name:     "root",
		Password: "root",
		Email:    "root@root.com",
		Phone:    "18200871880",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when a stub databse connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"name", "email", "phone", "created_at"}).
		AddRow(OldUser.Name, OldUser.Email, OldUser.Phone, time.Now())
	mock.ExpectQuery("select name, email, phone, created_at from userinfo").
		WithArgs(OldUser.Name, "0").WillReturnRows(rows)
	mock.ExpectPrepare("update userinfo").ExpectExec().WithArgs(User.Name, tools.Md5Encrypt(User.Password), User.Email, User.Phone, time.Now(), "root", "0").
		WillReturnResult(sqlmock.NewResult(1, 1))

	DB = db
	if err = UpdateUser("root", &User); err != nil {
		t.Errorf("error was not expected while create user:%s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expected were unfulfilled expectations:%s", err)
	}
}

func TestFindUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when a stub database connection", err)
	}
	defer db.Close()
	Names := []string{"tjl", "lzy", "root"}
	for _, name := range Names {
		rows := sqlmock.NewRows([]string{"password"}).AddRow(tools.Md5Encrypt(name))
		mock.ExpectQuery("select password from userinfo").WithArgs(name, "0").
			WillReturnRows(rows)
		DB = db
		if _, err = FindUser(name); err != nil {
			t.Errorf("error was not expected while check user:%s", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}
}

func TestGetUser(t *testing.T) {

}

func TestGetAllUsers(t *testing.T) {

}
