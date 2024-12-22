package Utils

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateUniqueFileName generates a unique file name based on the current timestamp and original extension.
func GenerateUniqueFileName(originalFileName string) string {
	extension := filepath.Ext(originalFileName)
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	return timestamp + extension
}

func UploadFile(ctx *gin.Context, formName string, uploadPath string) (string, error) {
	file, err := ctx.FormFile(formName)
	if err != nil {
		return "", err
	}

	uniqueFileName := GenerateUniqueFileName(file.Filename)

	savePath := filepath.Join(uploadPath, uniqueFileName)

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		return "", err
	}

	return deletePrefixPublic(savePath), nil
}

func deletePrefixPublic(path string) string {
	return strings.TrimPrefix(path, "public/")
}
