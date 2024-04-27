package brain

import (
	"github.com/akshay0074700747/chat-app/database/operations"
	"github.com/akshay0074700747/chat-app/entities"
)

type GlobalChatUsecase struct {
	Operations operations.GlobalChat
}

func NewGlobalChatUsecase(op operations.GlobalChat) *GlobalChatUsecase {
	return &GlobalChatUsecase{Operations: op}
}

func (global *GlobalChatUsecase) LoadInitialData() ([]entities.GlobalChat, error) {

	return global.Operations.FetchAll()

}

func (global *GlobalChatUsecase) AddMessage(req entities.GlobalChat) error  {
	
	return global.Operations.AddMessage(req)

}
