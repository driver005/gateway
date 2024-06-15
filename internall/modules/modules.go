package modules

import "gorm.io/gorm"

type Module struct {
	Name           string
	DB             *gorm.DB
	DefinitionFile string
	Dir            string
}

func NewModule(
	name string,
	db *gorm.DB,
	definitionFile string,
	dir string,
) *Module {
	return &Module{
		Name:           name,
		DB:             db,
		DefinitionFile: definitionFile,
		Dir:            dir,
	}
}
