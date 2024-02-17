package main

import (
	"context"

	"github.com/driver005/gateway/cmd"
)

//Swagger: swag init -g main.go  --parseDependency --parseInternal --parseDepth 1  --output docs/

var ctx = context.Background()

func main() {
	// spec := oas.NewOpenAPI()
	// spec.Parse(".", []string{}, "", false, "admin")
	// d, err := yaml.Marshal(&spec)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// _ = os.WriteFile("./admin.base.yaml", d, 0644)
	cmd.Execute()
}
