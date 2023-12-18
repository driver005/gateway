package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type ProductService struct {
	ctx                       context.Context
	repo                      *repository.ProductRepo
	productOptionRepository   *repository.ProductOptionRepo
	productVariantRepository  *repository.ProductVariantRepo
	productTypeRepository     *repository.ProductTypeRepo
	productTagRepository      *repository.ProductTagRepo
	imageRepository           *repository.ImageRepo
	productCategoryRepository *repository.ProductCategoryRepo
}

func NewProductService(
	ctx context.Context,
	repo *repository.ProductRepo,
	productOptionRepository *repository.ProductOptionRepo,
	productVariantRepository *repository.ProductVariantRepo,
	productTypeRepository *repository.ProductTypeRepo,
	productTagRepository *repository.ProductTagRepo,
	imageRepository *repository.ImageRepo,
	productCategoryRepository *repository.ProductCategoryRepo,
) *ProductService {
	return &ProductService{
		ctx,
		repo,
		productOptionRepository,
		productVariantRepository,
		productTypeRepository,
		productTagRepository,
		imageRepository,
		productCategoryRepository,
	}
}
