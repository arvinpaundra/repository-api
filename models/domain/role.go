package domain

import "time"

type Role struct {
	ID         string
	Role       string
	Visibility string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
