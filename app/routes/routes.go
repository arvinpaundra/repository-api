package routes

import (
	"github.com/arvinpaundra/repository-api/drivers"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	categoryController "github.com/arvinpaundra/repository-api/controllers/category"
	categoryService "github.com/arvinpaundra/repository-api/services/category"

	collectionController "github.com/arvinpaundra/repository-api/controllers/collection"
	collectionService "github.com/arvinpaundra/repository-api/services/collection"

	roleController "github.com/arvinpaundra/repository-api/controllers/role"
	roleService "github.com/arvinpaundra/repository-api/services/role"
)

type RouteConfig struct {
	Echo  *echo.Echo
	MySQl *gorm.DB
	Redis *redis.Client
}

func (rc *RouteConfig) New() {
	categoryRepository := drivers.NewCategoryRepository(rc.MySQl)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	categoryController := categoryController.NewCategoryController(categoryService)

	collectionRepository := drivers.NewCollectionRepository(rc.MySQl)
	collectionService := collectionService.NewCollectionService(collectionRepository)
	collectionController := collectionController.NewCollectionController(collectionService)

	roleRepository := drivers.NewRoleRepository(rc.MySQl)
	roleService := roleService.NewRoleService(roleRepository)
	roleController := roleController.NewRoleController(roleService)

	// API current version
	v1 := rc.Echo.Group("/api/v1")

	// category routes
	category := v1.Group("/categories")
	category.POST("", categoryController.HandlerCreateCategory)
	category.PUT("/:categoryId", categoryController.HandlerUpdateCategory)
	category.GET("", categoryController.HandlerFindAllCategories)
	category.GET("/:categoryId", categoryController.HandlerFindCategoryById)

	collection := v1.Group("/collections")
	collection.POST("", collectionController.HandlerCreateCollection)
	collection.PUT("/:collectionId", collectionController.HandlerUpdateCollection)
	collection.GET("", collectionController.HandlerFindAllCollections)
	collection.GET("/:collectionId", collectionController.HandlerFindCollectionById)

	role := v1.Group("/roles")
	role.POST("", roleController.HandlerCreateRole)
	role.PUT("/:roleId", roleController.HandlerUpdateRole)
	role.GET("", roleController.HandlerFindAllRoles)
	role.GET("/:roleId", roleController.HandlerFindRoleById)
}
