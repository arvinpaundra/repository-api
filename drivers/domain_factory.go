package drivers

import (
	"github.com/arvinpaundra/repository-api/drivers/mysql/category"
	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) category.CategoryRepository {
	return category.NewRepositorySQL(conn)
}
