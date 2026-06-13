package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins; tighten this in production
	},
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	hub := NewHub()

	mainPubSub := rdb.Subscribe(context.Background(), "ch:0")
	defer mainPubSub.Close()

	go func() {
		for message := range mainPubSub.Channel() {
			hub.BroadcastMessage(message.Payload)
		}
	}()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade error:", err)
			return
		}
		fmt.Println("ws client connected")

		client := &WSClient{conn: conn}
		hub.RegisterClient(client)

		go func() {
			defer func() {
				hub.UnregisterClient(client)
				conn.Close()
				fmt.Println("ws client disconnected")
			}()
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println("ws read error:", err)
					return
				}
				rdb.Publish(context.Background(), "ch:0", string(msg))
				fmt.Println("Publish message sent by ws client: " + string(msg))
			}
		}()
	})

	http.Handle("/", http.FileServer(http.Dir("./asset")))

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
