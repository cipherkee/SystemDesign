package main

import (
	"github.com/gorilla/websocket"
)

type WSClient struct {
	conn *websocket.Conn
}

func (c *WSClient) Send(message string) error {
	return c.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

type Hub struct {
	clients map[*WSClient]bool
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*WSClient]bool),
	}
}

func (h *Hub) RegisterClient(client *WSClient) {
	h.clients[client] = true
}

func (h *Hub) UnregisterClient(client *WSClient) {
	delete(h.clients, client)
}

func (h *Hub) BroadcastMessage(message string) {
	for client := range h.clients {
		if err := client.Send(message); err != nil {
			delete(h.clients, client)
		}
	}
}
