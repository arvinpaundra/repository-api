package utils

import "fmt"

func GetMemberCode(totalPemustaka int, departementCode string) string {
	// current totalPemustaka + 1 (current user that trying to register)
	totalPemustaka += 1

	if totalPemustaka < 10 {
		code := fmt.Sprintf("000%d-%s", totalPemustaka, departementCode)

		return code
	} else if totalPemustaka < 100 {
		code := fmt.Sprintf("00%d-%s", totalPemustaka, departementCode)

		return code
	} else if totalPemustaka < 1000 {
		code := fmt.Sprintf("0%d-%s", totalPemustaka, departementCode)

		return code
	} else {
		code := fmt.Sprintf("%d-%s", totalPemustaka, departementCode)

		return code
	}
}
