package services

import (
	"context"
	"errors"
	"io"

	"github.com/driver005/gateway/interfaces"
)

type DefaultFileService struct {
	ctx context.Context
}

func NewDefaultFileService() *DefaultFileService {
	return &DefaultFileService{
		context.Background(),
	}
}

func (s *DefaultFileService) SetContext(context context.Context) *DefaultFileService {
	s.ctx = context
	return s
}

func (s *DefaultFileService) Upload(fileData io.Reader) (*interfaces.FileServiceUploadResult, error) {
	return nil, errors.New("please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) UploadProtected(fileData io.Reader) (*interfaces.FileServiceUploadResult, error) {
	return nil, errors.New("please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) Delete(fileData map[string]interface{}) error {
	return errors.New("please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) GetUploadStreamDescriptor(fileData interfaces.UploadStreamDescriptorType) (*interfaces.FileServiceGetUploadStreamResult, error) {
	return nil, errors.New("please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) GetDownloadStream(fileData interfaces.GetUploadedFileType) (io.Reader, error) {
	return nil, errors.New("please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) GetPresignedDownloadUrl(fileData interfaces.GetUploadedFileType) (string, error) {
	return "", errors.New("please add a file service plugin in order to manipulate files in Medusa")
}
