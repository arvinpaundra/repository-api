package pemustaka

import (
	"context"
	"math"
	"mime/multipart"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/drivers/mysql/departement"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	requestAccess "github.com/arvinpaundra/repository-api/drivers/mysql/requestAccess"
	"github.com/arvinpaundra/repository-api/drivers/mysql/role"
	studyProgram "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram"
	"github.com/arvinpaundra/repository-api/drivers/mysql/user"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/response"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PemustakaServiceImpl struct {
	userRepository          user.UserRepository
	pemustakaRepository     pemustaka.PemustakaRepository
	studyProgramRepository  studyProgram.StudyProgramRepository
	departementRepository   departement.DepartementRepository
	roleRepository          role.RoleRepository
	requestAccessRepository requestAccess.RequestAccessRepository
	cloudinary              cloudinary.Cloudinary
	tx                      *gorm.DB
}

func NewPemustakaService(
	userRepository user.UserRepository,
	pemustakaRepository pemustaka.PemustakaRepository,
	studyProgramRepository studyProgram.StudyProgramRepository,
	departementRepository departement.DepartementRepository,
	roleRepository role.RoleRepository,
	requestAccessRepository requestAccess.RequestAccessRepository,
	cloudinary cloudinary.Cloudinary,
	tx *gorm.DB,
) PemustakaService {
	return PemustakaServiceImpl{
		userRepository:          userRepository,
		pemustakaRepository:     pemustakaRepository,
		studyProgramRepository:  studyProgramRepository,
		departementRepository:   departementRepository,
		roleRepository:          roleRepository,
		requestAccessRepository: requestAccessRepository,
		cloudinary:              cloudinary,
		tx:                      tx,
	}
}

func (service PemustakaServiceImpl) Register(ctx context.Context, req request.RegisterPemustakaRequest, supportEvidence *multipart.FileHeader) error {
	tx := service.tx.Begin()

	user, _ := service.userRepository.FindByEmail(ctx, req.Email)

	if user.Email != "" {
		return utils.ErrEmailAlreadyUsed
	}

	if _, err := service.studyProgramRepository.FindById(ctx, req.StudyProgramId); err != nil {
		return err
	}

	departement, err := service.departementRepository.FindById(ctx, req.DepartementId)

	if err != nil {
		return err
	}

	if _, err := service.roleRepository.FindById(ctx, req.RoleId); err != nil {
		return err
	}

	hashPassword := utils.HashPassword(req.Password)

	var userDomain domain.User
	userDomain.ID = uuid.NewString()
	userDomain.Email = req.Email
	userDomain.Password = hashPassword

	if err := service.userRepository.Save(ctx, tx, userDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	// get member code
	totalPemustaka, err := service.pemustakaRepository.GetTotalPemustakaByDepartementId(ctx, req.DepartementId)

	if err != nil {
		return err
	}

	memberCode := utils.GetMemberCode(totalPemustaka, departement.Code)

	pemustakaDomain := req.ToDomainPemustaka()
	pemustakaDomain.ID = uuid.NewString()
	pemustakaDomain.UserId = userDomain.ID
	pemustakaDomain.MemberCode = memberCode
	pemustakaDomain.IsActive = "0"
	pemustakaDomain.IsCollectedFinalProject = "0"
	pemustakaDomain.Avatar = configs.GetConfig("DEFAULT_AVATAR")

	filename := utils.GetFilename()

	requestAccessURL, err := service.cloudinary.Upload(ctx, "request-accesses", filename, supportEvidence)

	if err != nil {
		return err
	}

	if err := service.pemustakaRepository.Save(ctx, tx, pemustakaDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	requestAccessDomain := domain.RequestAccess{
		ID:              uuid.NewString(),
		PemustakaId:     pemustakaDomain.ID,
		SupportEvidence: requestAccessURL,
		Status:          "pending",
	}

	if err := service.requestAccessRepository.Save(ctx, tx, requestAccessDomain); err != nil {
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

func (service PemustakaServiceImpl) Login(ctx context.Context, req request.LoginPemustakaRequest) (string, error) {
	user, err := service.userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return "", err
	}

	ok := utils.ComparePassword(user.Password, req.Password)

	if !ok {
		return "", utils.ErrUserNotFound
	}

	pemustaka, err := service.pemustakaRepository.FindByUserId(ctx, user.ID)

	if err != nil {
		return "", utils.ErrPemustakaNotFound
	}

	if pemustaka.IsActive == "0" {
		return "", utils.ErrWaitingForAcceptance
	}

	token, _ := utils.GenerateToken(pemustaka.ID, pemustaka.Role.Role)

	return token, nil
}

func (service PemustakaServiceImpl) Update(ctx context.Context, req request.UpdatePemustakaRequest, avatar *multipart.FileHeader, pemustakaId string) error {
	tx := service.tx.Begin()

	pemustaka, err := service.pemustakaRepository.FindById(ctx, pemustakaId)

	if err != nil {
		return err
	}

	if _, err := service.studyProgramRepository.FindById(ctx, req.StudyProgramId); err != nil {
		return err
	}

	if _, err := service.departementRepository.FindById(ctx, req.DepartementId); err != nil {
		return err
	}

	var avatarURL string

	if avatar != nil {
		if pemustaka.Avatar != configs.GetConfig("DEFAULT_AVATAR") {
			if err := service.cloudinary.Delete(ctx, pemustaka.Avatar); err != nil {
				return err
			}
		}

		filename := utils.GetFilename()

		avatarURL, err = service.cloudinary.Upload(ctx, "avatars", filename, avatar)

		if err != nil {
			return err
		}
	}

	pemustakaDomain := req.ToDomainPemustaka()
	pemustakaDomain.Avatar = avatarURL

	if err = service.pemustakaRepository.Update(ctx, tx, pemustakaDomain, pemustakaId); err != nil {
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

func (service PemustakaServiceImpl) FindAll(ctx context.Context, query request.PemustakaRequestQuery, limit int, offset int) ([]response.PemustakaResponse, int, int, error) {
	pemustakas, totalRows, err := service.pemustakaRepository.FindAll(ctx, query, limit, offset)

	if err != nil {
		return []response.PemustakaResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToPemustakaArrayResponse(pemustakas), totalRows, int(totalPages), nil
}

func (service PemustakaServiceImpl) FindById(ctx context.Context, pemustakaId string) (response.PemustakaResponse, error) {
	pemustaka, err := service.pemustakaRepository.FindById(ctx, pemustakaId)

	if err != nil {
		return response.PemustakaResponse{}, err
	}

	return response.ToPemustakaResponse(pemustaka), nil
}
