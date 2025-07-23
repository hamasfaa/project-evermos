package impl

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hamasfaa/project-evermos/service"
)

type fileServiceImpl struct{}

func NewFileServiceImpl() service.FileService {
	return &fileServiceImpl{}
}

func (s *fileServiceImpl) ValidateImageType(contentType string) bool {
	allowedTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
	}

	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			return true
		}
	}
	return false
}

func (s *fileServiceImpl) UploadImage(file *multipart.FileHeader, uploadDir string) (string, error) {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
	filePath := filepath.Join(uploadDir, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy file content
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fmt.Sprintf("/uploads/toko/%s", filename), nil
}

func (s *fileServiceImpl) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}
