package models

// --- Data Types ---

type Input struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Output struct {
	Message string `json:"message"`
}

// --- Data Types ---

type QuestionRequest struct {
	Title   string
	Content string
	EmailId string
}

type Question struct {
	Id      string
	Title   string
	Content string
	Creator *User
}

type User struct {
	EmailId string
}
