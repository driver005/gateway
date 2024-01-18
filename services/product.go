package services

import (
	"context"
	"reflect"
	"slices"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ProductService struct {
	ctx context.Context
	r   Registry
}

func NewProductService(
	r Registry,
) *ProductService {
	return &ProductService{
		context.Background(),
		r,
	}
}

func (s *ProductService) SetContext(context context.Context) *ProductService {
	s.ctx = context
	return s
}

func (s *ProductService) List(selector *types.FilterableProduct, config *sql.Options) ([]models.Product, *utils.ApplictaionError) {
	products, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) ListAndCount(selector *types.FilterableProduct, config *sql.Options) ([]models.Product, *int64, *utils.ApplictaionError) {
	hasSalesChannelsRelation := false
	for _, r := range config.Relations {
		if r == "sales_channels" {
			hasSalesChannelsRelation = true
			break
		}
	}

	feature := true

	if feature && hasSalesChannelsRelation {
		newRelations := []string{}
		for _, r := range config.Relations {
			if r != "sales_channels" {
				newRelations = append(newRelations, r)
			}
		}
		config.Relations = newRelations
	}

	var count *int64
	var products []models.Product

	query := sql.BuildQuery(selector, config)
	if config.Q != nil {
		p, c, err := s.r.ProductRepository().GetFreeTextSearchResultsAndCount(config.Q, query, config.Relations)
		if err != nil {
			return nil, nil, err
		}
		products = p
		count = c
	} else {
		c, err := s.r.ProductRepository().FindAndCount(s.ctx, products, query)
		if err != nil {
			return nil, nil, err
		}
		count = c
	}
	if feature && hasSalesChannelsRelation {
		_, err := s.DecorateProductsWithSalesChannels(products)
		if err != nil {
			return nil, nil, err
		}
	}
	return products, count, nil
}

func (s *ProductService) Count(selector models.Product) (*int64, *utils.ApplictaionError) {
	query := sql.BuildQuery(selector, &sql.Options{})

	count, err := s.r.ProductRepository().Count(s.ctx, query)
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (s *ProductService) RetrieveById(id uuid.UUID, config *sql.Options) (*models.Product, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			nil,
		)
	}
	return s.Retrieve(models.Product{Model: core.Model{Id: id}}, config)
}

func (s *ProductService) RetrieveByHandle(productHandle string, config *sql.Options) (*models.Product, *utils.ApplictaionError) {
	if productHandle == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"productHandle" must be defined`,
			nil,
		)
	}

	return s.Retrieve(models.Product{Handle: productHandle}, config)
}

func (s *ProductService) RetrieveByExternalId(externalId string, config *sql.Options) (*models.Product, *utils.ApplictaionError) {
	if externalId == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"externalId" must be defined`,
			nil,
		)
	}
	return s.Retrieve(models.Product{ExternalId: externalId}, config)
}

func (s *ProductService) Retrieve(selector models.Product, config *sql.Options) (*models.Product, *utils.ApplictaionError) {
	var res *models.Product

	query := sql.BuildQuery(selector, config)

	if err := s.r.ProductRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) RetrieveVariants(id uuid.UUID, config *sql.Options) ([]models.ProductVariant, *utils.ApplictaionError) {
	requiredRelations := []string{"variants"}

	config.Relations = append(config.Relations, requiredRelations...)

	product, err := s.RetrieveById(id, config)
	if err != nil {
		return nil, err
	}
	return product.Variants, nil
}

func (s *ProductService) FilterProductsBySalesChannel(productIds uuid.UUIDs, salesChannelId uuid.UUID, config *sql.Options) ([]models.Product, *utils.ApplictaionError) {
	requiredRelations := []string{"sales_channels"}

	config.Relations = append(config.Relations, requiredRelations...)
	config.Specification = append(config.Specification, sql.In("id", productIds))

	products, err := s.List(&types.FilterableProduct{}, config)
	if err != nil {
		return nil, err
	}
	productSalesChannelsMap := make(map[uuid.UUID][]models.SalesChannel)
	for _, product := range products {
		productSalesChannelsMap[product.Id] = product.SalesChannels
	}
	var filteredProducts []models.Product
	for _, product := range products {
		for _, sc := range productSalesChannelsMap[product.Id] {
			if sc.Id == salesChannelId {
				filteredProducts = append(filteredProducts, product)
				break
			}
		}
	}
	return filteredProducts, nil
}

