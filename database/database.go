package database

import (
	"fmt"
	setting "github.com/moyrne/tractor/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDBEngine(s *setting.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		s.Username,
		s.Password,
		s.Host,
		s.DBName,
		s.Charset,
		s.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Operation interface {
	TableName() string
	Get(tx *gorm.DB) error
	Insert(tx *gorm.DB) error
	Update(tx *gorm.DB, values map[string]interface{}) error
	Delete(tx *gorm.DB) error
}
