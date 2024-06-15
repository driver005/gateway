package modules

import (
	"path/filepath"
	"strings"

	"github.com/driver005/gateway/internall/schema"
)

func (m *Module) Migrate() {
	filepath.Join(m.Dir, strings.ToLower(m.Name), "migrations", m.DefinitionFile)
	def, err := schema.NewDefinition(filepath.Join(m.Dir, strings.ToLower(m.Name), "migrations", m.DefinitionFile))
	if err != nil {
		panic(err)
	}

	if err := def.Generate(m.DB, filepath.Join(m.Dir, strings.ToLower(m.Name), "models"), filepath.Join(m.Dir, strings.ToLower(m.Name), "migrations")); err != nil {
		panic(err)
	}
}
