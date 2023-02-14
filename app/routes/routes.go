package routes

import (
	"github.com/arvinpaundra/repository-api/drivers"
	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	categoryController "github.com/arvinpaundra/repository-api/controllers/category"
	categoryService "github.com/arvinpaundra/repository-api/services/category"

	collectionController "github.com/arvinpaundra/repository-api/controllers/collection"
	collectionService "github.com/arvinpaundra/repository-api/services/collection"

	roleController "github.com/arvinpaundra/repository-api/controllers/role"
	roleService "github.com/arvinpaundra/repository-api/services/role"

	studyProgramController "github.com/arvinpaundra/repository-api/controllers/studyProgram"
	studyProgramService "github.com/arvinpaundra/repository-api/services/studyProgram"

	departementController "github.com/arvinpaundra/repository-api/controllers/departement"
	departementService "github.com/arvinpaundra/repository-api/services/departement"

	requestAccessController "github.com/arvinpaundra/repository-api/controllers/requestAccess"
	requestAccessService "github.com/arvinpaundra/repository-api/services/requestAccess"

	pemustakaController "github.com/arvinpaundra/repository-api/controllers/pemustaka"
	pemustakaService "github.com/arvinpaundra/repository-api/services/pemustaka"

	authController "github.com/arvinpaundra/repository-api/controllers/auth"
	authService "github.com/arvinpaundra/repository-api/services/auth"

	mailingController "github.com/arvinpaundra/repository-api/controllers/mailing"
	mailingService "github.com/arvinpaundra/repository-api/services/mailing"
)

type RouteConfig struct {
	Echo       *echo.Echo
	MySQl      *gorm.DB
	Redis      *redis.Client
	Mailing    *helper.Mailing
	Cloudinary cloudinary.Cloudinary
}

func (rc *RouteConfig) New() {
	categoryRepository := drivers.NewCategoryRepository(rc.MySQl)
	categorySrvc := categoryService.NewCategoryService(categoryRepository)
	categoryCtrl := categoryController.NewCategoryController(categorySrvc)

	collectionRepository := drivers.NewCollectionRepository(rc.MySQl)
	collectionSrvc := collectionService.NewCollectionService(collectionRepository)
	collectionCtrl := collectionController.NewCollectionController(collectionSrvc)

	roleRepository := drivers.NewRoleRepository(rc.MySQl)
	roleSrvc := roleService.NewRoleService(roleRepository)
	roleCtrl := roleController.NewRoleController(roleSrvc)

	studyProgramRepository := drivers.NewStudyProgramRepository(rc.MySQl)
	studyProgramSrvc := studyProgramService.NewStudyProgramService(studyProgramRepository)
	studyProgramCtrl := studyProgramController.NewStudyProgramController(studyProgramSrvc)

	departementRepository := drivers.NewDepartementRepository(rc.MySQl)
	departementSrvc := departementService.NewDepartementService(departementRepository, studyProgramRepository)
	departementCtrl := departementController.NewDepartementController(departementSrvc)

	exporationTokenRepository := drivers.NewExpirationRepository(rc.Redis)
	userRepository := drivers.NewUserRepository(rc.MySQl)
	pemustakaRepository := drivers.NewPemustakaRepository(rc.MySQl)
	requestAccessRepository := drivers.NewRequestAccessRepository(rc.MySQl)

	pemustakaSrvc := pemustakaService.NewPemustakaService(userRepository, pemustakaRepository, studyProgramRepository, departementRepository, roleRepository, requestAccessRepository, rc.Cloudinary, rc.MySQl)
	pemustakaCtrl := pemustakaController.NewPemustakaController(pemustakaSrvc)

	requestAccessSrvc := requestAccessService.NewRequestAccessService(requestAccessRepository, pemustakaRepository, rc.MySQl)
	requestAccessCtrl := requestAccessController.NewRequestAccessController(requestAccessSrvc)

	authSrvc := authService.NewAuthService(userRepository, exporationTokenRepository)
	authCtrl := authController.NewAuthController(authSrvc)

	mailingSrvc := mailingService.NewMailingService(exporationTokenRepository, userRepository, pemustakaRepository, *rc.Mailing)
	mailingCtrl := mailingController.NewMailingController(mailingSrvc)

	// API current version
	v1 := rc.Echo.Group("/api/v1")

	// category routes
	category := v1.Group("/categories")
	category.POST("", categoryCtrl.HandlerCreateCategory)
	category.PUT("/:categoryId", categoryCtrl.HandlerUpdateCategory)
	category.GET("", categoryCtrl.HandlerFindAllCategories)
	category.GET("/:categoryId", categoryCtrl.HandlerFindCategoryById)

	// collection routes
	collection := v1.Group("/collections")
	collection.POST("", collectionCtrl.HandlerCreateCollection)
	collection.PUT("/:collectionId", collectionCtrl.HandlerUpdateCollection)
	collection.GET("", collectionCtrl.HandlerFindAllCollections)
	collection.GET("/:collectionId", collectionCtrl.HandlerFindCollectionById)

	// collection routes
	role := v1.Group("/roles")
	role.POST("", roleCtrl.HandlerCreateRole)
	role.PUT("/:roleId", roleCtrl.HandlerUpdateRole)
	role.GET("", roleCtrl.HandlerFindAllRoles)
	role.GET("/:roleId", roleCtrl.HandlerFindRoleById)

	// study program routes
	studyProgram := v1.Group("/study-programs")
	studyProgram.POST("", studyProgramCtrl.HandlerCreateStudyProgram)
	studyProgram.GET("", studyProgramCtrl.HandlerFindAllStudyPrograms)

	detailStudyProgram := studyProgram.Group("/:studyProgramId")
	detailStudyProgram.PUT("", studyProgramCtrl.HandlerUpdateStudyProgram)
	detailStudyProgram.GET("", studyProgramCtrl.HandlerFindStudyProgramById)
	detailStudyProgram.GET("/departements", departementCtrl.HandlerFindDepartementsByStudyProgramId)

	// departement routes
	departement := v1.Group("/departements")
	departement.POST("", departementCtrl.HandlerCreateDepartement)
	departement.PUT("/:departementId", departementCtrl.HandlerUpdateDepartement)
	departement.GET("", departementCtrl.HandlerFindAllDepartements)
	departement.GET("/:departementId", departementCtrl.HandlerFindDepartementById)

	// auth routes
	auth := v1.Group("/auth")
	forgotPassword := auth.Group("/forgot-password")
	forgotPassword.POST("", mailingCtrl.HandlerSendForgotPasswordWithExpirationToken)
	forgotPassword.PUT("", authCtrl.HandlerForgotPassword)

	authPemustaka := auth.Group("/pemustaka")
	authPemustaka.POST("/register", pemustakaCtrl.HandlerRegister)
	authPemustaka.POST("/login", pemustakaCtrl.HandlerLogin)

	// pemustaka routes
	pemustaka := v1.Group("/pemustaka")
	pemustaka.PUT("/:pemustakaId", pemustakaCtrl.HandlerUpdatePemustaka)
	pemustaka.GET("", pemustakaCtrl.HandlerFindAllPemustaka)
	pemustaka.GET("/:pemustakaId", pemustakaCtrl.HandlerFindPemustakaById)

	// request access routes
	requestAccess := v1.Group("/request-accesses")
	requestAccess.GET("", requestAccessCtrl.HandlerFindAllRequestAccesses)
	requestAccess.GET("/:requestAccessId", requestAccessCtrl.HandlerFindRequestAccessById)
	requestAccess.PUT("/:requestAccessId", requestAccessCtrl.HandlerUpdateRequestAccess)
}
