package domain

import "time"

type Author struct {
	ID           string
	RepositoryId string
	PemustakaId  string
	Pemustaka    Pemustaka
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
