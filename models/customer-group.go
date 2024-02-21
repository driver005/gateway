package models

import "github.com/driver005/gateway/core"

//
// @oas:schema:CustomerGroup
// title: "Customer Group"
// description: "A customer group that can be used to organize customers into groups of similar traits."
// type: object
// required:
//   - created_at
//   - deleted_at
//   - id
//   - metadata
//   - name
//   - updated_at
// properties:
//   id:
//     description: The customer group's ID
//     type: string
//     example: cgrp_01G8ZH853Y6TFXWPG5EYE81X63
//   name:
//     description: The name of the customer group
//     type: string
//     example: VIP
//   customers:
//     description: The details of the customers that belong to the customer group.
//     type: array
//     x-expandable: "customers"
//     items:
//       $ref: "#/components/schemas/Customer"
//   price_lists:
//     description: The price lists that are associated with the customer group.
//     type: array
//     x-expandable: "price_lists"
//     items:
//       $ref: "#/components/schemas/PriceList"
//   created_at:
//     description: The date with timezone at which the resource was created.
//     type: string
//     format: date-time
//   updated_at:
//     description: The date with timezone at which the resource was updated.
//     type: string
//     format: date-time
//   deleted_at:
//     description: The date with timezone at which the resource was deleted.
//     nullable: true
//     type: string
//     format: date-time
//   metadata:
//     description: An optional key-value map with additional details
//     nullable: true
//     type: object
//     example: {car: "white"}
//     externalDocs:
//       description: "Learn about the metadata attribute, and how to delete and update it."
//       url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//

type CustomerGroup struct {
	core.SoftDeletableModel

	Name       string      `json:"name" gorm:"column:name"`
	Customers  []Customer  `json:"customers" gorm:"many2many:customer_group_customers"`
	PriceLists []PriceList `json:"price_lists" gorm:"many2many:price_list_customer_groups"`
}
