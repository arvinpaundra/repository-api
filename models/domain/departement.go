package domain

import "time"

type Departement struct {
	ID        string
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
