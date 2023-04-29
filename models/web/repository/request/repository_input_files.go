package request

import "mime/multipart"

type RepositoryInputFiles struct {
	ValidityPage        *multipart.FileHeader
	CoverAndListContent *multipart.FileHeader
	ChpOne              *multipart.FileHeader
	ChpTwo              *multipart.FileHeader
	ChpThree            *multipart.FileHeader
	ChpFour             *multipart.FileHeader
	ChpFive             *multipart.FileHeader
	Bibliography        *multipart.FileHeader
}
