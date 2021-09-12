package repository //主要處理ORM

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	. "webDemoBackend/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "admin:7754@tcp(192.168.0.18:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// db.Debug().Create(&BookDO)
	db.AutoMigrate(&BookDO{})

	// db.Create(&BookDO{ID: "16"})

	// var BookSlice []BookDO
	// db.Table("books").Find(&BookSlice)
	// fmt.Print(BookSlice, 1)
}

// func CreateData(db *gorm.DB) { //若已經新增則返回錯誤
// 	db.Create(&BookDO{Name: "bb", Price: 126})
// }

// func ReadData(db *gorm.DB) {
// 	db.First(&Book{}, 1) // 根据整形主键查找S
// 	//db.First(&Book{}, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
// }

// func UpdateData() {
// 	// Select with Struct (select zero value fields)
// db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
// // UPDATE users SET name='new_name', age=0 WHERE id=111;
// }

// func DeleteData(db *gorm.DB) {
// 	db.Delete(&Book{}, 10)
// }
