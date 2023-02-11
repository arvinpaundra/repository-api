package domain

import "time"

type Pemustaka struct {
	ID                      string
	UserId                  string
	StudyProgramId          string
	DepartementId           string
	RoleId                  string
	MemberCode              string
	Fullname                string
	IdentityNumber          string
	YearGen                 string
	Gender                  string
	Telp                    string
	BirthDate               string
	Address                 string
	IsCollectedFinalProject string
	IsActive                string
	Avatar                  string
	User                    User
	StudyProgram            StudyProgram
	Departement             Departement
	Role                    Role
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
