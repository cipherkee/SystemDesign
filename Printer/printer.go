package main

import "github.com/google/uuid"

type StackOverflow struct {
	qmap map[string]Question
}

func NewStackOverflow() *StackOverflow {
	return &StackOverflow{
		qmap: make(map[string]Question),
	}
}

func (so *StackOverflow) AddQuestion(q Question) {
	if _, exists := so.qmap[q.Title]; !exists {
		id := uuid.New().String()
		so.qmap[id] = q
	} else {
		// If the question already exists, you might want to update it or handle it differently
		// For now, we will just ignore the addition
		return
	}
}

func (so *StackOverflow) GetQuestion(title string) (Question, bool) {
	q, exists := so.qmap[title]
	return q, exists
}

func (so *StackOverflow) GetAllQuestions() []Question {
	questions := make([]Question, 0, len(so.qmap))
	for _, q := range so.qmap {
		questions = append(questions, q)
	}
	return questions
}
