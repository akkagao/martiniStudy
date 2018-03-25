package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var (
	DB *gorm.DB
)

func InitGorm() {
	db, err := gorm.Open("postgres", "host=localhost user=crazywolf dbname=crazywolf sslmode=disable password=crazywolf")
	DB = db
	if err != nil {
		fmt.Println(err)
	}
	//defer db.Close()

	// 开启日志
	db.LogMode(true)
	//设置链接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// 测试数据库是否链接成功
	err = db.DB().Ping()
	if err != nil {
		panic("connect database fail...")
		os.Exit(1)
	}

	println("connect database success...")

	showTests()
	db.Exec("update t_user set name = 'gaojianwen_tset' where id = 1")
	showTests()

}

func showTests() {
	var test struct {
		ID   int
		Name string
	}
	DB.Raw("select id,name from t_user where id = ?", 1).Scan(&test)
	fmt.Println(test)
}
