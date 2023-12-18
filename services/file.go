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

func NewDefaultFileService(
	ctx context.Context,
) *DefaultFileService {
	return &DefaultFileService{
		ctx,
	}
}

func (s *DefaultFileService) upload(fileData io.Reader) (*interfaces.FileServiceUploadResult, error) {
	return nil, errors.New("Please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) uploadProtected(fileData io.Reader) (*interfaces.FileServiceUploadResult, error) {
	return nil, errors.New("Please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) delete(fileData map[string]interface{}) error {
	return errors.New("Please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) getUploadStreamDescriptor(fileData interfaces.UploadStreamDescriptorType) (*interfaces.FileServiceGetUploadStreamResult, error) {
	return nil, errors.New("Please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) getDownloadStream(fileData interfaces.GetUploadedFileType) (io.Reader, error) {
	return nil, errors.New("Please add a file service plugin in order to manipulate files in Medusa")
}

func (s *DefaultFileService) getPresignedDownloadUrl(fileData interfaces.GetUploadedFileType) (string, error) {
	return "", errors.New("Please add a file service plugin in order to manipulate files in Medusa")
}
