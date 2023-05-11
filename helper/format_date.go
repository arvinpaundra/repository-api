package helper

import (
	"fmt"
	"time"
)

func FormatDate(date time.Time) string {
	// define indonesian months
	months := [...]string{
		"Januari", "Februari", "Maret", "April",
		"Mei", "Juni", "Juli", "Agustus",
		"September", "Oktober", "November", "Desember",
	}

	// get month index (0-11)
	monthIndex := int(date.Month()) - 1

	// get the indonesian month name
	monthName := months[monthIndex]

	formattedDate := fmt.Sprintf("%02d %s %d", date.Day(), monthName, date.Year())

	return formattedDate
}
