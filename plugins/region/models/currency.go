package models

type Currency struct {
	Code         string `gorm:"column:code;type:text;primaryKey" json:"code"`
	Symbol       string `gorm:"column:symbol;type:text;not null" json:"symbol"`
	SymbolNative string `gorm:"column:symbol_native;type:text;not null" json:"symbol_native"`
	Name         string `gorm:"column:name;type:text;not null" json:"name"`
}

func (*Currency) TableName() string {
	return "region_currency"
}
