package repository

import (
	"database/sql"
	"fmt"
	"webDemoBackend/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() *sql.DB {
	db, err := sql.Open("mysql", "user:7754@tcp(localhost:3306)/mysqldb?charset=utf8")

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

	return db
}()

func InserData(db sql.DB, bookDatas []entity.Book) {
	insertStmt, err := db.Prepare("INSERT INTO employee (id, name) VALUES (?, ?);")
	if err != nil {
		panic(err)
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(305, "Sue")
	if err != nil {
		panic(err)
	}
	fmt.Println("成功插入資料 305, Sue")
}
