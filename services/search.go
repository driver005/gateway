package services

import (
	"context"
)

type DefaultSearchService struct {
	ctx context.Context
	r   Registry
}

func NewDefaultSearchService(
	r Registry,
) *DefaultSearchService {
	return &DefaultSearchService{
		context.Background(),
		r,
	}
}

func (s *DefaultSearchService) SetContext(context context.Context) *DefaultSearchService {
	s.ctx = context
	return s
}

func (s *DefaultSearchService) CreateIndex(indexName string, options interface{}) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: createIndex must be overridden by a child class")
}

func (s *DefaultSearchService) GetIndex(indexName string) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: getIndex must be overridden by a child class")
}

func (s *DefaultSearchService) AddDocuments(indexName string, documents interface{}, types string) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: addDocuments must be overridden by a child class")
}

func (s *DefaultSearchService) ReplaceDocuments(indexName string, documents interface{}, types string) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: replaceDocuments must be overridden by a child class")
}

func (s *DefaultSearchService) DeleteDocument(indexName string, document_id interface{}) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: deleteDocument must be overridden by a child class")
}

func (s *DefaultSearchService) DeleteAllDocuments(indexName string) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: deleteAllDocuments must be overridden by a child class")
}

func (s *DefaultSearchService) Search(indexName string, query interface{}, options interface{}) map[string]interface{} {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: search must be overridden a the child class")
	return map[string]interface{}{
		"hits": []interface{}{},
	}
}

func (s *DefaultSearchService) UpdateSettings(indexName string, settings interface{}) {
	s.r.Context().Logger.Warn(s.ctx, "This is an empty method: updateSettings must be overridden by a child class")
}
