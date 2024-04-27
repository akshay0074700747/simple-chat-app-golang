package entities

import "time"

type GlobalChat struct {
	ID   uint   `gorm:"primaryKey;unique;not null"`
	User string `json:"from"`
	Message string `json:"message"`
	CreatedAt time.Time `gorm:"default:NOW()" json:"time"`
}

func (global *GlobalChat) Migrate_me() {

}
