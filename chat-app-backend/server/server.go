package server

import (
	"net/http"

	websockethandlers "github.com/akshay0074700747/chat-app/handlers/websocket_handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Engine struct {
	engine *mux.Router
}

func NewEngine(globalchathandler *websockethandlers.GlobalChatHandler) *Engine {

	r := mux.NewRouter()

	http.HandleFunc("/test", globalchathandler.GlobalChat)

	return &Engine{engine: r}
}

func (engine *Engine) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(engine.engine)
	http.Handle("/", handler)
	http.ListenAndServe(":3000", nil)
}
