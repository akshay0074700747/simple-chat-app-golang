package provider

import (
	"github.com/akshay0074700747/chat-app/brain"
	"github.com/akshay0074700747/chat-app/database/operations"
	websockethandlers "github.com/akshay0074700747/chat-app/handlers/websocket_handlers"
	"github.com/akshay0074700747/chat-app/server"
	"gorm.io/gorm"
)

func Provide(db *gorm.DB) *server.Engine  {

	globalchatop := operations.NewGlobalChat(db)
	globalchatuse := brain.NewGlobalChatUsecase(*globalchatop)
	globalchathandle := websockethandlers.NewGlobalChatHandler(*globalchatuse)



	engine := server.NewEngine(globalchathandle)
	return engine
}