package pemustaka

import (
	"context"
	"math"

	"github.com/arvinpaundra/repository-api/drivers/mysql/departement"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/role"
	studyProgram "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram"
	"github.com/arvinpaundra/repository-api/drivers/mysql/user"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/response"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PemustakaServiceImpl struct {
	userRepository         user.UserRepository
	pemustakaRepository    pemustaka.PemustakaRepository
	studyProgramRepository studyProgram.StudyProgramRepository
	departementRepository  departement.DepartementRepository
	roleRepository         role.RoleRepository
	tx                     *gorm.DB
}

func NewPemustakaService(
	userRepository user.UserRepository,
	pemustakaRepository pemustaka.PemustakaRepository,
	studyProgramRepository studyProgram.StudyProgramRepository,
	departementRepository departement.DepartementRepository,
	roleRepository role.RoleRepository,
	tx *gorm.DB,
) PemustakaService {
	return PemustakaServiceImpl{
		userRepository:         userRepository,
		pemustakaRepository:    pemustakaRepository,
		studyProgramRepository: studyProgramRepository,
		departementRepository:  departementRepository,
		roleRepository:         roleRepository,
		tx:                     tx,
	}
}

func (service PemustakaServiceImpl) Register(ctx context.Context, req request.RegisterPemustakaRequest) error {
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

	if err := service.pemustakaRepository.Save(ctx, tx, pemustakaDomain); err != nil {
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

	token, _ := utils.GenerateToken(pemustaka.ID, pemustaka.Role.Role)

	return token, nil
}

func (service PemustakaServiceImpl) Update(ctx context.Context, req request.UpdatePemustakaRequest, pemustakaId string) error {
	if _, err := service.pemustakaRepository.FindById(ctx, pemustakaId); err != nil {
		return err
	}

	if _, err := service.studyProgramRepository.FindById(ctx, req.StudyProgramId); err != nil {
		return err
	}

	if _, err := service.departementRepository.FindById(ctx, req.DepartementId); err != nil {
		return err
	}

	var avatarUrl string

	if req.Avatar != nil {
		// TODO: do upload image
		avatarUrl = ""
	}

	pemustaka := req.ToDomainPemustaka()
	pemustaka.Avatar = avatarUrl

	err := service.pemustakaRepository.Update(ctx, pemustaka, pemustakaId)

	if err != nil {
		return err
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
