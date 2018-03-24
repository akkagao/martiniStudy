package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var(
	DB *gorm.DB
)

func InitGorm() {
	db, err := gorm.Open("postgres", "host=localhost user=crazywolf dbname=crazywolf sslmode=disable password=crazywolf")
	DB = db
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	showTests()
	db.Exec("update test set name = 'gaojianwen' where id = 1")
	showTests()

}

func showTests() {
	var test struct {
		ID   int
		Name string
	}
	DB.Raw("select id,name from test where id = ?", 1).Scan(&test)
	fmt.Println(test)
}
