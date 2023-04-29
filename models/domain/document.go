package domain

import "time"

type Document struct {
	ID                  string
	RepositoryId        string
	ValidityPage        string
	CoverAndListContent string
	ChpOne              string
	ChpTwo              string
	ChpThree            string
	ChpFour             string
	ChpFive             string
	Bibliography        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
