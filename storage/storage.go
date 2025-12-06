package storage

import (
	"coffee-tracker/models"
	"encoding/json"
	"fmt"
	"os"
)

type Store struct {
	filePath string
}

func NewStore() *Store {
	return &Store{
		filePath: "coffees.json",
	}
}

func (s *Store) Save(coffee *models.Coffee) error {
	coffees, _ := s.loadAll()
	coffees = append(coffees, coffee)
	return s.saveAll(coffees)
}

func (s *Store) GetByGUID(guid string) (*models.Coffee, error) {
	coffees, err := s.loadAll()
	if err != nil {
		return nil, err
	}

	for _, coffee := range coffees {
		if coffee.GUID == guid {
			return coffee, nil
		}
	}
	return nil, fmt.Errorf("coffee not found")
}

func (s *Store) GetByName(name string) ([]*models.Coffee, error) {
	coffees, err := s.loadAll()
	if err != nil {
		return nil, err
	}

	var results []*models.Coffee
	for _, coffee := range coffees {
		if coffee.Name == name {
			results = append(results, coffee)
		}
	}
	return results, nil
}

func (s *Store) List(limit int) ([]*models.Coffee, error) {
	coffees, err := s.loadAll()
	if err != nil {
		return nil, err
	}

	if limit > len(coffees) {
		limit = len(coffees)
	}
	return coffees[:limit], nil
}

func (s *Store) loadAll() ([]*models.Coffee, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []*models.Coffee{}, nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var coffees []*models.Coffee
	if err := json.Unmarshal(data, &coffees); err != nil {
		return nil, err
	}

	return coffees, nil
}

func (s *Store) saveAll(coffees []*models.Coffee) error {
	data, err := json.MarshalIndent(coffees, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
