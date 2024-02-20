package models

//
// @oas:schema:Currency
// title: "Currency"
// description: "Currency"
// type: object
// required:
//   - code
//   - name
//   - symbol
//   - symbol_native
// properties:
//  code:
//    description: The 3 character ISO code for the currency.
//    type: string
//    example: usd
//    externalDocs:
//      url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//      description: See a list of codes.
//  symbol:
//    description: The symbol used to indicate the currency.
//    type: string
//    example: $
//  symbol_native:
//    description: The native symbol used to indicate the currency.
//    type: string
//    example: $
//  name:
//    description: The written name of the currency
//    type: string
//    example: US Dollar
//  includes_tax:
//    description: "Whether the currency prices include tax"
//    type: boolean
//    x-featureFlag: "tax_inclusive_pricing"
//    default: false
//

type Currency struct {
	Code         string `json:"code"  gorm:"column:code;primarykey"`
	Symbol       string `json:"symbol"  gorm:"column:symbol"`
	SymbolNative string `json:"symbol_native"  gorm:"column:symbol_native"`
	Name         string `json:"name"  gorm:"column:name"`
	IncludesTax  bool   `json:"includes_tax"  gorm:"column:includes_tax"`
}

// func (c *Currency) BeforeCreate(tx *gorm.DB) (err error) {
// 	c.Id, err = uuid.NewUUID()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c Currency) Validate() error {
// 	return validation.ValidateStruct(&c,
// 		validation.Field(&c.Id, validation.Required, is.UUID),
// 		validation.Field(&c.Code, validation.Match(regexp.MustCompile("^[a-z]{3}$"))),
// 		validation.Field(&c.IncludesTax),
// 		validation.Field(&c.Name),
// 		validation.Field(&c.Symbol),
// 		validation.Field(&c.SymbolNative),
// 	)
// }
