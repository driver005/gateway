package utils

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
)

func ToAddress(data *types.AddressPayload) *models.Address {
	return &models.Address{
		Model: core.Model{
			Metadata: data.Metadata,
		},
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Phone:       data.Phone,
		Company:     data.Company,
		Address1:    data.Address1,
		Address2:    data.Address2,
		City:        data.City,
		CountryCode: data.CountryCode,
		Province:    data.Province,
		PostalCode:  data.PostalCode,
	}
}
