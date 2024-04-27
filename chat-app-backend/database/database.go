package database

import (
	"github.com/akshay0074700747/chat-app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Group_tables interface {
	Migrate_me()
}

func FireUp(databaseaddress string) (*gorm.DB,error) {

	var db *gorm.DB
	var err error

	db,err = gorm.Open(postgres.Open(databaseaddress),&gorm.Config{})

	if err != nil {
		return nil,err
	}

	Migrte_all(db,&entities.GlobalChat{})

	return db,err

}

func Migrte_all(dbconn *gorm.DB, models ...Group_tables) {
	for _, model := range models {
		dbconn.AutoMigrate(model)
	}
}
