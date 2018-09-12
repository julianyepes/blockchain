package controllers

import (
	"fmt"
	"strings"
	"time"
)

const hashPrintBytes = 4

// formatShort is used to print out up to "hashPrintBytes"
// from a a slice of bytes
func formatShort(b []byte) string {
	if len(b) > hashPrintBytes {
		return fmt.Sprintf("%X", b[:hashPrintBytes])
	}
	return string(b)
}

// formatAsDate takes a int64 representing a unix time and
// returns the date in the format YYYY-MM-DD HH:MM:SS
func formatAsDate(t int64) string {
	date := fmt.Sprintf("%s", time.Unix(t, 0))
	splitDate := strings.Split(date, " ")
	return strings.Join([]string{splitDate[0], splitDate[1]}, " ")
}
