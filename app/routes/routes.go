package routes

import (
	controller "github.com/arvinpaundra/repository-api/controllers"
	"github.com/arvinpaundra/repository-api/drivers"
	service "github.com/arvinpaundra/repository-api/services"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RouteConfig struct {
	Echo  *echo.Echo
	MySQl *gorm.DB
	Redis *redis.Client
}

func (rc *RouteConfig) New() {
	categoryRepository := drivers.NewCategoryRepository(rc.MySQl)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	collectionRepository := drivers.NewCollectionRepository(rc.MySQl)
	collectionService := service.NewCollectionService(collectionRepository)
	collectionController := controller.NewCollectionController(collectionService)

	roleRepository := drivers.NewRoleRepository(rc.MySQl)
	roleService := service.NewRoleService(roleRepository)
	roleController := controller.NewRoleController(roleService)

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
