package services

import (
	"context"
	"fmt"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ProductCategoryService struct {
	ctx context.Context
	r   Registry
}

func NewProductCategoryService(
	r Registry,
) *ProductCategoryService {
	return &ProductCategoryService{
		context.Background(),
		r,
	}
}

func (s *ProductCategoryService) SetContext(context context.Context) *ProductCategoryService {
	s.ctx = context
	return s
}

func (s *ProductCategoryService) List(selector models.ProductCategory, config sql.Options, q *string) ([]models.ProductCategory, *utils.ApplictaionError) {
	collections, _, err := s.ListAndCount(selector, config, q)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (s *ProductCategoryService) ListAndCount(selector models.ProductCategory, config sql.Options, q *string) ([]models.ProductCategory, *int64, *utils.ApplictaionError) {
	includeDescendantsTree := true
	query := sql.BuildQuery(selector, config)

	res, count, err := s.r.ProductCategoryRepository().GetFreeTextSearchResultsAndCount(s.ctx, query, q, includeDescendantsTree)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *ProductCategoryService) Retrieve(selector models.ProductCategory, config sql.Options, q *string) (*models.ProductCategory, error) {
	query := sql.BuildQuery(selector, config)

	res, err := s.r.ProductCategoryRepository().FindOneWithDescendants(s.ctx, query, q)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductCategoryService) RetrieveById(id uuid.UUID, config sql.Options, q *string) (*models.ProductCategory, error) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(models.ProductCategory{Model: core.Model{Id: id}}, config, q)
}

func (s *ProductCategoryService) RetrieveByHandle(handle string, config sql.Options, q *string) (*models.ProductCategory, error) {
	if handle == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}

	return s.Retrieve(models.ProductCategory{Handle: handle}, config, q)
}

func (s *ProductCategoryService) Create(data *models.ProductCategory) (*models.ProductCategory, error) {
	query := sql.BuildQuery(models.ProductCategory{Model: core.Model{Id: data.ParentCategoryId.UUID}}, sql.Options{})
	siblingCount, err := s.r.ProductCategoryRepository().CountBy(s.ctx, []string{"parent_category_id"}, query)
	if err != nil {
		return nil, err
	}

	data.Rank = *siblingCount
	if err := s.transformParentIdToEntity(data); err != nil {
		return nil, err
	}

	if err := s.r.ProductCategoryRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}

	// err = s.eventBusService_.withTransaction(manager).emit(ProductCategoryService.Events.CREATED, map[string]interface{}{
	// 	"id": productCategory.ID,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return data, nil
}

func (s *ProductCategoryService) Update(productCategoryId uuid.UUID, update *models.ProductCategory) (*models.ProductCategory, error) {
	productCategory, err := s.RetrieveById(productCategoryId, sql.Options{}, nil)
	if err != nil {
		return nil, err
	}

	conditions := s.fetchReorderConditions(productCategory, update, false)
	if conditions.ShouldChangeRank || conditions.ShouldChangeParent {
		update.Rank = types.TempReorderRank
	}

	err = s.transformParentIdToEntity(update)
	if err != nil {
		return productCategory, err
	}

	update.Id = productCategory.Id

	if err := s.r.ProductCategoryRepository().Save(s.ctx, update); err != nil {
		return nil, err
	}

	err = s.performReordering(conditions)
	if err != nil {
		return productCategory, err
	}

	// err = s.eventBusService_.withTransaction(manager).emit(ProductCategoryService.Events.UPDATED, map[string]interface{}{
	// 	"id": productCategory.ID,
	// })
	// if err != nil {
	// 	return productCategory, err
	// }

	return nil, err
}

