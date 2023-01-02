package drivers

import (
	"github.com/arvinpaundra/repository-api/drivers/mysql/category"
	"github.com/arvinpaundra/repository-api/drivers/mysql/collection"
	"github.com/arvinpaundra/repository-api/drivers/mysql/role"
	studyProgram "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram"
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
