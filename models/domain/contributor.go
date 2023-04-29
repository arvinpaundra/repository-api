package domain

import "time"

type Contributor struct {
	ID            string
	RepositoryId  string
	PemustakaId   string
	ContributedAs string
	Pemustaka     Pemustaka
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
