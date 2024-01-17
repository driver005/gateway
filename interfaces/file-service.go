package interfaces

import (
	"io"
	"mime/multipart"

	"github.com/driver005/gateway/utils"
)

type FileServiceUploadResult struct {
	URL string
	Key string
}

type FileServiceGetUploadStreamResult struct {
	WriteStream io.Writer
	Promise     chan struct{}
	URL         string
	FileKey     string
}

type GetUploadedFileType struct {
	FileKey   string
	IsPrivate bool
}

type DeleteFileType struct {
	FileKey string
}

type UploadStreamDescriptorType struct {
	Name      string
	Ext       string
	IsPrivate bool
}

type IFileService interface {
	Upload(file *multipart.FileHeader) (*FileServiceUploadResult, *utils.ApplictaionError)
	UploadProtected(file *multipart.FileHeader) (*FileServiceUploadResult, *utils.ApplictaionError)
	Delete(fileData DeleteFileType) *utils.ApplictaionError
	GetUploadStreamDescriptor(fileData UploadStreamDescriptorType) (*FileServiceGetUploadStreamResult, *utils.ApplictaionError)
	GetDownloadStream(fileData GetUploadedFileType) (io.ReadCloser, *utils.ApplictaionError)
	GetPresignedDownloadUrl(fileData GetUploadedFileType) (string, *utils.ApplictaionError)
}