func (s *ProductService) ListTypes() ([]models.ProductType, *utils.ApplictaionError) {
	var res []models.ProductType

	if err := s.r.ProductTypeRepository().Find(s.ctx, res, sql.Query{}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) ListTagsByUsage(take int) ([]models.ProductTag, *utils.ApplictaionError) {
	productTags, err := s.r.ProductTagRepository().ListTagsByUsage(take)
	if err != nil {
		return nil, err
	}
	return productTags, nil
}

func (s *ProductService) IsProductInSalesChannels(id uuid.UUID, salesChannelIds uuid.UUIDs) (bool, *utils.ApplictaionError) {
	product, err := s.RetrieveById(id, &sql.Options{Relations: []string{"sales_channels"}})
	if err != nil {
		return false, err
	}

	var productsSalesChannels uuid.UUIDs
	for _, channel := range product.SalesChannels {
		productsSalesChannels = append(productsSalesChannels, channel.Id)
	}
	for _, id := range productsSalesChannels {
		if slices.Contains(salesChannelIds, id) {
			return true, nil
		}
	}
	return false, nil
}

func (s *ProductService) Create(data *types.CreateProductInput) (*models.Product, *utils.ApplictaionError) {
	var err *utils.ApplictaionError
	var product *models.Product

	if !reflect.ValueOf(data.Thumbnail).IsZero() && len(data.Images) > 0 {
		product.Thumbnail = data.Images[0]
	}
	if !reflect.ValueOf(data.IsGiftcard).IsZero() {
		product.Discountable = false
	}

	if data.ProfileId != uuid.Nil {
		product.Profiles = []models.ShippingProfile{{Model: core.Model{Id: data.ProfileId}}}
	}
	if len(data.Images) > 0 {
		var images []models.Image
		for _, i := range data.Images {
			images = append(images, models.Image{
				Url: i,
			})
		}
		product.Images, err = s.r.ImageRepository().UpsertImages(images)
		if err != nil {
			return nil, err
		}
	}
	if len(data.Tags) > 0 {
		var tags []models.ProductTag
		for _, t := range data.Tags {
			tags = append(tags, models.ProductTag{
				Model: core.Model{
					Id: t.Id,
				},
				Value: t.Value,
			})
		}
		product.Tags, err = s.r.ProductTagRepository().UpsertTags(tags)
		if err != nil {
			return nil, err
		}
	}
	if data.Type != nil {
		ty, err := s.r.ProductTypeRepository().UpsertType(&models.ProductType{
			Model: core.Model{
				Id: data.Type.Id,
			},
			Value: data.Type.Value,
		})
		if err != nil {
			return nil, err
		}
		product.TypeId = uuid.NullUUID{UUID: ty.Id}
	}

	feature := true
	featurev2 := true

	if feature && !featurev2 {
		if len(data.SalesChannels) > 0 {
			product.SalesChannels = []models.SalesChannel{}
			var salesChannelIds uuid.UUIDs
			for _, sc := range data.SalesChannels {
				salesChannelIds = append(salesChannelIds, sc.Id)
			}
			for _, id := range salesChannelIds {
				product.SalesChannels = append(product.SalesChannels, models.SalesChannel{Model: core.Model{Id: id}})
			}
		}
	}
	if len(data.Categories) > 0 {
		product.Categories = []models.ProductCategory{}
		var categoryIds uuid.UUIDs
		for _, c := range data.Categories {
			categoryIds = append(categoryIds, c.Id)
		}
		for _, id := range categoryIds {
			product.Categories = append(product.Categories, models.ProductCategory{Model: core.Model{Id: id}})
		}
	}

	if featurev2 {
		if len(data.SalesChannels) > 0 {
			for _, sc := range data.SalesChannels {
				_, err = s.r.SalesChannelService().SetContext(s.ctx).AddProducts(sc.Id, uuid.UUIDs{product.Id})
				if err != nil {
					return nil, err
				}
			}
		}
	}
	product.Options = []models.ProductOption{}
	for _, option := range data.Options {
		if err := s.r.ProductOptionRepository().Save(s.ctx, &models.ProductOption{Title: option.Title}); err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		product.Options = append(product.Options, models.ProductOption{Title: option.Title})
	}
	if data.Variants != nil {
		var variants []types.CreateProductVariantInput
		for _, variant := range data.Variants {
			var options []types.ProductVariantOption
			for i, option := range variant.Options {
				options = append(options, types.ProductVariantOption{OptionId: product.Options[i].Id, Value: option.Value})
			}
			variant.Options = options
			variants = append(variants, variant)
		}
		product.Variants, err = s.r.ProductVariantService().SetContext(s.ctx).Create(product.Id, nil, variants)
		if err != nil {
			return nil, err
		}
	}

	if err := s.r.ProductRepository().Save(s.ctx, product); err != nil {
		return nil, err
	}

	result, err := s.RetrieveById(product.Id, &sql.Options{Relations: []string{"options"}})
	if err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(ProductServiceEventsCreated, map[string]interface{}{"id": result.ID})
	// if err != nil {
	// 	return nil, err
	// }
	return result, nil
}

func (s *ProductService) Update(id uuid.UUID, data *types.UpdateProductInput) (*models.Product, *utils.ApplictaionError) {
	relations := []string{"tags", "images"}

	feature := true
	featurev2 := true
	if feature {
		relations = append(relations, "sales_channels")
	}
	product, err := s.RetrieveById(id, &sql.Options{Relations: relations})
	if err != nil {
		return nil, err
	}

	if product.Thumbnail == "" && reflect.ValueOf(data.Thumbnail).IsZero() && len(data.Images) > 0 {
		data.Thumbnail = data.Images[0]
	}

	if data.Type != nil {
		t, err := s.r.ProductTypeRepository().UpsertType(&models.ProductType{
			Model: core.Model{
				Id: data.Type.Id,
			},
			Value: data.Type.Value,
		})
		if err != nil {
			return nil, err
		}
		product.Type = t
	}
	if len(data.Tags) > 0 {
		var tags []models.ProductTag
		for _, t := range data.Tags {
			tags = append(tags, models.ProductTag{
				Model: core.Model{
					Id: t.Id,
				},
				Value: t.Value,
			})
		}
		t, err := s.r.ProductTagRepository().UpsertTags(tags)
		if err != nil {
			return nil, err
		}
		product.Tags = t
	}
	if len(data.Categories) > 0 {
		product.Categories = []models.ProductCategory{}
		var categoryIds uuid.UUIDs
		for _, c := range data.Categories {
			categoryIds = append(categoryIds, c.Id)
		}
		for _, id := range categoryIds {
			data.Categories = append(data.Categories, types.CreateProductProductCategoryInput{Id: id})
		}
	}
	if feature && !featurev2 {
		if len(data.SalesChannels) > 0 {
			product.SalesChannels = []models.SalesChannel{}
			var salesChannelIds uuid.UUIDs
			for _, sc := range data.SalesChannels {
				salesChannelIds = append(salesChannelIds, sc.Id)
			}
			for _, id := range salesChannelIds {
				data.SalesChannels = append(data.SalesChannels, types.CreateProductProductSalesChannelInput{Id: id})
			}
		}
	}

	if len(data.Images) > 0 {
		var images []models.Image
		for _, i := range data.Images {
			images = append(images, models.Image{
				Url: i,
			})
		}
		i, err := s.r.ImageRepository().UpsertImages(images)
		if err != nil {
			return nil, err
		}
		product.Images = i
	}

	product.Metadata = utils.MergeMaps(product.Metadata, data.Metadata)

	var options []models.ProductOption
	for _, o := range data.Options {
		options = append(options, models.ProductOption{
			Title: o.Title,
		})
	}

	// if len(data.Variants) > 0 {
	// 	var variants []models.ProductVariant
	// 	for _, v := range data.Variants {
	// 		variants = append(variants, models.ProductVariant{

	// 		})
	// 	}
	// 	product.Variants = variants
	// }

	if !reflect.ValueOf(data.Title).IsZero() {
		product.Title = data.Title
	}
	if !reflect.ValueOf(data.Subtitle).IsZero() {
		product.Subtitle = data.Subtitle
	}
	if !reflect.ValueOf(data.ProfileId).IsZero() {
		product.ProfileId = uuid.NullUUID{UUID: data.ProfileId}
	}
	if !reflect.ValueOf(data.Description).IsZero() {
		product.Description = data.Description
	}
	if !reflect.ValueOf(data.IsGiftcard).IsZero() {
		product.IsGiftcard = data.IsGiftcard
	}
	if !reflect.ValueOf(data.Discountable).IsZero() {
		product.Discountable = data.Discountable
	}
	if !reflect.ValueOf(data.Thumbnail).IsZero() {
		product.Thumbnail = data.Thumbnail
	}
	if !reflect.ValueOf(data.Handle).IsZero() {
		product.Handle = data.Handle
	}
	if !reflect.ValueOf(data.Status).IsZero() {
		product.Status = data.Status
	}
	if !reflect.ValueOf(data.CollectionId).IsZero() {
		product.CollectionId = uuid.NullUUID{UUID: data.CollectionId}
	}
	if !reflect.ValueOf(data.Options).IsZero() {
		product.Options = options
	}
	if !reflect.ValueOf(data.Weight).IsZero() {
		product.Weight = data.Weight
	}
	if !reflect.ValueOf(data.Length).IsZero() {
		product.Length = data.Length
	}
	if !reflect.ValueOf(data.Height).IsZero() {
		product.Height = data.Height
	}
	if !reflect.ValueOf(data.Width).IsZero() {
		product.Width = data.Width
	}
	if !reflect.ValueOf(data.HSCode).IsZero() {
		product.HsCode = data.HSCode
	}
	if !reflect.ValueOf(data.OriginCountry).IsZero() {
		product.OriginCountry = data.OriginCountry
	}
	if !reflect.ValueOf(data.MIdCode).IsZero() {
		product.MIdCode = data.MIdCode
	}
	if !reflect.ValueOf(data.Material).IsZero() {
		product.Material = data.Material
	}
	if data.Metadata != nil {
		product.Metadata = utils.MergeMaps(product.Metadata, data.Metadata)
	}
	if !reflect.ValueOf(data.ExternalId).IsZero() {
		product.ExternalId = data.ExternalId
	}

	if err = s.r.ProductRepository().Save(s.ctx, product); err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(ProductServiceEventsUpdated, map[string]interface{}{"id": product.ID, "fields": Object.keys(Update)})
	// if err != nil {
	// 	return nil, err
	// }
	return product, nil
}

func (s *ProductService) Delete(id uuid.UUID) *utils.ApplictaionError {
	product, err := s.RetrieveById(id, &sql.Options{Relations: []string{"variants.prices", "variants.options"}})
	if err != nil {
		return err
	}
	if product == nil {
		return nil
	}
	if err := s.r.ProductRepository().SoftRemove(s.ctx, product); err != nil {
		return err
	}
	// err = s.eventBus_.Emit(ProductServiceEventsDeleted, map[string]interface{}{"id": id})
	// if err != nil {
	// 	return nil, err
	// }
	return nil
}

func (s *ProductService) AddOption(id uuid.UUID, optionTitle string) (*models.Product, *utils.ApplictaionError) {
	product, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	for _, o := range product.Options {
		if o.Title == optionTitle {
			return nil, utils.NewApplictaionError(
				utils.DUPLICATE_ERROR,
				"An option with the title: "+optionTitle+" already exists",
				"500",
				nil,
			)
		}
	}

	option := &models.ProductOption{
		Title: optionTitle,
	}

	if err := s.r.ProductOptionRepository().Save(s.ctx, option); err != nil {
		return nil, err
	}

	for _, variant := range product.Variants {
		optionValue := &models.ProductOptionValue{
			OptionId: uuid.NullUUID{UUID: option.Id},
			Value:    "Default Value",
		}
		variant.Options = append(variant.Options, *optionValue)
	}

	if err := s.r.ProductRepository().Save(s.ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) ReorderVariants(id uuid.UUID, variantOrder uuid.UUIDs) (*models.Product, *utils.ApplictaionError) {
	product, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if len(product.Variants) != len(variantOrder) {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"models.Product variants and new variant order differ in length.",
			nil,
		)
	}

	var variants []models.ProductVariant
	for _, vId := range variantOrder {
		index := slices.IndexFunc(product.Variants, func(v models.ProductVariant) bool {
			return v.Id == vId
		})
		if index == -1 {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"models.Product has no variant with id: "+vId.String(),
				"500",
				nil,
			)
		}
		variants = append(variants, product.Variants[index])
	}

	product.Variants = variants

	if err := s.r.ProductRepository().Save(s.ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateOption(id uuid.UUID, optionId uuid.UUID, data *types.ProductOptionInput) (*models.Product, *utils.ApplictaionError) {
	product, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	for _, o := range product.Options {
		if o.Title == data.Title && o.Id != optionId {
			return nil, utils.NewApplictaionError(
				utils.NOT_FOUND,
				"An option with title "+data.Title+" already exists",
				"500",
				nil,
			)
		}
	}

	var option *models.ProductOption

	query := sql.BuildQuery(models.ProductOption{Model: core.Model{Id: optionId}}, &sql.Options{})

	if err := s.r.ProductOptionRepository().FindOne(s.ctx, option, query); err != nil {
		return nil, err
	}

	if data.Values != nil {
		option.Values = append(option.Values, data.Values...)
	}

	option.Title = data.Title

	if err := s.r.ProductOptionRepository().Save(s.ctx, option); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) RetrieveOptionByTitle(title string, id uuid.UUID) (*models.ProductOption, *utils.ApplictaionError) {
	product, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	for _, o := range product.Options {
		if o.Title == title {
			return &o, nil
		}
	}

	return nil, nil
}

func (s *ProductService) DeleteOption(id uuid.UUID, optionId uuid.UUID) (*models.Product, *utils.ApplictaionError) {
	product, err := s.RetrieveById(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	var option *models.ProductOption

	query := sql.BuildQuery(models.ProductOption{Model: core.Model{Id: optionId}}, &sql.Options{})

	if err := s.r.ProductOptionRepository().FindOne(s.ctx, option, query); err != nil {
		return nil, err
	}

	if option == nil {
		return nil, nil
	}

	if len(product.Variants) > 0 {
		firstVariant := product.Variants[0]
		firstIndex := slices.IndexFunc(firstVariant.Options, func(m models.ProductOptionValue) bool {
			return m.Id == optionId
		})
		if firstIndex == -1 {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"To Delete an option, first Delete all variants, such that when an option is deleted, no duplicate variants will exist.",
				"500",
				nil,
			)
		}

		for _, variant := range product.Variants {
			optionIndex := slices.IndexFunc(variant.Options, func(m models.ProductOptionValue) bool {
				return m.Id == optionId
			})
			if optionIndex == -1 || variant.Options[optionIndex].Value != firstVariant.Options[firstIndex].Value {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"To Delete an option, first Delete all variants, such that when an option is deleted, no duplicate variants will exist.",
					"500",
					nil,
				)
			}
		}
	}

	if err := s.r.ProductOptionRepository().Delete(s.ctx, option); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateShippingProfile(productIds uuid.UUIDs, profileId uuid.UUID) ([]models.Product, *utils.ApplictaionError) {
	var products []models.Product
	for _, id := range productIds {
		product, err := s.RetrieveById(id, &sql.Options{})
		if err != nil {
			return nil, err
		}

		if profileId == uuid.Nil {
			product.Profiles = []models.ShippingProfile{}
		} else {
			product.Profiles = []models.ShippingProfile{{Model: core.Model{Id: profileId}}}
		}

		products = append(products, *product)
	}

	if err := s.r.ProductRepository().SaveSlice(s.ctx, products); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) DecorateProductsWithSalesChannels(products []models.Product) ([]models.Product, *utils.ApplictaionError) {
	productIdSalesChannelMap, err := s.getSalesChannelModuleChannels(products)
	if err != nil {
		return nil, err
	}

	for i := range products {
		product := &products[i]
		product.SalesChannels = productIdSalesChannelMap[product.Id]
	}

	return products, nil
}

func (s *ProductService) getSalesChannelModuleChannels(products []models.Product) (map[uuid.UUID][]models.SalesChannel, *utils.ApplictaionError) {
	var data []models.Product

	var productIds uuid.UUIDs
	for _, product := range products {
		productIds = append(productIds, product.Id)
	}

	query := sql.BuildQuery(models.Product{}, &sql.Options{
		Specification: []sql.Specification{sql.In("id", productIds)},
		Selects:       []string{"sales_channels", "id"},
	})

	if err := s.r.ProductRepository().Find(s.ctx, data, query); err != nil {
		return nil, err
	}

	res := make(map[uuid.UUID][]models.SalesChannel)

	for _, record := range data {
		res[record.Id] = record.SalesChannels
	}

	return res, nil
}
