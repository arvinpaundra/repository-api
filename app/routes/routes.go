package routes

import (
	"github.com/arvinpaundra/repository-api/drivers"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/arvinpaundra/repository-api/helper/mailing"
	"github.com/arvinpaundra/repository-api/middlewares"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	categoryController "github.com/arvinpaundra/repository-api/controllers/category"
	categoryService "github.com/arvinpaundra/repository-api/services/category"

	identityCardController "github.com/arvinpaundra/repository-api/controllers/identityCard"
	identityCardService "github.com/arvinpaundra/repository-api/services/identityCard"

	reportController "github.com/arvinpaundra/repository-api/controllers/report"
	reportService "github.com/arvinpaundra/repository-api/services/report"

	staffController "github.com/arvinpaundra/repository-api/controllers/staff"
	staffService "github.com/arvinpaundra/repository-api/services/staff"

	dashboardController "github.com/arvinpaundra/repository-api/controllers/dashboard"
	dashboardSrvc "github.com/arvinpaundra/repository-api/services/dashboard"

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
	Mailing    *mailing.Mailing
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

	staffRepository := drivers.NewStaffRepository(rc.MySQl)
	staffSrvc := staffService.NewStaffService(rc.MySQl, userRepository, staffRepository, roleRepository, rc.Cloudinary)
	staffCtrl := staffController.NewStaffController(staffSrvc)

	pemustakaSrvc := pemustakaService.NewPemustakaService(userRepository, pemustakaRepository, studyProgramRepository, departementRepository, roleRepository, requestAccessRepository, rc.Cloudinary, rc.MySQl)
	pemustakaCtrl := pemustakaController.NewPemustakaController(pemustakaSrvc)

	requestAccessSrvc := requestAccessService.NewRequestAccessService(requestAccessRepository, pemustakaRepository, *rc.Mailing, rc.MySQl)
	requestAccessCtrl := requestAccessController.NewRequestAccessController(requestAccessSrvc)

	authSrvc := authService.NewAuthService(userRepository, exporationTokenRepository, rc.MySQl)
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
		categoryRepository,
		pemustakaRepository,
		authorRepository,
		contributorRepository,
		repoRepository,
		documentRepository,
		rc.Cloudinary,
		*rc.Mailing,
		rc.MySQl,
	)
	repositoryCtrl := repositoryController.NewRepositoryController(repositorySrvc)

	authorService := authorService.NewAuthorService(authorRepository, repoRepository, pemustakaRepository)
	authorCtrl := authorController.NewAuthorController(authorService)

	dashboardService := dashboardSrvc.NewDashboardService(pemustakaRepository, repoRepository, requestAccessRepository)
	dashboardCtrl := dashboardController.NewDashboardController(dashboardService)

	identityCardSrvc := identityCardService.NewIdentityCardService(pemustakaRepository)
	identityCardCtrl := identityCardController.NewIdentityCardController(identityCardSrvc)

	reportRepository := drivers.NewReportRepository(rc.MySQl)
	reportSrvc := reportService.NewReportService(pemustakaRepository, collectionRepository, staffRepository, reportRepository)
	reportCtrl := reportController.NewReportController(reportSrvc)

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
	departement.POST("", departementCtrl.HandlerCreateDepartement, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	departement.GET("", departementCtrl.HandlerFindAllDepartements)

	departementDetail := departement.Group("/:departementId")
	departementDetail.PUT("", departementCtrl.HandlerUpdateDepartement, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	departementDetail.GET("", departementCtrl.HandlerFindDepartementById)
	departementDetail.GET("/study-programs", studyProgramCtrl.HandlerFindByDepartementId)
	departementDetail.GET("/repositories", repositoryCtrl.HandlerFindByDepartementId)

	// study program routes
	studyProgram := v1.Group("/study-programs")
	studyProgram.POST("", studyProgramCtrl.HandlerCreateStudyProgram, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	studyProgram.GET("", studyProgramCtrl.HandlerFindAllStudyPrograms)

	studyProgramDetail := studyProgram.Group("/:studyProgramId")
	studyProgramDetail.PUT("", studyProgramCtrl.HandlerUpdateStudyProgram, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	studyProgramDetail.GET("", studyProgramCtrl.HandlerFindStudyProgramById)

	// auth routes
	auth := v1.Group("/auth")

	authUser := auth.Group("/users")
	authUser.PUT("/:userId/change-password", authCtrl.HandlerChangePassword, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))

	forgotPassword := auth.Group("/forgot-password")
	forgotPassword.POST("", mailingCtrl.HandlerSendForgotPasswordWithExpirationToken)
	forgotPassword.PUT("", authCtrl.HandlerForgotPassword)

	authStaff := auth.Group("/staff")
	authStaff.POST("/register", staffCtrl.HandlerRegister, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator"}))
	authStaff.POST("/login", staffCtrl.HandlerLogin)

	authPemustaka := auth.Group("/pemustaka")
	authPemustaka.POST("/register", pemustakaCtrl.HandlerRegister, middlewares.UploadSupportEvidence())
	authPemustaka.POST("/login", pemustakaCtrl.HandlerLogin)

	// pemustaka routes
	pemustaka := v1.Group("/pemustaka")
	pemustaka.GET("", pemustakaCtrl.HandlerFindAllPemustaka)
	pemustaka.POST("", pemustakaCtrl.HandleCreatePemustaka, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))

	pemustakaDetail := pemustaka.Group("/:pemustakaId", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))
	pemustakaDetail.PUT("", pemustakaCtrl.HandlerUpdatePemustaka, middlewares.UploadAvatarValidator())
	pemustakaDetail.GET("", pemustakaCtrl.HandlerFindPemustakaById)

	// request access routes
	requestAccess := v1.Group("/request-accesses", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan"}))
	requestAccess.GET("", requestAccessCtrl.HandlerFindAllRequestAccesses)
	requestAccess.GET("/total", requestAccessCtrl.HandlerGetTotal)

	requestAccessDetail := requestAccess.Group("/:requestAccessId")
	requestAccessDetail.GET("", requestAccessCtrl.HandlerFindRequestAccessById)
	requestAccessDetail.PUT("", requestAccessCtrl.HandlerUpdateRequestAccess)

	// repository routes
	repository := v1.Group("/repositories")
	repository.GET("", repositoryCtrl.HandlerFindAllRepositories)
	repository.GET("/total", repositoryCtrl.HandlerGetTotalRepository, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Kepala Perpustakaan", "Pustakawan"}))

	finalProject := repository.Group("/final-projects", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Mahasiswa"}), middlewares.UploadRepositoryFiles())
	finalProject.POST("", repositoryCtrl.HandlerCreateFinalProjectReport)
	finalProject.PUT("/:repositoryId", repositoryCtrl.HandlerUpdateFinalProjectReport)

	internshipReport := repository.Group("/internship-reports", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Mahasiswa"}), middlewares.UploadRepositoryFiles())
	internshipReport.POST("", repositoryCtrl.HandlerCreateInternshipReport)
	internshipReport.PUT("/:repositoryId", repositoryCtrl.HandlerUpdateInternshipReport)

	researchReport := repository.Group("/research-reports", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Dosen", "Mahasiswa"}), middlewares.UploadRepositoryFiles())
	researchReport.POST("", repositoryCtrl.HandlerCreateResearchReport)
	researchReport.PUT("/:repositoryId", repositoryCtrl.HandlerUpdateResearchReport)

	repositoryDetail := repository.Group("/:repositoryId")
	repositoryDetail.GET("", repositoryCtrl.HandlerFindRepositoryById)
	repositoryDetail.DELETE("", repositoryCtrl.HandlerDeleteRepository, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))
	repositoryDetail.PUT("/confirm", repositoryCtrl.HandlerConfirmRepository, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan"}))

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

	// staff routes
	staff := v1.Group("/staffs", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan"}))
	staff.GET("", staffCtrl.HandlerFindAllStaffs)

	staffDetail := staff.Group("/:staffId")
	staffDetail.GET("", staffCtrl.HandlerFindStaffById)
	staffDetail.PUT("", staffCtrl.HandlerUpdateStaff, middlewares.UploadAvatarValidator())
	staffDetail.PUT("/signatures", staffCtrl.HandlerUploadSignature)

	// dashboard routes
	dashboard := v1.Group("/dashboard", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Kepala Perpustakaan", "Pustakawan"}))
	dashboard.GET("/overview", dashboardCtrl.HandlerOverview)

	// identity card routes
	identityCard := v1.Group("/identity-card")
	identityCard.GET("/:pemustakaId", identityCardCtrl.HandlerGenerateIDCard, middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan", "Mahasiswa", "Dosen"}))

	// report routes
	report := v1.Group("/reports", middlewares.IsAuthenticated(), middlewares.CheckRoles([]string{"Administrator", "Pustakawan", "Kepala Perpustakaan"}))
	report.POST("/surat-keterangan-penyerahan-laporan", reportCtrl.HandlerGetSuratKeteranganPenyerahanLaporan)

	recapCollectedReport := report.Group("/recap-collected-report")
	recapCollectedReport.GET("", reportCtrl.HandlerRecapCollectedReport)
	recapCollectedReport.GET("/download", reportCtrl.HandlerDownloadRecapCollectedReport)
}
