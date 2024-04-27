package websockethandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/akshay0074700747/chat-app/brain"
	"github.com/akshay0074700747/chat-app/entities"
	"github.com/gorilla/websocket"
)

type GlobalChatHandler struct {
	Usecase     brain.GlobalChatUsecase
	Upgrader    websocket.Upgrader
	clients     map[*websocket.Conn]bool
	clientsLock sync.Mutex
}

func NewGlobalChatHandler(usecase brain.GlobalChatUsecase) *GlobalChatHandler {
	return &GlobalChatHandler{
		Usecase: usecase,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// Allow all origins
				return true
			},
		},
		clients: make(map[*websocket.Conn]bool),
	}
}

func (handler *GlobalChatHandler) GlobalChat(w http.ResponseWriter, r *http.Request) {
	conn, err := handler.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	handler.clientsLock.Lock()
	handler.clients[conn] = true
	handler.clientsLock.Unlock()

	defer func() {
		handler.clientsLock.Lock()
		delete(handler.clients, conn)
		handler.clientsLock.Unlock()

		conn.Close()
	}()

	// Send initial data to the connected client
	initialdata, err := handler.Usecase.LoadInitialData()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	for _, item := range initialdata {
		messageBytes, err := json.Marshal(item)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, messageBytes); err != nil {
			log.Println(err.Error())
			continue
		}
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				log.Println("WebSocket connection closed:", err)
				break
			}
			log.Println("WebSocket read error:", err)
			continue
		}

		log.Println("Received message:", string(message))

		var req entities.GlobalChat
		req.Message = string(message)

		if err := handler.Usecase.AddMessage(req); err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
			continue
		}

		handler.broadcast(message)
	}
}

func (handler *GlobalChatHandler) broadcast(message []byte) {
	handler.clientsLock.Lock()
	defer handler.clientsLock.Unlock()

	for conn := range handler.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("WebSocket write error:", err)
			conn.Close()
		}
	}
}
