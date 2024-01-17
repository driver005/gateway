package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ProductCollectionService struct {
	ctx context.Context
	r   Registry
}

func NewProductCollectionService(
	r Registry,
) *ProductCollectionService {
	return &ProductCollectionService{
		context.Background(),
		r,
	}
}

func (s *ProductCollectionService) SetContext(context context.Context) *ProductCollectionService {
	s.ctx = context
	return s
}

func (s *ProductCollectionService) Retrieve(collectionId uuid.UUID, config *sql.Options) (*models.ProductCollection, *utils.ApplictaionError) {
	if collectionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"collectionId" must be defined`,
			nil,
		)
	}
	var res *models.ProductCollection

	query := sql.BuildQuery(models.ProductCollection{Model: core.Model{Id: collectionId}}, config)

	if err := s.r.ProductCollectionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductCollectionService) RetrieveByHandle(collectionHandle string, config *sql.Options) (*models.ProductCollection, *utils.ApplictaionError) {
	if collectionHandle == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"collectionHandle" must be defined`,
			nil,
		)
	}

	var res *models.ProductCollection

	query := sql.BuildQuery(models.ProductCollection{Handle: collectionHandle}, config)

	if err := s.r.ProductCollectionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductCollectionService) List(selector models.ProductCollection, config *sql.Options, q *string, discountConditionId uuid.UUID) ([]models.ProductCollection, *utils.ApplictaionError) {
	collections, _, err := s.ListAndCount(selector, config, q, discountConditionId)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (s *ProductCollectionService) ListAndCount(selector models.ProductCollection, config *sql.Options, q *string, discountConditionId uuid.UUID) ([]models.ProductCollection, *int64, *utils.ApplictaionError) {
	var res []models.ProductCollection

	if q != nil {
		v := sql.ILike(*q)
		selector.Title = v
		selector.Handle = v
	}

	config.Relations = []string{}

	query := sql.BuildQuery(selector, config)

	if discountConditionId != uuid.Nil {
		return s.r.ProductCollectionRepository().FindAndCountByDiscountConditionId(s.ctx, discountConditionId, query)
	}

	count, err := s.r.ProductCollectionRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *ProductCollectionService) Create(data *types.CreateProductCollection) (*models.ProductCollection, *utils.ApplictaionError) {
	model := &models.ProductCollection{
		Model: core.Model{
			Metadata: data.Metadata,
		},
		Title:  data.Title,
		Handle: data.Handle,
	}

	if err := s.r.ProductCollectionRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}
	// s.eventBus_.withTransaction(manager).emit(ProductCollectionService.Events.CREATED, map[string]interface{}{"id": productCollection.id})
	return model, nil
}

func (s *ProductCollectionService) Update(collectionId uuid.UUID, data *types.UpdateProductCollection) (*models.ProductCollection, *utils.ApplictaionError) {
	productCollection, err := s.Retrieve(collectionId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	productCollection.Metadata = data.Metadata
	productCollection.Title = data.Title
	productCollection.Handle = data.Handle

	if err := s.r.ProductCollectionRepository().Save(s.ctx, productCollection); err != nil {
		return nil, err
	}

	// s.eventBus_.withTransaction(manager).emit(ProductCollectionService.Events.UPDATED, map[string]interface{}{"id": productCollection.id})
	return productCollection, nil
}

func (s *ProductCollectionService) Delete(collectionId uuid.UUID) *utils.ApplictaionError {
	productCollection, err := s.Retrieve(collectionId, &sql.Options{})
	if err != nil {
		return nil
	}

	if productCollection == nil {
		return nil
	}

	if err := s.r.ProductCollectionRepository().SoftRemove(s.ctx, productCollection); err != nil {
		return nil
	}
	// s.eventBus_.withTransaction(manager).emit(ProductCollectionService.Events.DELETED, map[string]interface{}{"id": productCollection.id})
	return nil
}

func (s *ProductCollectionService) AddProducts(collectionId uuid.UUID, productIds uuid.UUIDs) (*models.ProductCollection, *utils.ApplictaionError) {
	collection, err := s.Retrieve(collectionId, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return nil, err
	}

	_, err = s.r.ProductRepository().BulkAddToCollection(productIds, collection.Id)
	if err != nil {
		return nil, err
	}

	productCollection, err := s.Retrieve(collection.Id, &sql.Options{Relations: []string{"products"}})
	if err != nil {
		return nil, err
	}

	// s.eventBus_.withTransaction(manager).emit(ProductCollectionService.Events.PRODUCTS_ADDED, map[string]interface{}{"productCollection": productCollection, "productIds": productIds})
	return productCollection, nil
}

func (s *ProductCollectionService) RemoveProducts(collectionId uuid.UUID, productIds uuid.UUIDs) *utils.ApplictaionError {
	collection, err := s.Retrieve(collectionId, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return err
	}

	_, err = s.r.ProductRepository().BulkAddToCollection(productIds, collection.Id)
	if err != nil {
		return err
	}

	// s.eventBus_.withTransaction(manager).emit(ProductCollectionService.Events.PRODUCTS_REMOVED, map[string]interface{}{"productCollection": productCollection, "productIds": productIds})
	return nil
}
