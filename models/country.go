package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

//
// @oas:schema:Country
// title: "Country"
// description: "Country details"
// type: object
// required:
//   - display_name
//   - id
//   - iso_2
//   - iso_3
//   - name
//   - num_code
//   - region_id
// properties:
//   id:
//     description: The country's ID
//     type: string
//     example: 109
//   iso_2:
//     description: The 2 character ISO code of the country in lower case
//     type: string
//     example: it
//     externalDocs:
//       url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
//       description: See a list of codes.
//   iso_3:
//     description: The 2 character ISO code of the country in lower case
//     type: string
//     example: ita
//     externalDocs:
//       url: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3#Officially_assigned_code_elements
//       description: See a list of codes.
//   num_code:
//     description: The numerical ISO code for the country.
//     type: string
//     example: 380
//     externalDocs:
//       url: https://en.wikipedia.org/wiki/ISO_3166-1_numeric#Officially_assigned_code_elements
//       description: See a list of codes.
//   name:
//     description: The normalized country name in upper case.
//     type: string
//     example: ITALY
//   display_name:
//     description: The country name appropriate for display.
//     type: string
//     example: Italy
//   region_id:
//     description: The region ID this country is associated with.
//     nullable: true
//     type: string
//     example: reg_01G1G5V26T9H8Y0M4JNE3YGA4G
//   region:
//     description: The details of the region the country is associated with.
//     x-expandable: "region"
//     nullable: true
//     $ref: "#/components/schemas/Region"
//

type Country struct {
	core.SoftDeletableModel

	Iso2        string        `json:"iso_2" gorm:"column:iso_2;uniqueIndex"`
	Iso3        string        `json:"iso_3" gorm:"column:iso_3"`
	NumCode     string        `json:"num_code" gorm:"column:num_code"`
	Name        string        `json:"name" gorm:"column:name"`
	DisplayName string        `json:"display_name" gorm:"column:display_name"`
	RegionId    uuid.NullUUID `json:"region_id" gorm:"column:region_id"`
	Region      *Region       `json:"region" gorm:"foreignKey:RegionId"`
}
