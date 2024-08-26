package Utils

import (
	"path/filepath"
	"strconv"
	"time"
)

// GenerateUniqueFileName generates a unique file name based on the current timestamp and original extension.
func GenerateUniqueFileName(originalFileName string) string {
	extension := filepath.Ext(originalFileName)
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	return timestamp + extension
}
