package models

// Currency
type Currency struct {
	// core.Model

	// The 3 character ISO code for the currency.
	Code string `json:"code" gorm:"primarykey"`

	// The symbol used to indicate the currency.
	Symbol string `json:"symbol"`

	// The native symbol used to indicate the currency.
	SymbolNative string `json:"symbol_native"`

	// The written name of the currency
	Name string `json:"name"`

	// [EXPERIMENTAL] Does the currency prices include tax
	IncludesTax bool `json:"includes_tax" gorm:"default:null"`
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
