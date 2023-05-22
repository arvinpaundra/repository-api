package request

import "mime/multipart"

type RepositoryInputFiles map[string]*multipart.FileHeader
