package domain

import "time"

type Staff struct {
	ID        string
	UserId    string
	RoleId    string
	Fullname  string
	Nip       string
	Telp      string
	Address   string
	Gender    string
	BirthDate string
	IsActive  string
	Avatar    string
	Signature string
	User      User
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
