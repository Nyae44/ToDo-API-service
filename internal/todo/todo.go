package todo

import (
	"context"
	"errors"
	"fmt"
	"my-first-api/internal/db"
	"strings"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) Add(todo string) error {
	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("failed to read from db: %w", err)
	}
	for _, t := range items {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}
	return nil
}
func (svc *Service) Search(query string) ([]string, error) {
	items, err := svc.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read from db: %w", err)
	}

	var results []string
	for _, todo := range items {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			results = append(results, todo.Task)
		}
	}
	return results, nil
}
func (svc *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read from db: %w", err)
	}
	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}
	return results, nil
}
