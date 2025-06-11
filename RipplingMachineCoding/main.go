package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// --- Main Function with Client Code ---

func main() {
	server := &Server{}
	go server.Start(":8080")

	time.Sleep(1 * time.Second)

	// --- Client: GET with JSON body ---
	input := Input{Name: "Keerthana", Age: 30}
	jsonData, _ := json.Marshal(input)

	client := &http.Client{}
	HandleGet(client, jsonData)

	handlePut(client, jsonData)
}

func HandleGet(client *http.Client, jsonData []byte) {

	req, err := http.NewRequest("GET", "http://localhost:8080/get-body", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Failed to create GET request:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("GET with body failed:", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET Response:", string(body))
}

func handlePut(client *http.Client, jsonData []byte) {

	req, err := http.NewRequest("POST", "http://localhost:8080/post", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("PUT request failed:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	putResp, err := client.Do(req)
	if err != nil {
		log.Fatal("GET with body failed:", err)
	}

	body, _ := io.ReadAll(putResp.Body)
	fmt.Println("PUT Response:", string(body))
}
