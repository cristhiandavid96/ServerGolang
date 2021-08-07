package commons

import (
	"log"

	"github.com/cristhiandavid96/ServerGolang/models"

	"github.com/jinzhu/gorm"

	"github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/farmacia?charset=utf8")

	if error != nil {
		log.Fatal(error)
	}

	return db
}
