package repository

import (
	"fmt"
	"strings"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShippingProfileRepo struct {
	sql.Repository[models.ShippingProfile]
}

func ShippingProfileRepository(db *gorm.DB) *ShippingProfileRepo {
	return &ShippingProfileRepo{*sql.NewRepository[models.ShippingProfile](db)}
}

func (r *ShippingProfileRepo) FindByProducts(productIds uuid.UUIDs) (map[string]models.ShippingProfile, *utils.ApplictaionError) {
	query := `
		SELECT *
		FROM shipping_profiles sp
		INNER JOIN product_shipping_profile psp ON psp.profile_id = sp.id
		WHERE psp.product_id IN (%v)
	`

	// create a comma-separated list of placeholders for the product IDs
	placeholders := make([]string, len(productIds))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	// build the query with the placeholders
	query = fmt.Sprintf(query, strings.Join(placeholders, ","))

	// execute the query
	rows, err := r.Db().Raw(query, productIds.Strings()).Rows()
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	defer rows.Close()

	// create a map to store the shipping profiles
	var shippingProfiles map[string]models.ShippingProfile

	// iterate over the rows
	for rows.Next() {
		var profile models.ShippingProfile

		// scan the row into the ShippingProfile struct
		err := r.Db().ScanRows(rows, profile)
		if err != nil {
			return nil, r.HandleDBError(err)
		}

		// get the product ID from the row
		var productId string
		err = rows.Scan(&productId)
		if err != nil {
			return nil, r.HandleDBError(err)
		}

		// append the shipping profile to the corresponding product ID
		shippingProfiles[productId] = profile
	}

	// check for any errors during iteration
	err = rows.Err()
	if err != nil {
		return nil, r.HandleDBError(err)
	}

	return shippingProfiles, nil
}
