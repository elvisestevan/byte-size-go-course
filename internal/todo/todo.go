package todo

import (
	"fmt"
	"strings"
)

type Service struct {
	items []Item
}

type Item struct {
	Task   string
	Status string
}

func NewService() *Service {
	return &Service{
		items: make([]Item, 0),
	}
}

func (s *Service) Add(todo string) error {
	if todo == "" {
		return fmt.Errorf("todo item cannot be empty")
	}
	existing, _ := s.Get(todo)
	if existing.Task != "" {
		return fmt.Errorf("todo item already exists")
	}
	s.items = append(s.items, Item{
		Task:   todo,
		Status: "pending",
	})
	return nil
}

func (s *Service) Get(todo string) (Item, error) {
	for _, item := range s.items {
		if item.Task == todo {
			return item, nil
		}
	}
	return Item{}, fmt.Errorf("todo item not found")
}

func (s *Service) GetAll() []Item {
	return s.items
}

func (s *Service) Delete(todo string) error {
	for i, item := range s.items {
		if item.Task == todo {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo item not found")
}

func (s *Service) Search(term string) []Item {
	results := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		if strings.Contains(item.Task, term) {
			results = append(results, item)
		}
	}
	return results
}
