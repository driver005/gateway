package services

import (
	"context"
	"reflect"
	"testing"
)

func TestNewDefaultSearchService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *DefaultSearchService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultSearchService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultSearchService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultSearchService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
		want *DefaultSearchService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultSearchService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultSearchService_CreateIndex(t *testing.T) {
	type args struct {
		indexName string
		options   interface{}
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.CreateIndex(tt.args.indexName, tt.args.options)
		})
	}
}

func TestDefaultSearchService_GetIndex(t *testing.T) {
	type args struct {
		indexName string
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.GetIndex(tt.args.indexName)
		})
	}
}

func TestDefaultSearchService_AddDocuments(t *testing.T) {
	type args struct {
		indexName string
		documents interface{}
		types     string
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddDocuments(tt.args.indexName, tt.args.documents, tt.args.types)
		})
	}
}

func TestDefaultSearchService_ReplaceDocuments(t *testing.T) {
	type args struct {
		indexName string
		documents interface{}
		types     string
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.ReplaceDocuments(tt.args.indexName, tt.args.documents, tt.args.types)
		})
	}
}

func TestDefaultSearchService_DeleteDocument(t *testing.T) {
	type args struct {
		indexName   string
		document_id interface{}
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DeleteDocument(tt.args.indexName, tt.args.document_id)
		})
	}
}

func TestDefaultSearchService_DeleteAllDocuments(t *testing.T) {
	type args struct {
		indexName string
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DeleteAllDocuments(tt.args.indexName)
		})
	}
}

func TestDefaultSearchService_Search(t *testing.T) {
	type args struct {
		indexName string
		query     interface{}
		options   map[string]interface{}
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Search(tt.args.indexName, tt.args.query, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultSearchService.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultSearchService_UpdateSettings(t *testing.T) {
	type args struct {
		indexName string
		settings  interface{}
	}
	tests := []struct {
		name string
		s    *DefaultSearchService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.UpdateSettings(tt.args.indexName, tt.args.settings)
		})
	}
}
