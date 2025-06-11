package main

type Question struct {
	Title       string
	Description string
	Star        int
}

type GetQuestionRequest struct {
	id string
}
