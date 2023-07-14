package staff

import (
	"context"
	"math"
	"mime/multipart"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/drivers/mysql/role"
	"github.com/arvinpaundra/repository-api/drivers/mysql/staff"
	"github.com/arvinpaundra/repository-api/drivers/mysql/user"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/staff/request"
	"github.com/arvinpaundra/repository-api/models/web/staff/response"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StaffServiceImpl struct {
	tx              *gorm.DB
	userRepository  user.UserRepository
	staffRepository staff.StaffRepository
	roleRepository  role.RoleRepository
	cloudinary      cloudinary.Cloudinary
}

func NewStaffService(
	tx *gorm.DB,
	userRepository user.UserRepository,
	staffRepository staff.StaffRepository,
	roleRepository role.RoleRepository,
	cloudinary cloudinary.Cloudinary,
) StaffService {
	return StaffServiceImpl{
		tx:              tx,
		userRepository:  userRepository,
		staffRepository: staffRepository,
		roleRepository:  roleRepository,
		cloudinary:      cloudinary,
	}
}

func (service StaffServiceImpl) Register(ctx context.Context, req request.RegisterStaffRequest) error {
	tx := service.tx.Begin()

	user, _ := service.userRepository.FindByEmail(ctx, req.Email)

	if user.Email != "" {
		return utils.ErrEmailAlreadyUsed
	}

	if _, err := service.roleRepository.FindById(ctx, req.RoleId); err != nil {
		return err
	}

	password := utils.HashPassword(configs.GetConfig("DEFAULT_USER_PASSWORD"))

	userDomain := domain.User{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: password,
	}

	if err := service.userRepository.Save(ctx, tx, userDomain); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	staffDomain := req.ToDomainStaff()
	staffDomain.ID = uuid.NewString()
	staffDomain.UserId = userDomain.ID
	staffDomain.Avatar = configs.GetConfig("DEFAULT_AVATAR")

	if err := service.staffRepository.Save(ctx, tx, staffDomain); err != nil {
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

func (service StaffServiceImpl) Login(ctx context.Context, req request.LoginStaffRequest) (string, error) {
	user, err := service.userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return "", err
	}

	if ok := utils.ComparePassword(user.Password, req.Password); !ok {
		return "", utils.ErrUserNotFound
	}

	staff, err := service.staffRepository.FindByUserId(ctx, user.ID)

	if err != nil {
		return "", err
	}

	if staff.IsActive == "0" {
		return "", utils.ErrAccountNotActivated
	}

	token, _ := utils.GenerateToken(staff.ID, staff.Role.Role)

	return token, nil
}

func (service StaffServiceImpl) Update(ctx context.Context, req request.UpdateStaffRequest, avatar *multipart.FileHeader, staffId string) error {
	tx := service.tx.Begin()

	staff, err := service.staffRepository.FindById(ctx, staffId)

	if err != nil {
		return err
	}

	if _, err := service.roleRepository.FindById(ctx, req.RoleId); err != nil {
		return err
	}

	if user, _ := service.userRepository.FindByEmail(ctx, req.Email); user.Email != "" && user.ID != staff.UserId {
		return utils.ErrEmailAlreadyUsed
	}

	var userDomain domain.User
	userDomain.Email = req.Email

	if err := service.userRepository.Update(ctx, tx, userDomain, staff.User.Email); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	var avatarURL string

	if avatar != nil {
		if staff.Avatar != configs.GetConfig("DEFAULT_AVATAR") {
			if err := service.cloudinary.Delete(ctx, staff.Avatar); err != nil {
				return err
			}
		}

		avatarURL, err = service.cloudinary.Upload(ctx, "avatars", utils.GetFilename(), avatar)

		if err != nil {
			return err
		}
	}

	staffDomain := req.ToDomainStaff()
	staffDomain.Avatar = avatarURL

	if err := service.staffRepository.Update(ctx, tx, staffId, staffDomain); err != nil {
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

func (service StaffServiceImpl) FindAll(ctx context.Context, query request.StaffRequestQuery, limit int, offset int) ([]response.StaffResponse, int, int, error) {
	staffs, totalRows, err := service.staffRepository.FindAll(ctx, query, limit, offset)

	if err != nil {
		return []response.StaffResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToStaffArrayResponse(staffs), totalRows, int(totalPages), nil
}

func (service StaffServiceImpl) FindById(ctx context.Context, staffId string) (response.StaffResponse, error) {
	staff, err := service.staffRepository.FindById(ctx, staffId)

	if err != nil {
		return response.StaffResponse{}, err
	}

	return response.ToStaffResponse(staff), nil
}

func (service StaffServiceImpl) UploadSignature(ctx context.Context, signature *multipart.FileHeader, staffId string) error {
	tx := service.tx.Begin()

	staff, err := service.staffRepository.FindById(ctx, staffId)

	if err != nil {
		return err
	}

	if staff.Signature != "" {
		err := service.cloudinary.Delete(ctx, staff.Signature)

		if err != nil {
			return err
		}
	}

	signatureUrl, err := service.cloudinary.Upload(ctx, "signatures", utils.GetFilename(), signature)

	if err != nil {
		return err
	}

	if err := service.staffRepository.Update(ctx, tx, staffId, domain.Staff{Signature: signatureUrl}); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return err
		}

		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}
