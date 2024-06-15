package schema

import (
	"fmt"
	"io"
	"regexp"
	"slices"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Migrator struct {
	DB        *gorm.DB
	Tables    []*Table
	Relations map[string][]Relation
}

type Table struct {
	Name        string
	Fields      []*Field
	SqlFields   int
	ForeignKeys map[string]ForeignKey
	Sql         string
	CreatedAt   time.Time
	migrator    *Migrator
}

type Field struct {
	Name    string
	Type    string
	GormTag string
	JsonTag string
	Sql     string
}

type Relation struct {
	Name      string
	Type      string
	GormTag   string
	IsList    bool
	IsPointer bool
	Sql       string
}

func NewMigrator() *Migrator {
	return &Migrator{
		Relations: make(map[string][]Relation),
	}
}

func (m *Migrator) Schema(path string) error {
	for _, t := range m.Tables {
		var fields strings.Builder

		for _, f := range t.Fields {
			fields.WriteString(fmt.Sprintf("\n    %s %s `gorm:\"%s\" json:\"%s\"`", toUpper(f.Name), f.Type, f.GormTag, f.JsonTag))
		}

		data := strings.Replace(SCHEMASKELLETON, "{.name}", toUpper(t.Name), -1)
		data = strings.Replace(data, "{.tableName}", t.Name, -1)
		data = strings.Replace(data, "{.fields}", fields.String(), -1)

		f, err := NewFile(path, t, nil, "go")
		if err != nil {
			return err
		}

		if err := f.WithCreate(func(w io.Writer) error {
			if _, err := w.Write([]byte(data)); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) Sql(path string) error {
	for _, t := range m.Tables {
		now := time.Now()
		var sql strings.Builder

		up, err := NewFile(path, t, &Version{
			Format:    "unixNano",
			Direction: "up",
			Time:      now,
		}, "sql")
		if err != nil {
			return err
		}

		down, err := NewFile(path, t, &Version{
			Format:    "unixNano",
			Direction: "down",
			Time:      now,
		}, "sql")
		if err != nil {
			return err
		}

		for i, f := range t.Fields {
			if f.Sql != "" {
				if i == t.SqlFields-1 {
					sql.WriteString(f.Sql + "\n")
				} else {
					sql.WriteString(f.Sql + ",\n")
				}
			}
		}

		upSql := strings.Replace(t.Sql, "?", "\n"+sql.String(), -1)
		downSql := fmt.Sprintf("DROP TABLE \"%s\" CASCADE;", t.Name)

		if err := up.WithCreate(func(w io.Writer) error {
			if _, err := w.Write([]byte(upSql)); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}

		if err := down.WithCreate(func(w io.Writer) error {
			if _, err := w.Write([]byte(downSql)); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) AddColumn(column Column, sqlString string) {
	var sql strings.Builder
	var gormTag strings.Builder
	omitempty := ""
	sql.WriteString(fmt.Sprintf(`"%s" %s`, column.Name, column.Type))
	gormTag.WriteString(fmt.Sprintf("column:%s;type:%s;", column.Name, column.Type))

	if column.Unsigned {
		gormTag.WriteString("unsigned;")
		sql.WriteString(" UNSIGNED")
	}
	if column.Autoincrement {
		gormTag.WriteString("autoIncrement;")
		sql.WriteString(" AUTO_INCREMENT")
	}
	if column.Length != 0 {
		gormTag.WriteString(fmt.Sprintf("size:%v;", column.Length))
	}
	if column.Default != "" {
		gormTag.WriteString(fmt.Sprintf("default:%s;", column.Default))
	}
	if column.GormTag != "" {
		gormTag.WriteString(column.GormTag)
	}
	if !column.Nullable {
		gormTag.WriteString("not null;")
		sql.WriteString(" NOT NULL")
	} else {
		gormTag.WriteString("null;")
		omitempty = ",omitempty"
		sql.WriteString(" NULL")
	}

	if column.Primary || column.Name == "id" {
		gormTag.WriteString("primaryKey;")
		sql.WriteString(" PRIMARY KEY")
	}

	if sqlString == "nil" {
		sql.Reset()
	} else if sqlString != "" {
		sql.Reset()
		sql.WriteString(sqlString)
	}

	if sql.Len() != 0 {
		t.SqlFields = t.SqlFields + 1
	}

	t.Fields = append(t.Fields, &Field{
		Name:    column.Name,
		Type:    gormType(column.MappedType),
		JsonTag: fmt.Sprintf(`%s%s`, column.Name, omitempty),
		GormTag: gormTag.String(),
		Sql:     sql.String(),
	})
}

func (t *Table) DropColumn(name string) {
	for i, v := range t.Fields {
		if v.Name == name {
			t.Fields = slices.Delete(t.Fields, i, i+1)
		}
	}
}

func (t *Table) AddIndex(index Index) {
	var tag strings.Builder
	sql := ""
	tag.WriteString(fmt.Sprintf("index:%s,priority:1", index.KeyName))

	if index.Unique {
		tag.WriteString(",unique")
		sql += "UNIQUE "
	}
	if index.Composite {
		tag.WriteString(fmt.Sprintf(",composite:%s", index.KeyName))
	}

	if index.Expression == "" {
		sql = fmt.Sprintf(`CREATE %sINDEX %s ON "%s" ("%s");\n`, sql, index.KeyName, t.Name, strings.Join(index.ColumnNames, `", "`))
	} else {
		sql = index.Expression + ";\n"
	}

	for _, v := range t.Fields {
		for _, i := range index.ColumnNames {
			if v.Name == i {
				v.GormTag += tag.String()
				return
			}
		}
	}

	t.Sql += sql
}

func (t *Table) DropIndex(indexName string) {
	for _, v := range t.Fields {
		if v.Name == t.Name {
			reg := regexp.MustCompile(`index:` + indexName + `,(.*?);`)
			v.GormTag = reg.ReplaceAllString(v.GormTag, "${1}")
		}
	}
}

func (t *Table) AddCheck(check Check) {
	for _, v := range t.Fields {
		if v.Name == check.Name {
			v.GormTag += fmt.Sprintf("check:%s", check.Expression)
			v.Sql += fmt.Sprintf("CONSTRAINT %s CHECK (%s),\n", check.Name, check.Expression)
			return
		}
	}
}

func (t *Table) DropCheck(name string) {
	for _, v := range t.Fields {
		if v.Name == name {
			reg := regexp.MustCompile(`check:(.*?);`)
			v.GormTag = reg.ReplaceAllString(v.GormTag, "${1}")
		}
	}
}

func (t *Table) AddRelation(relation Relation) {
	relationName := toUpper(relation.Type)
	mappedType := "*" + relationName

	if relation.IsList {
		mappedType = strings.Replace(mappedType, "*", "[]", 1)
	}

	column := Column{
		Name:          relation.Name,
		Type:          relationName,
		Unsigned:      false,
		Autoincrement: false,
		Primary:       false,
		Nullable:      true,
		Length:        0,
		Default:       "",
		MappedType:    mappedType,
		GormTag:       relation.GormTag,
	}

	t.AddColumn(column, relation.Sql)
}

func (t *Table) DropRelation(name string) {
	for _, c := range t.Fields {
		if c.Name == name {
			t.DropColumn(name)
		}
	}
}

func (m *Migrator) Tabel(row *Row) *Table {
	for _, t := range m.Tables {
		if t.Name == row.Name {
			return t
		}
	}

	table := &Table{
		Name:      row.Name,
		Sql:       fmt.Sprintf("CREATE TABLE IF NOT EXISTS \"%s\" (?);\n", row.Name),
		CreatedAt: time.Now(),
		Fields:    []*Field{},
		migrator:  m,
	}

	m.Tables = append(m.Tables, table)

	m.AddColumn(table, row.Columns)
	m.AddCheck(table, row.Checks)
	m.AddIndex(table, row.Indexes)
	m.AddForeignKey(row.ForeignKeys)

	return table
}

func (m *Migrator) GetTabel(name string) *Table {
	for _, t := range m.Tables {
		if t.Name == name {
			return t
		}
	}
	return nil
}

func (m *Migrator) AddColumn(t *Table, columns map[string]Column) {
	for _, column := range columns {
		t.AddColumn(column, "")
	}
}

func (m *Migrator) AddIndex(t *Table, indexes []Index) {
	for _, index := range indexes {
		t.AddIndex(index)
	}
}

func (m *Migrator) AddCheck(t *Table, checks []Check) {
	for _, check := range checks {
		t.AddCheck(check)
	}
}

func (m *Migrator) AddForeignKey(keys map[string]ForeignKey) {
	relations := map[string][]Relation{}

	if len(keys) == 2 {
		foreignKeys := make([]Relation, 0, 2)
		for _, fk := range keys {
			localTable := strings.Split(fk.LocalTableName, ".")[1]
			relationName := strings.Split(fk.ReferencedTableName, ".")[1]
			field := Relation{
				Name:      relationName,
				Type:      relationName,
				GormTag:   fmt.Sprintf("many2many:%s;", localTable),
				IsList:    true,
				IsPointer: false,
				Sql:       "nil",
			}

			foreignKeys = append(foreignKeys, field)
		}
		if len(foreignKeys) == 2 {
			relations[foreignKeys[0].Type] = append(relations[foreignKeys[0].Type], foreignKeys[1])
			relations[foreignKeys[1].Type] = append(relations[foreignKeys[1].Type], foreignKeys[0])
		}
	} else {
		// Add foreign key fields and relationships
		for _, fk := range keys {
			unique := false
			referencedTable := strings.Split(fk.LocalTableName, ".")[1]
			localTable := strings.Split(fk.ReferencedTableName, ".")[1]

			field := Relation{
				Name:      referencedTable,
				Type:      referencedTable,
				GormTag:   fmt.Sprintf("foreignKey:%s;", localTable),
				IsList:    !unique,
				IsPointer: unique,
				Sql:       "nil",
			}

			relations[localTable] = append(relations[localTable], field)
		}
	}

	for _, foreignKey := range keys {
		var sql strings.Builder
		var constraint strings.Builder

		if foreignKey.ReferencedTableName != "" {
			sql.WriteString(fmt.Sprintf(`CONSTRAINT "%s" FOREIGN KEY ("%s") REFERENCES "%s" ("%s") `, foreignKey.ConstraintName, foreignKey.ColumnNames[0], strings.Split(foreignKey.ReferencedTableName, ".")[1], strings.Join(foreignKey.ReferencedColumnNames, ", ")))
		}
		if foreignKey.UpdateRule != "" {
			sql.WriteString(fmt.Sprintf("ON UPDATE %s ", foreignKey.UpdateRule))
			constraint.WriteString(fmt.Sprintf("OnUpdate:%s,", foreignKey.UpdateRule))
		}
		if foreignKey.DeleteRule != "" {
			sql.WriteString(fmt.Sprintf("ON DELETE %s", foreignKey.DeleteRule))
			constraint.WriteString(fmt.Sprintf("OnDelete:%s", foreignKey.DeleteRule))
		}

		localTable := strings.Split(foreignKey.LocalTableName, ".")[1]
		referencedTable := strings.Split(foreignKey.ReferencedTableName, ".")[1]

		field := Relation{
			Name:      referencedTable,
			Type:      referencedTable,
			IsList:    false,
			IsPointer: true,
			Sql:       sql.String(),
		}

		if constraint.String() != "" {
			field.GormTag = fmt.Sprintf("constraint:%s;polymorphic:Owner;", constraint.String())
		}

		relations[localTable] = append(relations[localTable], field)
	}

	for table, relation := range relations {
		t := m.GetTabel(table)
		for _, r := range relation {
			t.AddRelation(r)
		}
	}
}
