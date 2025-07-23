package service

import "mime/multipart"

type FileService interface {
	UploadImage(file *multipart.FileHeader, uploadDir string) (string, error)
	ValidateImageType(contentType string) bool
	DeleteFile(filePath string) error
}
