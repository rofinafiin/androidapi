package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreateMariaGormConnection(dbConfig string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB.Statement.RaiseErrorOnNotFound = true

	if err != nil {
		panic(err)
	}
	return DB
}
