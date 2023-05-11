package repository

import (
	"context"
	"math"
	"strings"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/drivers/mysql/author"
	"github.com/arvinpaundra/repository-api/drivers/mysql/collection"
	"github.com/arvinpaundra/repository-api/drivers/mysql/contributor"
	"github.com/arvinpaundra/repository-api/drivers/mysql/departement"
	"github.com/arvinpaundra/repository-api/drivers/mysql/document"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/repository"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/repository/request"
	"github.com/arvinpaundra/repository-api/models/web/repository/response"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepositoryServiceImpl struct {
	collectionRepository  collection.CollectionRepository
	departementRepository departement.DepartementRepository
	pemustakaRepository   pemustaka.PemustakaRepository
	authorRepository      author.AuthorRepository
	contributorRepository contributor.ContributorRepository
	repository            repository.Repository
	documentRepository    document.DocumentRepository
	cloudinary            cloudinary.Cloudinary
	tx                    *gorm.DB
}

func NewRepositoryService(
	collectionRepository collection.CollectionRepository,
	departementRepository departement.DepartementRepository,
	pemustakaRepository pemustaka.PemustakaRepository,
	authorRepository author.AuthorRepository,
	contributorRepository contributor.ContributorRepository,
	repository repository.Repository,
	documentRepository document.DocumentRepository,
	cloudinary cloudinary.Cloudinary,
	tx *gorm.DB,
) RepositoryService {
	return RepositoryServiceImpl{
		collectionRepository:  collectionRepository,
		departementRepository: departementRepository,
		pemustakaRepository:   pemustakaRepository,
		authorRepository:      authorRepository,
		contributorRepository: contributorRepository,
		repository:            repository,
		documentRepository:    documentRepository,
		cloudinary:            cloudinary,
		tx:                    tx,
	}
}

func (service RepositoryServiceImpl) CreateFinalProjectReport(ctx context.Context, repositoryDTO request.CreateFinalProjectReportRequest, files request.RepositoryInputFiles) error {
	tx := service.tx.Begin()

	if _, err := service.collectionRepository.FindById(ctx, configs.GetConfig("ID_FINAL_PROJECT")); err != nil {
		return err
	}

	if _, err := service.departementRepository.FindById(ctx, repositoryDTO.DepartementId); err != nil {
		return err
	}

	pemustaka, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.Author)

	if err != nil {
		return err
	}

	if pemustaka.IsCollectedFinalProject == "1" {
		return utils.ErrAlreadyCollectedFinalProject
	}

	if _, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.FirstMentor); err != nil {
		return utils.ErrMentorNotFound
	}

	if _, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.SecondMentor); err != nil {
		return utils.ErrMentorNotFound
	}

	if _, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.FirstExaminer); err != nil {
		return utils.ErrExaminerNotFound
	}

	if _, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.SecondExaminer); err != nil {
		return utils.ErrExaminerNotFound
	}

	validityPageURL, err := service.cloudinary.Upload(ctx, "validity-pages", utils.GetFilename(), files.ValidityPage)

	if err != nil {
		return err
	}

	coverAndListContentURL, err := service.cloudinary.Upload(ctx, "covers", utils.GetFilename(), files.CoverAndListContent)

	if err != nil {
		return err
	}

	chpOneURL, err := service.cloudinary.Upload(ctx, "bab1", utils.GetFilename(), files.ChpOne)

	if err != nil {
		return err
	}

	chpTwoURL, err := service.cloudinary.Upload(ctx, "bab2", utils.GetFilename(), files.ChpTwo)

	if err != nil {
		return err
	}

	chpThreeURL, err := service.cloudinary.Upload(ctx, "bab3", utils.GetFilename(), files.ChpThree)

	if err != nil {
		return err
	}

	chpFourURL, err := service.cloudinary.Upload(ctx, "bab4", utils.GetFilename(), files.ChpFour)

	if err != nil {
		return err
	}

	chpFiveURL, err := service.cloudinary.Upload(ctx, "bab5", utils.GetFilename(), files.ChpFive)

	if err != nil {
		return err
	}

	bibliographyURL, err := service.cloudinary.Upload(ctx, "bibliographies", utils.GetFilename(), files.Bibliography)

	if err != nil {
		return err
	}

	repositoryDomain := domain.Repository{
		ID:            uuid.NewString(),
		CollectionId:  configs.GetConfig("ID_FINAL_PROJECT"),
		DepartementId: repositoryDTO.DepartementId,
		Title:         repositoryDTO.Title,
		Abstract:      repositoryDTO.Abstract,
		DateValidated: repositoryDTO.DateValidated,
		Improvement:   repositoryDTO.Improvement,
		RelatedTitle:  repositoryDTO.RelatedTitle,
		UpdateDesc:    repositoryDTO.UpdateDesc,
		Status:        "pending",
	}

	if err := service.repository.Save(ctx, tx, repositoryDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	authorDomain := domain.Author{
		ID:           uuid.NewString(),
		RepositoryId: repositoryDomain.ID,
		PemustakaId:  repositoryDTO.Author,
	}

	if err := service.authorRepository.Save(ctx, tx, []domain.Author{authorDomain}); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	firstMentor := domain.Contributor{
		ID:            uuid.NewString(),
		RepositoryId:  repositoryDomain.ID,
		PemustakaId:   repositoryDTO.FirstMentor,
		ContributedAs: "Pembimbing 1",
	}

	if err := service.contributorRepository.Save(ctx, tx, firstMentor); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	secondMentor := domain.Contributor{
		ID:            uuid.NewString(),
		RepositoryId:  repositoryDomain.ID,
		PemustakaId:   repositoryDTO.SecondMentor,
		ContributedAs: "Pembimbing 2",
	}

	if err := service.contributorRepository.Save(ctx, tx, secondMentor); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	firstExaminer := domain.Contributor{
		ID:            uuid.NewString(),
		RepositoryId:  repositoryDomain.ID,
		PemustakaId:   repositoryDTO.FirstExaminer,
		ContributedAs: "Penguji 1",
	}

	if err := service.contributorRepository.Save(ctx, tx, firstExaminer); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	secondExaminer := domain.Contributor{
		ID:            uuid.NewString(),
		RepositoryId:  repositoryDomain.ID,
		PemustakaId:   repositoryDTO.SecondExaminer,
		ContributedAs: "Penguji 2",
	}

	if err := service.contributorRepository.Save(ctx, tx, secondExaminer); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	documentDomain := domain.Document{
		ID:                  uuid.NewString(),
		RepositoryId:        repositoryDomain.ID,
		ValidityPage:        validityPageURL,
		CoverAndListContent: coverAndListContentURL,
		ChpOne:              chpOneURL,
		ChpTwo:              chpTwoURL,
		ChpThree:            chpThreeURL,
		ChpFour:             chpFourURL,
		ChpFive:             chpFiveURL,
		Bibliography:        bibliographyURL,
	}

	if err := service.documentRepository.Save(ctx, tx, documentDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	pemustakaDomain := domain.Pemustaka{
		IsCollectedFinalProject: "1",
	}

	if err := service.pemustakaRepository.Update(ctx, tx, pemustakaDomain, repositoryDTO.Author); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}

func (service RepositoryServiceImpl) CreateInternshipReport(ctx context.Context, repositoryDTO request.CreateInternshipReportRequest, files request.RepositoryInputFiles) error {
	tx := service.tx.Begin()

	if _, err := service.collectionRepository.FindById(ctx, configs.GetConfig("ID_INTERNSHIP_REPORT")); err != nil {
		return err
	}

	if _, err := service.departementRepository.FindById(ctx, repositoryDTO.DepartementId); err != nil {
		return err
	}

	pemustaka, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.Author)

	if err != nil {
		return err
	}

	if pemustaka.IsCollectedInternshipReport == "1" {
		return utils.ErrAlreadyCollectedInternshipReport
	}

	if _, err := service.pemustakaRepository.FindById(ctx, repositoryDTO.Mentor); err != nil {
		return err
	}

	validityPageURL, err := service.cloudinary.Upload(ctx, "validity-pages", utils.GetFilename(), files.ValidityPage)

	if err != nil {
		return err
	}

	coverAndListContentURL, err := service.cloudinary.Upload(ctx, "covers", utils.GetFilename(), files.CoverAndListContent)

	if err != nil {
		return err
	}

	chpOneURL, err := service.cloudinary.Upload(ctx, "bab1", utils.GetFilename(), files.ChpOne)

	if err != nil {
		return err
	}

	chpTwoURL, err := service.cloudinary.Upload(ctx, "bab2", utils.GetFilename(), files.ChpTwo)

	if err != nil {
		return err
	}

	chpThreeURL, err := service.cloudinary.Upload(ctx, "bab3", utils.GetFilename(), files.ChpThree)

	if err != nil {
		return err
	}

	chpFourURL, err := service.cloudinary.Upload(ctx, "bab4", utils.GetFilename(), files.ChpFour)

	if err != nil {
		return err
	}

	chpFiveURL, err := service.cloudinary.Upload(ctx, "bab5", utils.GetFilename(), files.ChpFive)

	if err != nil {
		return err
	}

	bibliographyURL, err := service.cloudinary.Upload(ctx, "bibliographies", utils.GetFilename(), files.Bibliography)

	if err != nil {
		return err
	}

	repositoryDomain := domain.Repository{
		ID:            uuid.NewString(),
		CollectionId:  configs.GetConfig("ID_INTERNSHIP_REPORT"),
		DepartementId: repositoryDTO.DepartementId,
		Title:         repositoryDTO.Title,
		Improvement:   "0",
		DateValidated: repositoryDTO.DateValidated,
		Status:        "pending",
	}

	if err := service.repository.Save(ctx, tx, repositoryDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	authorDomain := domain.Author{
		ID:           uuid.NewString(),
		RepositoryId: repositoryDomain.ID,
		PemustakaId:  repositoryDTO.Author,
	}

	if err := service.authorRepository.Save(ctx, tx, []domain.Author{authorDomain}); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	contributorDomain := domain.Contributor{
		ID:            uuid.NewString(),
		RepositoryId:  repositoryDomain.ID,
		PemustakaId:   repositoryDTO.Mentor,
		ContributedAs: "Pembimbing Magang",
	}

	if err := service.contributorRepository.Save(ctx, tx, contributorDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	documentDomain := domain.Document{
		ID:                  uuid.NewString(),
		RepositoryId:        repositoryDomain.ID,
		ValidityPage:        validityPageURL,
		CoverAndListContent: coverAndListContentURL,
		ChpOne:              chpOneURL,
		ChpTwo:              chpTwoURL,
		ChpThree:            chpThreeURL,
		ChpFour:             chpFourURL,
		ChpFive:             chpFiveURL,
		Bibliography:        bibliographyURL,
	}

	if err := service.documentRepository.Save(ctx, tx, documentDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	pemustakaDomain := domain.Pemustaka{
		IsCollectedInternshipReport: "1",
	}

	if err := service.pemustakaRepository.Update(ctx, tx, pemustakaDomain, repositoryDTO.Author); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}

func (service RepositoryServiceImpl) CreateResearchReport(ctx context.Context, repositoryDTO request.CreateResearchReportRequest, files request.RepositoryInputFiles) error {
	tx := service.tx.Begin()

	if _, err := service.collectionRepository.FindById(ctx, repositoryDTO.CollectionId); err != nil {
		return err
	}

	if _, err := service.departementRepository.FindById(ctx, repositoryDTO.DepartementId); err != nil {
		return err
	}

	authors := strings.Split(repositoryDTO.Authors[0], ",")

	for i := range authors {
		if _, err := service.pemustakaRepository.FindById(ctx, authors[i]); err != nil {
			return err
		}
	}

	validityPageURL, err := service.cloudinary.Upload(ctx, "validity-pages", utils.GetFilename(), files.ValidityPage)

	if err != nil {
		return err
	}

	coverAndListContentURL, err := service.cloudinary.Upload(ctx, "covers", utils.GetFilename(), files.CoverAndListContent)

	if err != nil {
		return err
	}

	chpOneURL, err := service.cloudinary.Upload(ctx, "bab1", utils.GetFilename(), files.ChpOne)

	if err != nil {
		return err
	}

	chpTwoURL, err := service.cloudinary.Upload(ctx, "bab2", utils.GetFilename(), files.ChpTwo)

	if err != nil {
		return err
	}

	chpThreeURL, err := service.cloudinary.Upload(ctx, "bab3", utils.GetFilename(), files.ChpThree)

	if err != nil {
		return err
	}

	chpFourURL, err := service.cloudinary.Upload(ctx, "bab4", utils.GetFilename(), files.ChpFour)

	if err != nil {
		return err
	}

	chpFiveURL, err := service.cloudinary.Upload(ctx, "bab5", utils.GetFilename(), files.ChpFive)

	if err != nil {
		return err
	}

	bibliographyURL, err := service.cloudinary.Upload(ctx, "bibliographies", utils.GetFilename(), files.Bibliography)

	if err != nil {
		return err
	}

	repositoryDomain := domain.Repository{
		ID:            uuid.NewString(),
		CollectionId:  repositoryDTO.CollectionId,
		DepartementId: repositoryDTO.DepartementId,
		Title:         repositoryDTO.Title,
		Abstract:      repositoryDTO.Abstract,
		DateValidated: repositoryDTO.DateValidated,
		Improvement:   "0",
		Status:        "pending",
	}

	if err := service.repository.Save(ctx, tx, repositoryDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	authorsDomain := make([]domain.Author, 0)

	for i := range authors {
		authorsDomain = append(authorsDomain, domain.Author{
			ID:           uuid.NewString(),
			RepositoryId: repositoryDomain.ID,
			PemustakaId:  authors[i],
		})
	}

	if err := service.authorRepository.Save(ctx, tx, authorsDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	documentDomain := domain.Document{
		ID:                  uuid.NewString(),
		RepositoryId:        repositoryDomain.ID,
		ValidityPage:        validityPageURL,
		CoverAndListContent: coverAndListContentURL,
		ChpOne:              chpOneURL,
		ChpTwo:              chpTwoURL,
		ChpThree:            chpThreeURL,
		ChpFour:             chpFourURL,
		ChpFive:             chpFiveURL,
		Bibliography:        bibliographyURL,
	}

	if err := service.documentRepository.Save(ctx, tx, documentDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}

func (service RepositoryServiceImpl) Delete(ctx context.Context, repositoryId string) error {
	repository, err := service.repository.FindById(ctx, repositoryId)

	if err != nil {
		return err
	}

	document, err := service.documentRepository.FindByRepositoryId(ctx, repository.ID)

	if err != nil {
		return err
	}

	if err := service.repository.Delete(ctx, repositoryId); err != nil {
		return err
	}

	// delete files from cloud
	if err := service.cloudinary.Delete(ctx, document.ValidityPage); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.CoverAndListContent); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.ChpOne); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.ChpTwo); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.ChpThree); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.ChpFour); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.ChpFive); err != nil {
		return err
	}

	if err := service.cloudinary.Delete(ctx, document.Bibliography); err != nil {
		return err
	}

	return nil
}

