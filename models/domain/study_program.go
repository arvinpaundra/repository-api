package domain

import "time"

type StudyProgram struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
