package domain

import "time"

type RequestAccess struct {
	ID              string
	PemustakaId     string
	SupportEvidence string
	Status          string
	Reasons         string
	Pemustaka       Pemustaka
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
