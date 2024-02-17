package services

import (
	"context"
	"io"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
)

func TestNewDefaultFileService(t *testing.T) {
	tests := []struct {
		name string
		want *DefaultFileService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultFileService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultFileService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFileService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *DefaultFileService
		args args
		want *DefaultFileService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultFileService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFileService_Upload(t *testing.T) {
	type args struct {
		fileData io.Reader
	}
	tests := []struct {
		name    string
		s       *DefaultFileService
		args    args
		want    *interfaces.FileServiceUploadResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Upload(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultFileService.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultFileService.Upload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFileService_UploadProtected(t *testing.T) {
	type args struct {
		fileData io.Reader
	}
	tests := []struct {
		name    string
		s       *DefaultFileService
		args    args
		want    *interfaces.FileServiceUploadResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UploadProtected(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultFileService.UploadProtected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultFileService.UploadProtected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFileService_Delete(t *testing.T) {
	type args struct {
		fileData map[string]interface{}
	}
	tests := []struct {
		name    string
		s       *DefaultFileService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Delete(tt.args.fileData); (err != nil) != tt.wantErr {
				t.Errorf("DefaultFileService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultFileService_GetUploadStreamDescriptor(t *testing.T) {
	type args struct {
		fileData interfaces.UploadStreamDescriptorType
	}
	tests := []struct {
		name    string
		s       *DefaultFileService
		args    args
		want    *interfaces.FileServiceGetUploadStreamResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUploadStreamDescriptor(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultFileService.GetUploadStreamDescriptor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultFileService.GetUploadStreamDescriptor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFileService_GetDownloadStream(t *testing.T) {
	type args struct {
		fileData interfaces.GetUploadedFileType
	}
	tests := []struct {
		name    string
		s       *DefaultFileService
		args    args
		want    io.Reader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetDownloadStream(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultFileService.GetDownloadStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultFileService.GetDownloadStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFileService_GetPresignedDownloadUrl(t *testing.T) {
	type args struct {
		fileData interfaces.GetUploadedFileType
	}
	tests := []struct {
		name    string
		s       *DefaultFileService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPresignedDownloadUrl(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultFileService.GetPresignedDownloadUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultFileService.GetPresignedDownloadUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
