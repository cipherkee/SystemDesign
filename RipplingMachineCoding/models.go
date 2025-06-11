package main

// --- Data Types ---

type Input struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Output struct {
	Message string `json:"message"`
}
