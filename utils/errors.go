package utils

import "errors"

var (
	// error authentication failed
	ErrAuthenticationFailed = errors.New("Wrong email or password")

	// error user not found
	ErrUserNotFound = errors.New("User not found")

	// ErrCategoryNotFound error category not found
	ErrCategoryNotFound = errors.New("Category not found")

	// error collection not found
	ErrCollectionNotFound = errors.New("Collection not found")

	// error role not found
	ErrRoleNotFound = errors.New("Role not found")

	// error study program not found
	ErrStudyProgramNotFound = errors.New("Study program not found")

	// error pemustaka not found
	ErrPemustakaNotFound = errors.New("Pemustaka not found")

	// error examiner not found
	ErrExaminerNotFound = errors.New("Examiner not found")

	// error mentor not found
	ErrMentorNotFound = errors.New("Mentor not found")

	// error departement not found
	ErrDepartementNotFound = errors.New("Departement not found")

	// error staff not found
	ErrStaffNotFound = errors.New("Staff not found")

	// error head of library not found
	ErrHeadOfLibraryNotFound = errors.New("Head of library not found")

	// error repository not found
	ErrRepositoryNotFound = errors.New("Repository not found")

	// error document of repository not found
	ErrDocumentsNotFound = errors.New("Documents not found")

	// error request access not found
	ErrRequestAccessNotFound = errors.New("Request access not found")

	// error email already used
	ErrEmailAlreadyUsed = errors.New("Email already used")

	// error email not found
	ErrEmailNotFound = errors.New("Email not found")

	// error invalid password length
	ErrInvalidPasswordLength = errors.New("Minimum password length are 6 characters long")

	// error token has been expired
	ErrTokenExpired = errors.New("Token has been expired")

	// error token has been sent
	ErrTokenHasBeenSent = errors.New("Token has been sent")

	// error token not match
	ErrTokenNotMatch = errors.New("Token not matched")

	// error link has been expired
	ErrLinkExpired = errors.New("Link has been expired")

	// error pemustaka already collected final project
	ErrAlreadyCollectedFinalProject = errors.New("Pemustaka already collected final project")

	// error pemustaka already collected internship report
	ErrAlreadyCollectedInternshipReport = errors.New("Pemustaka already collected internship report")

	// error fields are not equal
	ErrFieldsAreNotEqual = errors.New("The fields must be equal")

	// error internal server error
	ErrInternalServerError = errors.New("Internal server error")

	// error invalid token header
	ErrInvalidTokenHeader = errors.New("Invalid token header")

	// error forbidden access
	ErrForbiddenAccess = errors.New("Access forbidden")

	// error unauthorized user
	ErrUnAuthorized = errors.New("User unauthorized")

	// error account status not activated
	ErrAccountNotActivated = errors.New("Account not activated yet")

	// error head of library have not upload the signature yet
	ErrHeadOfLibraryNotUploadSignature = errors.New("Head of library is not upload the signature yet")

	// error pemustaka not collected final project report
	ErrNotCollectedFinalProject = errors.New("Pemustka is not collected final project report yet")

	// error pemustaka not collected internship report
	ErrNotCollectedInternshipReport = errors.New("Pemustaka is not collected internship report yet")
)
