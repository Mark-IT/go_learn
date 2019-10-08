package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"		// mysql驱动
	"github.com/jmoiron/sqlx"				// 封装mysql操作
)

/*
测试表
CREATE TABLE person (
    user_id int primary key auto_increment,
    username varchar(260),
    sex varchar(260),
    email varchar(260)
);

CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)


*/

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
