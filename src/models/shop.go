package models

type Shop struct {
	ID     string `json:"id"`
	Name  string `json:"name"`
	Introduction string `json:"introduction"`
	Location  string    `json:"location"`
	Level int `json:"level"`
	Score float64 `json:"score"`
}

func NewShop(id string, name string, introduction string, location string, level int, score float64) (*Shop, error) {
	return &Shop{
		ID:     id,
		Name:  name,
		Introduction: introduction,
		Location:  location,
		Level: level,
		Score: score,
	}, nil
}

func (s *Shop) Update(name string, introduction string, location string,level int, score float64) error {
	s.Name = name
	s.Introduction = introduction
	s.Location = location
	s.Level = level
	s.Score = score
	return nil
}