package utils

import "fmt"

func GetMemberCode(totalPemustaka int, departementCode string) string {
	// current totalPemustaka + 1 (current user that trying to register)
	totalPemustaka += 1

	code := fmt.Sprintf("%d%s", totalPemustaka, departementCode)

	return code
}
