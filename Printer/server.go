package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Server struct to encapsulate the HTTP server logic
type Server struct {
	Port string
	so   *StackOverflow
}

func NewServer() *Server {
	return &Server{
		Port: "8080",
		so:   NewStackOverflow(),
	}
}

// Start method to start the server and keep it running
func (s *Server) Start() {
	http.HandleFunc("/", s.handleHelloWorld)

	http.HandleFunc("/add", s.AddQuestion)

	http.HandleFunc("/get", s.GetQuestion)

	http.HandleFunc("/getall", s.GetAllQuestions)

	fmt.Printf("Server is running on http://localhost:%s\n", s.Port)
	err := http.ListenAndServe(":"+s.Port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Handler method for the "/" endpoint
func (s *Server) handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func (s *Server) AddQuestion(w http.ResponseWriter, r *http.Request) {
	var q Question

	// Here you would typically parse the request body to fill the Question struct
	// For simplicity, let's assume q is filled with some data
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Ensure the body is closed after reading

	err = json.Unmarshal(body, &q)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	s.so.AddQuestion(q)
}

func (s *Server) GetQuestion(w http.ResponseWriter, r *http.Request) {
	var req GetQuestionRequest
	// Here you would typically parse the request body to fill the Question struct
	// For simplicity, let's assume q is filled with some data
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Ensure the body is closed after reading

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	q, found := s.so.GetQuestion(req.id)
	if found {
		response, err := json.Marshal(q)
		if err != nil {
			http.Error(w, "Error marshalling question", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		http.Error(w, "Question not found", http.StatusNotFound)
	}
}

func (s *Server) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions := s.so.GetAllQuestions()
	response, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Error marshalling questions", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
