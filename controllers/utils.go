package controllers

import "fmt"

const hashPrintBytes = 4

// formatShort is used to print out up to "hashPrintBytes"
// from a a slice of bytes
func formatShort(b []byte) string {
	if len(b) > hashPrintBytes {
		return fmt.Sprintf("%X", b[:hashPrintBytes])
	}
	return string(b)
}