func (service RepositoryServiceImpl) FindAll(ctx context.Context, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error) {
	repositories, totalRows, err := service.repository.FindAll(ctx, query, limit, offset)

	if err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRepositoriesResponse(repositories), totalRows, int(totalPages), nil
}

func (service RepositoryServiceImpl) FindById(ctx context.Context, repositoryId string) (response.DetailRepositoryResponse, error) {
	repositoryDomain, err := service.repository.FindById(ctx, repositoryId)

	if err != nil {
		return response.DetailRepositoryResponse{}, err
	}

	authorsDomain, err := service.authorRepository.FindByRepositoryId(ctx, repositoryId)

	if err != nil {
		return response.DetailRepositoryResponse{}, err
	}

	contributorsDomain, err := service.contributorRepository.FindByRepositoryId(ctx, repositoryId)

	if err != nil {
		return response.DetailRepositoryResponse{}, err
	}

	documentsDomain, err := service.documentRepository.FindByRepositoryId(ctx, repositoryId)

	if err != nil {
		return response.DetailRepositoryResponse{}, utils.ErrDocumentsNotFound
	}

	authors := response.ToArrayAuthorResponse(authorsDomain)
	contributors := response.ToArrayContributorResponse(contributorsDomain)
	documents := response.ToRepositoryDocumentsResponse(documentsDomain)

	return response.ToRepositoryResponse(repositoryDomain, authors, contributors, documents), nil
}

