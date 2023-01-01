package drivers

import (
	"github.com/arvinpaundra/repository-api/drivers/mysql/category"
	"github.com/arvinpaundra/repository-api/drivers/mysql/collection"
	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) category.CategoryRepository {
	return category.NewRepositorySQL(conn)
}

func NewCollectionRepository(conn *gorm.DB) collection.CollectionRepository {
	return collection.NewRepositorySQL(conn)
}
