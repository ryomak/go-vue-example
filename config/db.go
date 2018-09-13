package config

import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func InitDB(conf DBConfig) *gorm.DB {
	param := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" +
		conf.Database + "?parseTime=true&loc=UTC"

	db, err := gorm.Open("mysql", param)
	if err != nil {
		panic(err)
	}
	return db
}