func (s *ProductCategoryService) Delete(productCategoryId uuid.UUID) error {
	productCategory, err := s.RetrieveById(productCategoryId, sql.Options{
		Relations: []string{"category_children"},
	}, nil)
	if err != nil {
		return err
	}

	if productCategory == nil {
		return nil
	}

	conditions := s.fetchReorderConditions(productCategory, &models.ProductCategory{
		ParentCategoryId: productCategory.ParentCategoryId,
		Rank:             productCategory.Rank,
	}, true)

	if len(productCategory.CategoryChildren) > 0 {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf("Deleting models.ProductCategory (%s) with category children is not allowed", productCategoryId),
			"500",
			nil,
		)
	}

	if err := s.r.ProductCategoryRepository().Delete(s.ctx, productCategory); err != nil {
		return err
	}

	err = s.performReordering(conditions)
	if err != nil {
		return err
	}

	// err = s.eventBusService_.withTransaction(manager).emit(ProductCategoryService.Events.DELETED, map[string]interface{}{
	// 	"id": productCategory.ID,
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *ProductCategoryService) AddProducts(productCategoryId uuid.UUID, productIDs uuid.UUIDs) error {
	if err := s.r.ProductCategoryRepository().AddProducts(productCategoryId, productIDs); err != nil {
		return err
	}

	return nil
}

func (s *ProductCategoryService) RemoveProducts(productCategoryId uuid.UUID, productIDs uuid.UUIDs) error {
	if err := s.r.ProductCategoryRepository().RemoveProducts(productCategoryId, productIDs); err != nil {
		return err
	}

	return nil
}

func (s *ProductCategoryService) fetchReorderConditions(productCategory *models.ProductCategory, input *models.ProductCategory, shouldDeleteElement bool) types.ReorderConditions {
	originalParentId := productCategory.ParentCategoryId
	targetParentId := input.ParentCategoryId
	originalRank := productCategory.Rank
	targetRank := input.Rank
	shouldChangeParent := targetParentId.UUID != uuid.Nil && targetParentId != originalParentId
	shouldChangeRank := shouldChangeParent || (targetRank != 0 && originalRank != targetRank)
	return types.ReorderConditions{
		TargetCategoryId:    productCategory.Id,
		OriginalParentId:    originalParentId.UUID,
		TargetParentId:      targetParentId.UUID,
		OriginalRank:        originalRank,
		TargetRank:          targetRank,
		ShouldChangeParent:  shouldChangeParent,
		ShouldChangeRank:    shouldChangeRank,
		ShouldIncrementRank: false,
		ShouldDeleteElement: shouldDeleteElement,
	}
}

