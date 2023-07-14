package utils

import "errors"

var (
	// error authentication failed
	ErrAuthenticationFailed = errors.New("Email atau kata sandi salah")

	// error user not found
	ErrUserNotFound = errors.New("Pengguna tidak ditemukan")

	// ErrCategoryNotFound error category not found
	ErrCategoryNotFound = errors.New("Kategori tidak ditemukan")

	// error collection not found
	ErrCollectionNotFound = errors.New("Koleksi tidak ditemukan")

	// error role not found
	ErrRoleNotFound = errors.New("Role tidak ditemukan")

	// error study program not found
	ErrStudyProgramNotFound = errors.New("Program studi tidak ditemukan")

	// error pemustaka not found
	ErrPemustakaNotFound = errors.New("Pemustaka tidak ditemukan")

	// error examiner not found
	ErrExaminerNotFound = errors.New("Penguji tidak ditemukan")

	// error mentor not found
	ErrMentorNotFound = errors.New("Pembimbing tidak ditemukan")

	// error departement not found
	ErrDepartementNotFound = errors.New("Jurusan tidak ditemukan")

	// error staff not found
	ErrStaffNotFound = errors.New("Petugas tidak ditemukan")

	// error head of library not found
	ErrHeadOfLibraryNotFound = errors.New("Kepala perpustakaan tidak ditemukan")

	// error repository not found
	ErrRepositoryNotFound = errors.New("Karya tulis ilmiah tidak ditemukan")

	// error document of repository not found
	ErrDocumentsNotFound = errors.New("Dokumen tidak ditemukan")

	// error request access not found
	ErrRequestAccessNotFound = errors.New("Permintaan akses tidak ditemukan")

	// error email already used
	ErrEmailAlreadyUsed = errors.New("Email sudah digunakan")

	// error email not found
	ErrEmailNotFound = errors.New("Email tidak ditemukan")

	// error invalid password length
	ErrInvalidPasswordLength = errors.New("Minimum panjang password yaitu 6 karakter")

	// error token has been expired
	ErrTokenExpired = errors.New("Token telah kadaluarsa")

	// error token has been sent
	ErrTokenHasBeenSent = errors.New("Token telah dikirim")

	// error token not match
	ErrTokenNotMatch = errors.New("Token salah")

	// error link has been expired
	ErrLinkExpired = errors.New("Link telah kadaluarsa")

	// error pemustaka already collected final project
	ErrAlreadyCollectedFinalProject = errors.New("Pemustaka sudah mengumpulkan Laporan Tugas Akhir")

	// error pemustaka already collected internship report
	ErrAlreadyCollectedInternshipReport = errors.New("Pemustaka sudah mengumpulkan Laporan Hasil Magang")

	// error fields are not equal
	ErrFieldsAreNotEqual = errors.New("Input tidak sama")

	// error internal server error
	ErrInternalServerError = errors.New("Internal server error")

	// error invalid token header
	ErrInvalidTokenHeader = errors.New("Invalid token header")

	// error forbidden access
	ErrForbiddenAccess = errors.New("Access forbidden")

	// error unauthorized user
	ErrUnAuthorized = errors.New("User unauthorized")

	// error account status not activated
	ErrAccountNotActivated = errors.New("Akun belum diaktivasi")

	// error pemustaka not collected final project report
	ErrNotCollectedFinalProject = errors.New("Pemustka belum mengunggah Laporan Tugas Akhir")

	// error pemustaka not collected internship report
	ErrNotCollectedInternshipReport = errors.New("Pemustaka belum mengunggah Laporan Hasil Magang")

	// error file type must be .pdf
	ErrFileTypeMustBePDF = errors.New("Hanya dapat mengunggah berkas berformat .pdf")
)
