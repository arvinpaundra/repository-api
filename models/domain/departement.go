package domain

import "time"

type Departement struct {
	ID             string
	StudyProgramId string
	Name           string
	Code           string
	StudyProgram   StudyProgram
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
