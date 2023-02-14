package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GetFilename() string {
	// generate filename from uuid with removed hyphens
	filename := strings.Replace(uuid.NewString(), "-", "", -1)

	return filename
}
