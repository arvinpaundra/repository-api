package utils

import "errors"

var (
	// error authentication failed
	ErrAuthenticationFailed = errors.New("wrong email or password")

	// error user not found
	ErrUserNotFound = errors.New("user not found")

	// ErrCategoryNotFound error category not found
	ErrCategoryNotFound = errors.New("category not found")

	// error collection not found
	ErrCollectionNotFound = errors.New("collection not found")

	// error role not found
	ErrRoleNotFound = errors.New("role not found")

	// error study program not found
	ErrStudyProgramNotFound = errors.New("study program not found")

	// error pemustaka not found
	ErrPemustakaNotFound = errors.New("pemustaka not found")

	// error departement not found
	ErrDepartementNotFound = errors.New("departement not found")

	// error staff not found
	ErrStaffNotFound = errors.New("staff not found")

	// error repository not found
	ErrRepositoryNotFound = errors.New("repository not found")

	// error request access not found
	ErrRequestAccessNotFound = errors.New("request access not found")

	// error email already used
	ErrEmailAlreadyUsed = errors.New("email already used")

	// error email not found
	ErrEmailNotFound = errors.New("email not found")

	// error invalid password length
	ErrInvalidPasswordLength = errors.New("minimum password length are 6 characters long")

	// error token has been expired
	ErrTokenExpired = errors.New("token has been expired")

	// error token has been sent
	ErrTokenHasBeenSent = errors.New("token has been sent")

	// error token not match
	ErrTokenNotMatch = errors.New("token not matched")

	// error link has been expired
	ErrLinkExpired = errors.New("link has been expired")

	// error waiting for acceptance to activate pemustaka account
	ErrWaitingForAcceptance = errors.New("account is waiting for acceptance")

	// error internal server error
	ErrInternalServerError = errors.New("internal server error")
)
