package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("postgres", "sspzfyzexjagcv:32a3e10c18c74c343411ff1e1b8939e07ea73f168e9778506d6f21cffc0fb3d2@ec2-54-217-250-0.eu-west-1.compute.amazonaws.com:5432/d1cina0osrksr5")
	if err != nil {
		return
	}
	defer db.Close()
	DB = db
}
