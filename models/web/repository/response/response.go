package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type RepositoryResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	DateValidated string    `json:"date_validated"`
	Collection    string    `json:"collection"`
	Departement   string    `json:"departement"`
	Authors       []Author  `json:"authors"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DetailRepositoryResponse struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	Abstract      string        `json:"abstract"`
	Improvement   string        `json:"improvement"`
	RelatedTitle  string        `json:"related_title"`
	UpdateDesc    string        `json:"update_desc"`
	DateValidated string        `json:"date_validated"`
	Status        string        `json:"status"`
	Collection    string        `json:"collection"`
	Departement   string        `json:"departement"`
	Authors       []Author      `json:"authors"`
	Contributors  []Contributor `json:"contributors,omitempty"`
	Documents     Documents     `json:"documents,omitempty"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type Documents struct {
	ValidityPage        string `json:"validity_page"`
	CoverAndListContent string `json:"cover_and_list_content"`
	ChpOne              string `json:"chp_one"`
	ChpTwo              string `json:"chp_two"`
	ChpThree            string `json:"chp_three"`
	ChpFour             string `json:"chp_four"`
	ChpFive             string `json:"chp_five"`
	Bibliography        string `json:"bibliography"`
}

type Author struct {
	AuthorId    string `json:"author_id"`
	PemustakaId string `json:"pemustaka_id"`
	Fullname    string `json:"fullname"`
}

type Contributor struct {
	ContributorId string `json:"contributor_id"`
	PemustakaId   string `json:"pemustaka_id"`
	Fullname      string `json:"fullname"`
	ContributedAs string `json:"contributed_as"`
}

func ToRepositoryDocumentsResponse(documents domain.Document) Documents {
	return Documents{
		ValidityPage:        documents.ValidityPage,
		CoverAndListContent: documents.CoverAndListContent,
		ChpOne:              documents.ChpOne,
		ChpTwo:              documents.ChpTwo,
		ChpThree:            documents.ChpThree,
		ChpFour:             documents.ChpFour,
		ChpFive:             documents.ChpFive,
		Bibliography:        documents.Bibliography,
	}
}

func ToArrayAuthorResponse(authorsFromDomain []domain.Author) []Author {
	var authors []Author

	for _, author := range authorsFromDomain {
		authors = append(authors, Author{
			AuthorId:    author.ID,
			PemustakaId: author.Pemustaka.ID,
			Fullname:    author.Pemustaka.Fullname,
		})
	}

	return authors
}

func ToArrayContributorResponse(contributorsFromDomain []domain.Contributor) []Contributor {
	var contributors []Contributor

	for _, contributor := range contributorsFromDomain {
		contributors = append(contributors, Contributor{
			ContributorId: contributor.ID,
			PemustakaId:   contributor.PemustakaId,
			Fullname:      contributor.Pemustaka.Fullname,
			ContributedAs: contributor.ContributedAs,
		})
	}

	return contributors
}

func ToRepositoryResponse(repository domain.Repository, authors []Author, contributors []Contributor, documents Documents) DetailRepositoryResponse {
	return DetailRepositoryResponse{
		ID:            repository.ID,
		Title:         repository.Title,
		Abstract:      repository.Abstract,
		Improvement:   repository.Improvement,
		RelatedTitle:  repository.RelatedTitle,
		UpdateDesc:    repository.UpdateDesc,
		DateValidated: repository.DateValidated,
		Status:        repository.Status,
		Collection:    repository.Collection.Name,
		Departement:   repository.Departement.Name,
		Authors:       authors,
		Contributors:  contributors,
		Documents:     documents,
		CreatedAt:     repository.CreatedAt,
		UpdatedAt:     repository.UpdatedAt,
	}
}

func ToRepositoriesResponse(repositoriesFromDomain []domain.Repository) []RepositoryResponse {
	var repositories []RepositoryResponse

	for _, repository := range repositoriesFromDomain {
		repositories = append(repositories, RepositoryResponse{
			ID:            repository.ID,
			Title:         repository.Title,
			DateValidated: repository.DateValidated,
			Collection:    repository.Collection.Name,
			Departement:   repository.Departement.Name,
			Authors:       ToArrayAuthorResponse(repository.Authors),
			CreatedAt:     repository.CreatedAt,
			UpdatedAt:     repository.UpdatedAt,
		})
	}

	return repositories
}
