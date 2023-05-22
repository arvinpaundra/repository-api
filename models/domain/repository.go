package domain

import "time"

type Repository struct {
	ID            string
	CollectionId  string
	DepartementId string
	CategoryId    string
	Title         string
	Abstract      string
	Improvement   string
	RelatedTitle  string
	UpdateDesc    string
	DateValidated string
	Status        string
	Collection    Collection
	Departement   Departement
	Category      Category
	Documents     Document
	Authors       []Author
	Contributors  []Contributor
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
