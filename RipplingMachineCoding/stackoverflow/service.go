package stackoverflow

import (
	"github.com/cipherkee/ripplingmachinecoding/models"
	"github.com/google/uuid"
)

var (
	users = map[string]*models.User{
		"1": &models.User{
			"abc@gmail.com",
		},
		"2": &models.User{
			"def@gmail.com",
		},
	}
)

type Storage struct {
	questions map[string]*models.Question
	users     map[string]*models.User
}

func NewStorage() *Storage {
	return &Storage{
		questions: map[string]*models.Question{},
		users:     map[string]*models.User{},
	}
}

type Stackoverflow struct {
	db *Storage
}

func NewStackOverflow() *Stackoverflow {
	return &Stackoverflow{
		db: NewStorage(),
	}
}

func (s *Stackoverflow) GetQuestions() map[string]*models.Question {
	return s.db.questions
}

func (s *Stackoverflow) PostQuestion(q *models.QuestionRequest) {
	id := uuid.New().String()

	s.db.questions[id] = &models.Question{
		Id:      id,
		Title:   q.Title,
		Content: q.Content,
		Creator: &models.User{
			EmailId: q.EmailId,
		},
	}
}