func (service RepositoryServiceImpl) FindByAuthorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error) {
	if _, err := service.pemustakaRepository.FindById(ctx, pemustakaId); err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	repositories, totalRows, err := service.repository.FindByAuthorId(ctx, pemustakaId, query, limit, offset)

	if err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRepositoriesResponse(repositories), totalRows, int(totalPages), nil
}

func (service RepositoryServiceImpl) FindByMentorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error) {
	if _, err := service.pemustakaRepository.FindById(ctx, pemustakaId); err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	repositories, totalRows, err := service.repository.FindByMentorId(ctx, pemustakaId, query, limit, offset)

	if err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRepositoriesResponse(repositories), totalRows, int(totalPages), nil
}

func (service RepositoryServiceImpl) FindByExaminerId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error) {
	if _, err := service.pemustakaRepository.FindById(ctx, pemustakaId); err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	repositories, totalRows, err := service.repository.FindByExaminerId(ctx, pemustakaId, query, limit, offset)

	if err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRepositoriesResponse(repositories), totalRows, int(totalPages), nil
}

func (service RepositoryServiceImpl) FindByCollectionId(ctx context.Context, collectionId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error) {
	if _, err := service.collectionRepository.FindById(ctx, collectionId); err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	repositories, totalRows, err := service.repository.FindByCollectionId(ctx, collectionId, query, limit, offset)

	if err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRepositoriesResponse(repositories), totalRows, int(totalPages), nil
}

func (service RepositoryServiceImpl) FindByDepartementId(ctx context.Context, departementId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error) {
	if _, err := service.departementRepository.FindById(ctx, departementId); err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	repositories, totalRows, err := service.repository.FindByDepartementId(ctx, departementId, query, limit, offset)

	if err != nil {
		return []response.RepositoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRepositoriesResponse(repositories), totalRows, int(totalPages), nil
}

func (service RepositoryServiceImpl) GetTotal(ctx context.Context, status string) (int, error) {
	total, err := service.repository.GetTotal(ctx, status)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (service RepositoryServiceImpl) Confirm(ctx context.Context, req request.ConfirmRequest, repositoryId string) error {
	tx := service.tx.Begin()

	if _, err := service.repository.FindById(ctx, repositoryId); err != nil {
		return err
	}

	repositoryDomain := domain.Repository{
		Status: req.Status,
	}

	if err := service.repository.Update(ctx, tx, repositoryDomain, repositoryId); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}
