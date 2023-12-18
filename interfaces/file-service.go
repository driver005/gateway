package interfaces

import (
	"io"
	"mime/multipart"
)

type FileServiceUploadResult struct {
	// define the structure of FileServiceUploadResult
}

type DeleteFileType struct {
	// define the structure of DeleteFileType
}

type UploadStreamDescriptorType struct {
	// define the structure of UploadStreamDescriptorType
}

type FileServiceGetUploadStreamResult struct {
	// define the structure of FileServiceGetUploadStreamResult
}

type GetUploadedFileType struct {
	// define the structure of GetUploadedFileType
}

type IFileService interface {
	Upload(file *multipart.FileHeader) (*FileServiceUploadResult, error)
	UploadProtected(file *multipart.FileHeader) (*FileServiceUploadResult, error)
	Delete(fileData DeleteFileType) error
	GetUploadStreamDescriptor(fileData UploadStreamDescriptorType) (*FileServiceGetUploadStreamResult, error)
	GetDownloadStream(fileData GetUploadedFileType) (io.ReadCloser, error)
	GetPresignedDownloadUrl(fileData GetUploadedFileType) (string, error)
}

func IsFileService(object interface{}) bool {
	_, ok := object.(IFileService)
	return ok
}
