package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	dbConn *sql.DB
	err error 
)

func init() {
	dbConn, err = sql.Open("mysql", "root:zhaomeiping@tcp(192.168.189.155:3306)/VIDEO_DB?charset=utf8")
	if err != nil {
		fmt.Println("mysql Open failed")
		panic(err.Error())
	}
	//defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		fmt.Println("mysql Ping failed")
		panic(err.Error())
	}
}
