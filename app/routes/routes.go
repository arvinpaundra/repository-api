package routes

import (
	"github.com/arvinpaundra/repository-api/drivers"
	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/arvinpaundra/repository-api/middlewares"
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

	repositoryController "github.com/arvinpaundra/repository-api/controllers/repository"
	repositoryService "github.com/arvinpaundra/repository-api/services/repository"

	authorController "github.com/arvinpaundra/repository-api/controllers/author"
	authorService "github.com/arvinpaundra/repository-api/services/author"
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

	departementRepository := drivers.NewDepartementRepository(rc.MySQl)
	departementSrvc := departementService.NewDepartementService(departementRepository)
	departementCtrl := departementController.NewDepartementController(departementSrvc)

	studyProgramRepository := drivers.NewStudyProgramRepository(rc.MySQl)
	studyProgramSrvc := studyProgramService.NewStudyProgramService(studyProgramRepository, departementRepository)
	studyProgramCtrl := studyProgramController.NewStudyProgramController(studyProgramSrvc)

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

	authorRepository := drivers.NewAuthorRepository(rc.MySQl)

	contributorRepository := drivers.NewContributorRepository(rc.MySQl)

	documentRepository := drivers.NewDocumentRepository(rc.MySQl)

	repoRepository := drivers.NewRepository(rc.MySQl)
	repositorySrvc := repositoryService.NewRepositoryService(
		collectionRepository,
		departementRepository,
		pemustakaRepository,
		authorRepository,
		contributorRepository,
		repoRepository,
		documentRepository,
		rc.Cloudinary,
		rc.MySQl,
	)
	repositoryCtrl := repositoryController.NewRepositoryController(repositorySrvc)

	authorService := authorService.NewAuthorService(authorRepository, repoRepository, pemustakaRepository)
	authorCtrl := authorController.NewAuthorController(authorService)

	// API version
	v1 := rc.Echo.Group("/api/v1")

	// category routes
	category := v1.Group("/categories")
	category.POST("", categoryCtrl.HandlerCreateCategory)
	category.GET("", categoryCtrl.HandlerFindAllCategories)
	categoryDetail := category.Group("/:categoryId")
	categoryDetail.PUT("", categoryCtrl.HandlerUpdateCategory)
	categoryDetail.GET("", categoryCtrl.HandlerFindCategoryById)

	// collection routes
	collection := v1.Group("/collections")
	collection.POST("", collectionCtrl.HandlerCreateCollection, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	collection.GET("", collectionCtrl.HandlerFindAllCollections)
	collectionDetail := collection.Group("/:collectionId")
	collectionDetail.PUT("", collectionCtrl.HandlerUpdateCollection, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	collectionDetail.GET("", collectionCtrl.HandlerFindCollectionById)
	collectionDetail.GET("/repositories", repositoryCtrl.HandlerFindByCollectionId)

	// collection routes
	role := v1.Group("/roles")
	role.POST("", roleCtrl.HandlerCreateRole, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	role.PUT("/:roleId", roleCtrl.HandlerUpdateRole, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	role.GET("", roleCtrl.HandlerFindAllRoles)
	role.GET("/:roleId", roleCtrl.HandlerFindRoleById)

	// departement routes
	departement := v1.Group("/departements")
	departement.POST("", departementCtrl.HandlerCreateDepartement, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	departement.GET("", departementCtrl.HandlerFindAllDepartements)

	departementDetail := departement.Group("/:departementId")
	departementDetail.PUT("", departementCtrl.HandlerUpdateDepartement, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	departementDetail.GET("", departementCtrl.HandlerFindDepartementById)
	departementDetail.GET("/study-programs", studyProgramCtrl.HandlerFindByDepartementId)
	departementDetail.GET("/repositories", repositoryCtrl.HandlerFindByDepartementId)

	// study program routes
	studyProgram := v1.Group("/study-programs")
	studyProgram.POST("", studyProgramCtrl.HandlerCreateStudyProgram, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	studyProgram.GET("", studyProgramCtrl.HandlerFindAllStudyPrograms)

	studyProgramDetail := studyProgram.Group("/:studyProgramId")
	studyProgramDetail.PUT("", studyProgramCtrl.HandlerUpdateStudyProgram, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	studyProgramDetail.GET("", studyProgramCtrl.HandlerFindStudyProgramById)

	// auth routes
	auth := v1.Group("/auth")

	authUser := auth.Group("/users")
	authUser.PUT("/:userId/change-password", authCtrl.HandlerChangePassword, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))

	forgotPassword := auth.Group("/forgot-password")
	forgotPassword.POST("", mailingCtrl.HandlerSendForgotPasswordWithExpirationToken)
	forgotPassword.PUT("", authCtrl.HandlerForgotPassword)

	authPemustaka := auth.Group("/pemustaka")
	authPemustaka.POST("/register", pemustakaCtrl.HandlerRegister)
	authPemustaka.POST("/login", pemustakaCtrl.HandlerLogin)

	// pemustaka routes
	pemustaka := v1.Group("/pemustaka")
	pemustaka.GET("", pemustakaCtrl.HandlerFindAllPemustaka)

	pemustakaDetail := pemustaka.Group("/:pemustakaId", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Mahasiswa", "Dosen"}))
	pemustakaDetail.PUT("", pemustakaCtrl.HandlerUpdatePemustaka)
	pemustakaDetail.GET("", pemustakaCtrl.HandlerFindPemustakaById)

	// request access routes
	requestAccess := v1.Group("/request-accesses", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	requestAccess.GET("", requestAccessCtrl.HandlerFindAllRequestAccesses)

	requestAccessDetail := requestAccess.Group("/:requestAccesId")
	requestAccessDetail.GET("", requestAccessCtrl.HandlerFindRequestAccessById)
	requestAccessDetail.PUT("", requestAccessCtrl.HandlerUpdateRequestAccess)

	// repository routes
	repository := v1.Group("/repositories")
	repository.GET("", repositoryCtrl.HandlerFindAllRepositories)
	repository.POST("/final-projects", repositoryCtrl.HandlerCreateFinalProjectReport, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Mahasiswa"}))
	repository.POST("/internship-report", repositoryCtrl.HandlerCreateInternshipReport, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Mahasiswa"}))
	repository.POST("/research-report", repositoryCtrl.HandlerCreateResearchReport, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))

	repositoryDetail := repository.Group("/:repositoryId")
	repositoryDetail.GET("", repositoryCtrl.HandlerFindRepositoryById)
	repositoryDetail.DELETE("", repositoryCtrl.HandlerDeleteRepository, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))

	// author routes
	author := v1.Group("/authors")
	author.GET("/:pemustakaId/repositories", repositoryCtrl.HandlerFindByAuthorId, middlewares.IsAuthenticated())
	author.DELETE("/:pemustakaId/repositories/:repositoryId", authorCtrl.HandlerDeleteAuthor, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))

	// mentor routes
	mentor := v1.Group("/mentors", middlewares.IsAuthenticated())
	mentor.GET("/:pemustakaId/repositories", repositoryCtrl.HandlerFindByMentorId)

	// examiner routes
	examiner := v1.Group("/examiners", middlewares.IsAuthenticated())
	examiner.GET("/:pemustakaId/repositories", repositoryCtrl.HandlerFindByExaminerId)
}
