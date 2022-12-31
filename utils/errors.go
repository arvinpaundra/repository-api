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

	// error email already used
	ErrEmailAlreadyUsed = errors.New("email already used")

	// error invalid password length
	ErrInvalidPasswordLength = errors.New("minimum password length are 6 characters long")

	// error internal server error
	ErrInternalServerError = errors.New("internal server error")
)
