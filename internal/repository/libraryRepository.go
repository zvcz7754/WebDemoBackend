package repository //主要處理ORM

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	. "webDemoBackend/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "root:7754@tcp(localhost:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func init() {
	if err != nil {
		panic("failed to connect database")
	}

	var BookSlice []BookDO
	db.Table("books").Find(&BookSlice)
	fmt.Print(BookSlice, 1)
}

func CreateData(db *gorm.DB) {
	db.Create(&BookDO{Name: "bb", Price: 126})
}

func ReadData(db *gorm.DB) {
	db.First(&Book{}, 1) // 根据整形主键查找
	//db.First(&Book{}, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
}

func DeleteData(db *gorm.DB) {
	db.Delete(&Book{}, 10)
}
