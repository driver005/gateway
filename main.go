package main

import (
	"context"
	"log"
	"os"

	"github.com/driver005/gateway/oas"
	"gopkg.in/yaml.v2"
)

//Swagger: swag init -g main.go  --parseDependency --parseInternal --parseDepth 1  --output docs/

var ctx = context.Background()

func main() {
	spec := oas.NewOpenAPI()
	spec.Parse(".", []string{}, "", false, "admin")
	d, err := yaml.Marshal(&spec)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	_ = os.WriteFile("./admin.base.yaml", d, 0644)
}
