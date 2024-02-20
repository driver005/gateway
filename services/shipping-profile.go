package services

import (
	"context"
	"reflect"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ShippingProfileService struct {
	ctx context.Context
	r   Registry
}

func NewShippingProfileService(
	r Registry,
) *ShippingProfileService {
	return &ShippingProfileService{
		context.Background(),
		r,
	}
}

func (s *ShippingProfileService) SetContext(context context.Context) *ShippingProfileService {
	s.ctx = context
	return s
}

func (s *ShippingProfileService) List(selector *types.FilterableShippingProfile, config *sql.Options) ([]models.ShippingProfile, *utils.ApplictaionError) {
	var res []models.ShippingProfile

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 50
		config.Order = "created_at DESC"
	}

	query := sql.BuildQuery(selector, config)

	if err := s.r.ShippingProfileRepository().Find(s.ctx, &res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ShippingProfileService) GetMapProfileIdsByProductIds(productIds uuid.UUIDs) (*map[uuid.UUID]uuid.UUID, *utils.ApplictaionError) {
	mappedProfiles := make(map[uuid.UUID]uuid.UUID)
	if len(productIds) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"productIds" must be defined`,
			nil,
		)
	}

	var shippingProfiles []models.ShippingProfile
	query := sql.BuildQuery(models.ShippingProfile{}, &sql.Options{
		Relations: []string{"products"},
		Selects:   []string{"id", "products.id"},
	})
	if err := s.r.ShippingProfileRepository().Specification(sql.In("id", productIds)).Find(s.ctx, &shippingProfiles, query); err != nil {
		return nil, err
	}

	for _, profile := range shippingProfiles {
		for _, product := range profile.Products {
			mappedProfiles[product.Id] = profile.Id
		}
	}
	return &mappedProfiles, nil
}

func (s *ShippingProfileService) Retrieve(profileId uuid.UUID, config *sql.Options) (*models.ShippingProfile, *utils.ApplictaionError) {
	if profileId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"profileId" must be defined`,
			nil,
		)
	}
	var res *models.ShippingProfile = &models.ShippingProfile{}

	query := sql.BuildQuery(models.ShippingProfile{Model: core.Model{Id: profileId}}, config)

	if err := s.r.ShippingProfileRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ShippingProfileService) RetrieveForProducts(productIds uuid.UUIDs) (map[string]models.ShippingProfile, *utils.ApplictaionError) {
	if len(productIds) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"productIds" must be defined`,
			nil,
		)
	}

	productProfilesMap, err := s.r.ShippingProfileRepository().FindByProducts(productIds)
	if err != nil {
		return nil, err
	}
	if len(productProfilesMap) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`No Profile found for products with id: `+strings.Join(productIds.Strings(), ", "),
			nil,
		)
	}
	return productProfilesMap, nil
}

func (s *ShippingProfileService) RetrieveDefault() (*models.ShippingProfile, *utils.ApplictaionError) {
	var res *models.ShippingProfile = &models.ShippingProfile{}

	query := sql.BuildQuery(models.ShippingProfile{Type: models.ShippingProfileTypeDefault}, &sql.Options{})

	if err := s.r.ShippingProfileRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ShippingProfileService) CreateDefault() (*models.ShippingProfile, *utils.ApplictaionError) {
	profile, err := s.RetrieveDefault()
	if err != nil {
		return nil, err
	}
	if profile == nil {
		var profile *models.ShippingProfile = &models.ShippingProfile{}
		profile.Name = "Default Shipping Profile"
		profile.Type = models.ShippingProfileTypeDefault

		if err := s.r.ShippingProfileRepository().Save(s.ctx, profile); err != nil {
			return nil, err
		}

		return profile, nil
	}
	return profile, nil
}

func (s *ShippingProfileService) RetrieveGiftCardDefault() (*models.ShippingProfile, *utils.ApplictaionError) {
	var res *models.ShippingProfile = &models.ShippingProfile{}

	query := sql.BuildQuery(models.ShippingProfile{Type: models.ShippingProfileTypeGiftCard}, &sql.Options{})

	if err := s.r.ShippingProfileRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ShippingProfileService) CreateGiftCardDefault() (*models.ShippingProfile, *utils.ApplictaionError) {
	profile, err := s.RetrieveGiftCardDefault()
	if err != nil {
		return nil, err
	}
	if profile == nil {
		var profile *models.ShippingProfile = &models.ShippingProfile{}
		profile.Name = "Gift Card Profile"
		profile.Type = models.ShippingProfileTypeGiftCard

		if err := s.r.ShippingProfileRepository().Save(s.ctx, profile); err != nil {
			return nil, err
		}

		return profile, nil
	}
	return profile, nil
}

func (s *ShippingProfileService) Create(data *types.CreateShippingProfile) (*models.ShippingProfile, *utils.ApplictaionError) {
	model := &models.ShippingProfile{
		Model: core.Model{
			Metadata: data.Metadata,
		},
		Name: data.Name,
		Type: data.Type,
	}

	if err := s.r.ShippingProfileRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *ShippingProfileService) Update(profileId uuid.UUID, data *types.UpdateShippingProfile) (*models.ShippingProfile, *utils.ApplictaionError) {
	profile, err := s.Retrieve(profileId, &sql.Options{})
	if err != nil {
		return nil, err
	}
	if data.Products != nil {
		_, err = s.AddProducts(profile.Id, data.Products)
		if err != nil {
			return nil, err
		}
	}

	if data.ShippingOptions != nil {
		_, err = s.AddShippingOptions(profile.Id, data.ShippingOptions)
		if err != nil {
			return nil, err
		}
	}

	if data.Metadata != nil {
		profile.Metadata = utils.MergeMaps(profile.Metadata, data.Metadata)
	}

	if !reflect.ValueOf(data.Name).IsZero() {
		profile.Name = data.Name
	}
	if !reflect.ValueOf(data.Type).IsZero() {
		profile.Type = data.Type
	}

	if err := s.r.ShippingProfileRepository().Save(s.ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ShippingProfileService) Delete(profileId uuid.UUID) *utils.ApplictaionError {
	profile, err := s.Retrieve(profileId, &sql.Options{})
	if err != nil {
		return err
	}
	if profile == nil {
		return utils.NewApplictaionError(
			utils.CONFLICT,
			"Profile not existing",
			nil,
		)
	}
	err = s.r.ShippingProfileRepository().SoftRemove(s.ctx, profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShippingProfileService) AddProduct(profileId uuid.UUID, productId uuid.UUID) (*models.ShippingProfile, *utils.ApplictaionError) {
	return s.AddProducts(profileId, uuid.UUIDs{productId})
}

func (s *ShippingProfileService) AddProducts(profileId uuid.UUID, productIds uuid.UUIDs) (*models.ShippingProfile, *utils.ApplictaionError) {
	_, err := s.r.ProductService().SetContext(s.ctx).UpdateShippingProfile(productIds, profileId)
	if err != nil {
		return nil, err
	}
	return s.Retrieve(profileId, &sql.Options{})
}

func (s *ShippingProfileService) RemoveProducts(profileId uuid.UUID, productIds uuid.UUIDs) (*models.ShippingProfile, *utils.ApplictaionError) {
	_, err := s.r.ProductService().SetContext(s.ctx).UpdateShippingProfile(productIds, profileId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ShippingProfileService) AddShippingOption(profileId uuid.UUID, optionId uuid.UUID) (*models.ShippingProfile, *utils.ApplictaionError) {
	return s.AddShippingOptions(profileId, uuid.UUIDs{optionId})
}

func (s *ShippingProfileService) AddShippingOptions(profileId uuid.UUID, optionIds uuid.UUIDs) (*models.ShippingProfile, *utils.ApplictaionError) {
	_, err := s.r.ShippingOptionService().SetContext(s.ctx).UpdateShippingProfile(optionIds, profileId)
	if err != nil {
		return nil, err
	}
	return s.Retrieve(profileId, &sql.Options{
		Relations: []string{"products.profiles", "shipping_options.profile"},
	})

}

func (s *ShippingProfileService) FetchCartOptions(cart *models.Cart) ([]models.ShippingOption, *utils.ApplictaionError) {
	profileIds, err := s.GetProfilesInCart(cart)
	if err != nil {
		return nil, err
	}

	customShippingOptions, err := s.r.CustomShippingOptionService().SetContext(s.ctx).List(models.CustomShippingOption{CartId: uuid.NullUUID{UUID: cart.Id}}, &sql.Options{
		Selects: []string{"id", "shipping_option_id", "price"},
	})
	if err != nil {
		return nil, err
	}
	hasCustomShippingOptions := len(customShippingOptions) > 0
	var selector uuid.UUIDs
	if hasCustomShippingOptions {
		for _, cso := range customShippingOptions {
			selector = append(selector, cso.ShippingOption.Id)
		}
	}
	rawOpts, err := s.r.ShippingOptionService().SetContext(s.ctx).List(&types.FilterableShippingOption{AdminOnly: false}, &sql.Options{
		Relations:     []string{"requirements", "profile"},
		Specification: []sql.Specification{sql.In("profile_id", profileIds), sql.In("id", selector)},
	})
	if err != nil {
		return nil, err
	}

	var shippingOptions []models.ShippingOption
	if hasCustomShippingOptions {
		customShippingOptionsMap := make(map[uuid.UUID]models.CustomShippingOption)

		for _, cso := range customShippingOptions {
			customShippingOptionsMap[cso.ShippingOptionId.UUID] = cso
		}

		for _, raw := range rawOpts {
			customOption, ok := customShippingOptionsMap[raw.Id]
			if ok {
				raw.Amount = customOption.Price
			}
			shippingOptions = append(shippingOptions, raw)
		}

		return shippingOptions, nil
	}
	for _, raw := range rawOpts {
		_, err := s.r.ShippingOptionService().SetContext(s.ctx).ValidateCartOption(&raw, cart)
		if err == nil {
			shippingOptions = append(shippingOptions, raw)
		}
	}

	return shippingOptions, nil
}

func (s *ShippingProfileService) GetProfilesInCart(cart *models.Cart) (uuid.UUIDs, *utils.ApplictaionError) {
	var profileIds uuid.UUIDs

	feature := true

	if feature {
		var productIds uuid.UUIDs

		for _, item := range cart.Items {
			if item.Variant != nil {
				productIds = append(productIds, item.Variant.ProductId.UUID)
			}
		}

		productShippingProfileMap, err := s.GetMapProfileIdsByProductIds(productIds)
		if err != nil {
			return nil, err
		}
		for _, profileId := range *productShippingProfileMap {
			profileIds = append(profileIds, profileId)
		}
	} else {
		for _, item := range cart.Items {
			if item.Variant.Product != nil {
				profileIds = append(profileIds, item.Variant.Product.ProfileId.UUID)
			}
		}
	}
	return profileIds, nil
}
