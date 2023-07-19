package domain

import "time"

type StudyProgram struct {
	ID            string
	DepartementId string
	Name          string
	CoverColor    string
	Departement   Departement
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
