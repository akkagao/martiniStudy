package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitGorm() {
	db, err := gorm.Open("postgres", "host=localhost user=crazywolf dbname=crazywolf sslmode=disable password=crazywolf")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var test struct{
		ID int
		Name string
	}
	db.Raw("select id,name from test where id = ?",1).Scan(&test)
	fmt.Println(test)
}
