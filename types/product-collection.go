package types

import "github.com/driver005/gateway/core"

// @oas:schema:AdminPostCollectionsReq
// type: object
// description: The product collection's details.
// required:
//   - title
//
// properties:
//
//	title:
//	  type: string
//	  description: The title of the collection.
//	handle:
//	  type: string
//	  description: An optional handle to be used in slugs. If none is provided, the kebab-case version of the title will be used.
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateProductCollection struct {
	Title    string     `json:"title"`
	Handle   string     `json:"handle,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostCollectionsCollectionReq
// type: object
// description: The product collection's details to update.
// properties:
//
//	title:
//	  type: string
//	  description: The title of the collection.
//	handle:
//	  type: string
//	  description: An optional handle to be used in slugs. If none is provided, the kebab-case version of the title will be used.
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateProductCollection struct {
	Title    string     `json:"title,omitempty" validate:"omitempty"`
	Handle   string     `json:"handle,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}
