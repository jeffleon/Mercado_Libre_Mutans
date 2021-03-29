package database

import (

	"fmt"
	"github.com/jeffleon/mutansApi/model"
	"github.com/jinzhu/gorm"
	_ "database/sql"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"


)

var (
	DBConn *gorm.DB
)

func InitDatabase(){

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "user", "password", "127.0.0.1:5432", "mutans")
	var err error
	if DBConn, err = gorm.Open("mysql", dsn); err != nil {
		panic(err)
	}
	fmt.Println("Database connection successfully openned")

	DBConn.AutoMigrate(&model.Mutants{})

	fmt.Println("Database Migrated")
}