func (s *ProductCategoryService) performReordering(conditions types.ReorderConditions) error {
	shouldChangeParent := conditions.ShouldChangeParent
	shouldChangeRank := conditions.ShouldChangeRank
	shouldDeleteElement := conditions.ShouldDeleteElement

	if !(shouldChangeParent || shouldChangeRank || shouldDeleteElement) {
		return nil
	}

	if shouldChangeParent {
		err := s.shiftSiblings(types.ReorderConditions{
			TargetCategoryId:    conditions.TargetCategoryId,
			OriginalParentId:    conditions.OriginalParentId,
			TargetParentId:      conditions.OriginalParentId,
			OriginalRank:        conditions.TargetRank,
			TargetRank:          conditions.OriginalRank,
			ShouldChangeParent:  true,
			ShouldChangeRank:    false,
			ShouldIncrementRank: false,
			ShouldDeleteElement: false,
		})
		if err != nil {
			return err
		}
	}

	if shouldChangeParent && shouldChangeRank {
		err := s.shiftSiblings(types.ReorderConditions{
			TargetCategoryId:    conditions.TargetCategoryId,
			OriginalParentId:    conditions.TargetParentId,
			TargetParentId:      conditions.TargetParentId,
			OriginalRank:        conditions.TargetRank,
			TargetRank:          conditions.TargetRank,
			ShouldChangeParent:  true,
			ShouldChangeRank:    true,
			ShouldIncrementRank: true,
			ShouldDeleteElement: false,
		})
		if err != nil {
			return err
		}
	}

	if (!shouldChangeParent && shouldChangeRank) || shouldDeleteElement {
		err := s.shiftSiblings(types.ReorderConditions{
			TargetCategoryId:    conditions.TargetCategoryId,
			OriginalParentId:    conditions.OriginalParentId,
			TargetParentId:      conditions.OriginalParentId,
			OriginalRank:        conditions.OriginalRank,
			TargetRank:          conditions.TargetRank,
			ShouldChangeParent:  false,
			ShouldChangeRank:    true,
			ShouldIncrementRank: false,
			ShouldDeleteElement: shouldDeleteElement,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ProductCategoryService) shiftSiblings(conditions types.ReorderConditions) error {
	shouldIncrementRank := conditions.ShouldIncrementRank
	targetRank := conditions.TargetRank
	shouldChangeParent := conditions.ShouldChangeParent
	originalRank := conditions.OriginalRank
	targetParentId := conditions.TargetParentId
	targetCategoryId := conditions.TargetCategoryId
	shouldDeleteElement := conditions.ShouldDeleteElement

	query := sql.BuildQuery(models.ProductCategory{
		ParentCategoryId: uuid.NullUUID{UUID: targetParentId},
		Model:            core.Model{Id: targetCategoryId},
	}, sql.Options{
		Null: []string{"parent_category_id"},
		Not:  []string{"id"},
	})
	siblingCount, err := s.r.ProductCategoryRepository().CountBy(s.ctx, []string{"parent_category_id", "id"}, query)
	if err != nil {
		return err
	}

	var targetCategory *models.ProductCategory

	query = sql.BuildQuery(models.ProductCategory{
		Model:            core.Model{Id: targetCategoryId},
		ParentCategoryId: uuid.NullUUID{UUID: targetParentId},
		Rank:             types.TempReorderRank,
	}, sql.Options{
		Null: []string{"parent_category_id"},
	})

	if err := s.r.ProductCategoryRepository().FindOne(s.ctx, targetCategory, query); err != nil {
		return err
	}

	if targetRank == 0 || targetRank > *siblingCount {
		targetRank = *siblingCount
	}

	var rankCondition sql.Specification
	if shouldChangeParent || shouldDeleteElement {
		rankCondition = sql.GreaterOrEqual("rank", targetRank)
	} else if originalRank > targetRank {
		shouldIncrementRank = true
		rankCondition = sql.Equal("rank", targetRank)
	} else {
		shouldIncrementRank = false
		rankCondition = sql.Equal("rank", targetRank)
	}

	var siblingsToShift []models.ProductCategory

	query = sql.BuildQuery(models.ProductCategory{
		Model:            core.Model{Id: targetCategoryId},
		ParentCategoryId: uuid.NullUUID{UUID: targetParentId},
	}, sql.Options{
		Specification: []sql.Specification{rankCondition},
		Not:           []string{"id"},
		Null:          []string{"parent_category_id"},
	})

	if err := s.r.ProductCategoryRepository().Find(s.ctx, siblingsToShift, query); err != nil {
		return err
	}

	for _, sibling := range siblingsToShift {
		if sibling.Id == targetCategoryId {
			continue
		}

		if shouldIncrementRank {
			sibling.Rank++
		} else {
			sibling.Rank--
		}

		if err := s.r.ProductCategoryRepository().Save(s.ctx, &sibling); err != nil {
			return err
		}
	}

	if targetCategory == nil {
		return nil
	}

	targetCategory.Rank = targetRank
	if err := s.r.ProductCategoryRepository().Save(s.ctx, targetCategory); err != nil {
		return err
	}

	return nil
}

func (s *ProductCategoryService) transformParentIdToEntity(update *models.ProductCategory) error {
	parentCategoryId := update.ParentCategoryId
	if parentCategoryId.UUID == uuid.Nil {
		return nil
	}

	parentCategory, err := s.RetrieveById(parentCategoryId.UUID, sql.Options{}, nil)
	if err != nil {
		return err
	}

	update.ParentCategory = parentCategory

	return nil
}
