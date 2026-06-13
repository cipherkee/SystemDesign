package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cipherkee/ripplingmachinecoding/models"
)

const (
	getSample     = "http://localhost:8080/get-body"
	postSample    = "http://localhost:8080/post"
	getQuestions  = "http://localhost:8080/getquestions"
	postquestions = "http://localhost:8080/postquestion"
)

// --- Main Function with Client Code ---

func main() {

	time.Sleep(1 * time.Second)

	// --- Client: GET with JSON body ---
	// input := models.Input{Name: "Keerthana", Age: 30}
	// jsonData, _ := json.Marshal(input)

	client := &http.Client{}

	// HandleGet(client, jsonData)

	// handlePut(client, jsonData)

	question := models.QuestionRequest{
		Title:   "Question1",
		Content: "How are you ?",
		EmailId: "abc@gmail.com",
	}

	jsonData, _ := json.Marshal(question)

	question2 := models.QuestionRequest{
		Title:   "Question2",
		Content: "What is it ?",
		EmailId: "def@gmail.com",
	}

	jsonData2, _ := json.Marshal(question2)

	handlePut(client, postquestions, jsonData)
	handlePut(client, postquestions, jsonData2)
	HandleGet(client, getQuestions)
}

func HandleGet(client *http.Client, url string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to create GET request:", err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET Response:", string(body))
}

func handlePut(client *http.Client, url string, jsonData []byte) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("PUT request failed:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	putResp, err := client.Do(req)
	if err != nil {
		log.Fatal("GET with body failed:", err)
	}

	body, _ := io.ReadAll(putResp.Body)
	fmt.Println("POST Response:", string(body))
}
