package drivers

import (
	"github.com/arvinpaundra/repository-api/drivers/mysql/author"
	"github.com/arvinpaundra/repository-api/drivers/mysql/category"
	"github.com/arvinpaundra/repository-api/drivers/mysql/collection"
	"github.com/arvinpaundra/repository-api/drivers/mysql/contributor"
	"github.com/arvinpaundra/repository-api/drivers/mysql/departement"
	"github.com/arvinpaundra/repository-api/drivers/mysql/document"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/repository"
	requestAccess "github.com/arvinpaundra/repository-api/drivers/mysql/requestAccess"
	"github.com/arvinpaundra/repository-api/drivers/mysql/role"
	studyProgram "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram"
	"github.com/arvinpaundra/repository-api/drivers/mysql/user"
	expirationToken "github.com/arvinpaundra/repository-api/drivers/redis/expirationToken"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) category.CategoryRepository {
	return category.NewRepositorySQL(conn)
}

func NewCollectionRepository(conn *gorm.DB) collection.CollectionRepository {
	return collection.NewRepositorySQL(conn)
}

func NewRoleRepository(conn *gorm.DB) role.RoleRepository {
	return role.NewSQLRepository(conn)
}

func NewStudyProgramRepository(conn *gorm.DB) studyProgram.StudyProgramRepository {
	return studyProgram.NewSQLRepository(conn)
}

func NewDepartementRepository(conn *gorm.DB) departement.DepartementRepository {
	return departement.NewSQLRepository(conn)
}

func NewExpirationRepository(conn *redis.Client) expirationToken.ExpirationTokenRepository {
	return expirationToken.NewRedisRepository(conn)
}

func NewPemustakaRepository(conn *gorm.DB) pemustaka.PemustakaRepository {
	return pemustaka.NewSQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) user.UserRepository {
	return user.NewSQLRepository(conn)
}

func NewRequestAccessRepository(conn *gorm.DB) requestAccess.RequestAccessRepository {
	return requestAccess.NewSQLRepository(conn)
}

func NewRepository(conn *gorm.DB) repository.Repository {
	return repository.NewSQLRepository(conn)
}

func NewAuthorRepository(conn *gorm.DB) author.AuthorRepository {
	return author.NewSQLRepository(conn)
}

func NewContributorRepository(conn *gorm.DB) contributor.ContributorRepository {
	return contributor.NewSQLRepository(conn)
}

func NewDocumentRepository(conn *gorm.DB) document.DocumentRepository {
	return document.NewSQLRepository(conn)
}
