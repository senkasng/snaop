package db

import (

	//"github.com/labstack/echo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"../util"
)


var connect string 


type DB struct { //db 结构体
	Mydb *sql.DB
}

func init() { //解析数据库连接
	obj := util.ParseObject()
	connect = obj["mysql"]["connect"]
	
}

func Newdb() DB {
	var dbinstance  DB
	db, err := sql.Open("mysql", connect)
	if err != nil {
		fmt.Println(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("db connect failed")
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(1)
	dbinstance.Mydb = db
	return dbinstance
}


func (db DB) Close() error {
	err := db.Mydb.Close()
	if err != nil {
		return err
	}
	return nil
}