package db

//Database connections with gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"w4s/models"
)

func SetupModels() *gorm.DB {

	if os.Getenv("BD-LOCATION") == "0" {
		data := "mysql"
		stringconection := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/w4s?charset=utf8&parseTime=True&loc=Local"
		if err := os.Setenv("db-STRING", stringconection); err != nil {
			panic(err)
		}
		if err := os.Setenv("db-PROGRAM", data); err != nil {
			panic(err)
		}
	} else {
		data := "postgres"
		stringconection := os.Getenv("DATABASE_URL")
		if err := os.Setenv("db-STRING", stringconection); err != nil {
			panic(err)
		}
		if err := os.Setenv("db-PROGRAM", data); err != nil {
			panic(err)
		}
	}
	db, err := gorm.Open(os.Getenv("db-PROGRAM"), os.Getenv("db-STRING"))
	//db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.UserAccountBadListToken{})
	db.AutoMigrate(&models.Table{})
	//db.AutoMigrate(&models.Picture{})

	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.TypeofTable{})

	db.AutoMigrate(&models.LogoffListTokens{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.PermissionTable{})

	return db
}
