package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"trivium/internal/domain/repositorier"
	presentation_repositorier "trivium/internal/presentation/repositorier"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // CORS handled by middleware
	},
}

type WsCryptoController struct {
	cryptoHistoryRepo repositorier.CryptoHistoryRepository
	clients           map[*websocket.Conn]bool
	mu                sync.RWMutex
	broadcast         chan interface{}
}

func NewWsCryptoController(cryptoHistoryRepo repositorier.CryptoHistoryRepository) *WsCryptoController {
	ctrl := &WsCryptoController{
		cryptoHistoryRepo: cryptoHistoryRepo,
		clients:           make(map[*websocket.Conn]bool),
		broadcast:         make(chan interface{}, 100),
	}

	go ctrl.handleBroadcast()
	return ctrl
}

func (c *WsCryptoController) SetupRoutes(router presentation_repositorier.HttpRepositorier) {
	router.HandleFunc("/ws/crypto", c.HandleWebSocket, http.MethodGet)
}

func (c *WsCryptoController) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	c.mu.Lock()
	c.clients[conn] = true
	c.mu.Unlock()

	log.Printf("New WebSocket client connected. Total: %d", len(c.clients))

	defer func() {
		c.mu.Lock()
		delete(c.clients, conn)
		c.mu.Unlock()
		conn.Close()
		log.Printf("WebSocket client disconnected. Total: %d", len(c.clients))
	}()

	// Keep connection alive, read messages (ping/pong)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (c *WsCryptoController) BroadcastPrice(data interface{}) {
	select {
	case c.broadcast <- data:
	default:
		// channel full, skip
	}
}

func (c *WsCryptoController) handleBroadcast() {
	for msg := range c.broadcast {
		jsonData, err := json.Marshal(msg)
		if err != nil {
			continue
		}

		c.mu.RLock()
		for client := range c.clients {
			err := client.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				client.Close()
				go func(conn *websocket.Conn) {
					c.mu.Lock()
					delete(c.clients, conn)
					c.mu.Unlock()
				}(client)
			}
		}
		c.mu.RUnlock()
	}
}
