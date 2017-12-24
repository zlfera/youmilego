package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("postgres", "user=u0_a124 dbname=youmile_dev")
	if err != nil {
		return
	}
	defer db.Close()
	DB = db
}
