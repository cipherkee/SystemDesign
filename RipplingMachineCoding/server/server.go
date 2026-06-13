package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cipherkee/ripplingmachinecoding/models"
	"github.com/cipherkee/ripplingmachinecoding/stackoverflow"
)

// --- Server Struct and Methods ---

type Server struct {
	svc stackoverflow.Stackoverflow
}

func (s *Server) handleGetWithBody(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var input models.Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response := models.Output{
		Message: fmt.Sprintf("Hello %s, you are %d years old (from JSON body)!", input.Name, input.Age),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	var input models.Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response := models.Output{Message: fmt.Sprintf("Received name: %s, age: %d", input.Name, input.Age)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) Start(addr string) {
	http.HandleFunc("/get-body", s.handleGetWithBody)
	http.HandleFunc("/post", s.handlePost)

	http.HandleFunc("/postquestion", s.handlePostQuestion)
	http.HandleFunc("/getquestions", s.handleGetQuestions)

	fmt.Println("Server starting on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	server := &Server{
		svc: *stackoverflow.NewStackOverflow(),
	}

	server.Start(":8080")
}

func (s *Server) handlePostQuestion(w http.ResponseWriter, r *http.Request) {
	var input models.QuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	s.svc.PostQuestion(&input)

	//response := models.Output{Message: fmt.Sprintf("Received name: %s, age: %d", input.Name, input.Age)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func (s *Server) handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	response := s.svc.GetQuestions()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
