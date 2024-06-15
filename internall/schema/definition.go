package schema

import (
	"fmt"

	"gorm.io/gorm"
)

type Structur struct {
}

type Definition struct {
	Namespaces []string `json:"namespaces"`
	Name       string   `json:"name"`
	Tables     []Row    `json:"tables"`
}

func NewDefinition(filename string) (*Definition, error) {
	var data Definition

	f, err := LoadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := f.Unmarshal(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *Definition) Generate(db *gorm.DB, modelDir string, migrationDir string) error {
	migrator := NewMigrator()
	for _, table := range d.Tables {
		if err := table.Migration(migrator); err != nil {
			return err
		}
	}
	if err := migrator.Schema(modelDir); err != nil {
		return err
	}

	if err := migrator.Sql(migrationDir); err != nil {
		return err
	}
	return nil
}

type Row struct {
	Columns     map[string]Column     `json:"columns,omitempty"`
	Name        string                `json:"name"`
	Schema      string                `json:"schema"`
	Indexes     []Index               `json:"indexes,omitempty"`
	Checks      []Check               `json:"checks,omitempty"`
	ForeignKeys map[string]ForeignKey `json:"foreignKeys,omitempty"`
}

func (t *Row) Migration(migrator *Migrator) error {
	t.Up(migrator)

	// migrator.Schema()
	if err := t.Down(migrator); err != nil {
		return err
	}

	return nil
}

func (r *Row) Up(migrator *Migrator) {
	migrator.Tabel(r)

	// _, err := NewFile("./test", table, "up")
	// if err != nil {
	// 	return nil, err
	// }

	// if err := f.WithCreate(func(w io.Writer) error {
	// 	if _, err = w.Write([]byte(m.Fields)); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }); err != nil {
	// 	return err
	// }
}

func (r *Row) Down(migrator *Migrator) error {
	table := migrator.Tabel(r)

	_, err := NewFile("./test", table, &Version{
		Format:    "unixNano",
		Direction: "down",
	}, "sql")
	if err != nil {
		return err
	}

	_ = fmt.Sprintf("DROP TABLE %s CASCADE", r.Name)

	// if err := f.WithCreate(func(w io.Writer) error {
	// 	if _, err = w.Write([]byte(sql)); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }); err != nil {
	// 	return err
	// }

	return nil
}

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Unsigned      bool   `json:"unsigned"`
	Autoincrement bool   `json:"autoincrement"`
	Primary       bool   `json:"primary"`
	Nullable      bool   `json:"nullable"`
	Length        int    `json:"length,omitempty"`
	Default       string `json:"default,omitempty"`
	MappedType    string `json:"mappedType"`
	GormTag       string `json:"-"`
}

type Index struct {
	KeyName     string   `json:"keyName"`
	ColumnNames []string `json:"columnNames"`
	Composite   bool     `json:"composite"`
	Primary     bool     `json:"primary"`
	Unique      bool     `json:"unique"`
	Expression  string   `json:"expression,omitempty"`
}

type Check struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
	Definition string `json:"definition"`
}

type ForeignKey struct {
	ConstraintName        string   `json:"constraintName"`
	ColumnNames           []string `json:"columnNames"`
	LocalTableName        string   `json:"localTableName"`
	ReferencedColumnNames []string `json:"referencedColumnNames"`
	ReferencedTableName   string   `json:"referencedTableName"`
	DeleteRule            string   `json:"deleteRule"`
	UpdateRule            string   `json:"updateRule"`
}
