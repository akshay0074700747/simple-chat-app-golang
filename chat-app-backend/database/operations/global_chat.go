package operations

import (
	"github.com/akshay0074700747/chat-app/entities"
	"gorm.io/gorm"
)

type GlobalChat struct {
	DB *gorm.DB
}

func NewGlobalChat(db *gorm.DB) *GlobalChat {
	return &GlobalChat{DB: db}
}

func (global *GlobalChat) FetchAll() ([]entities.GlobalChat, error) {

	var data []entities.GlobalChat

	query := `SELECT * FROM global_chats ORDER BY created_at ASC;`

	return data, global.DB.Raw(query).Scan(&data).Error

}

func (global *GlobalChat) AddMessage(req entities.GlobalChat) error {
	
	query := `INSERT INTO global_chats (message) VALUES($1);`

	return global.DB.Exec(query,req.Message).Error

}
