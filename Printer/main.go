package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Create a new server instance
	server := NewServer()

	// Start the server in a separate goroutine to keep it running
	go server.Start()

	// Call the server endpoint from the main function
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error calling server endpoint:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response from the server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response from server:", string(body))

	q := Question{
		Title:       "Sample Question",
		Description: "This is a sample question for testing.",
		Star:        5,
	}

	AddQuestion(q)
	GetAllQuestions()
}

func AddQuestion(q Question) {
	qmarshal, err := json.Marshal(q)
	if err != nil {
		fmt.Println("Error marshaling question:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer(qmarshal)) // Example of adding a question
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	fmt.Println("Question added successfully, response status:", resp.Status)
	defer resp.Body.Close()
}

func GetAllQuestions() {
	resp, err := http.Get("http://localhost:8080/getall") // Example of getting all questions
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var questions []Question
	err = json.Unmarshal(body, &questions)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return
	}

	for _, q := range questions {
		fmt.Printf("Title: %s, Description: %s, Star: %d\n", q.Title, q.Description, q.Star)
	}
}
