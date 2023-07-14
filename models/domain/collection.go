package domain

import "time"

type Collection struct {
	ID         string
	Name       string
	Visibility string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
