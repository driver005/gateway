package main

import (
	"github.com/driver005/gateway/cmd"
)

func main() {
	cmd.Execute()
}

// spec := oas.NewOpenAPI()
// spec.Parse(".", []string{}, "", false, "admin")
// d, err := yaml.Marshal(&spec)
// if err != nil {
// 	log.Fatalf("error: %v", err)
// }
// _ = os.WriteFile("./admin.base.yaml", d, 0644)

// g := gen.NewGenerator(gen.Config{
// 	OutPath:           "./query",
// 	FieldNullable:     false,
// 	FieldCoverable:    false,
// 	FieldWithIndexTag: true,
// 	FieldWithTypeTag:  true,
// })

// gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/medusa"))
// g.UseDB(gormdb)

// g.ApplyBasic(
// 	// Generate structs from all tables of current database
// 	g.GenerateAllTable()...,
// )
// // Generate the code
// g.Execute()
