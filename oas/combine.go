package oas

import (
	"fmt"
	"strings"
)

func CombineOAS(adminOAS, storeOAS OpenAPI) OpenAPI {
	prepareOASForCombine(&adminOAS, "admin")
	prepareOASForCombine(&storeOAS, "store")
	combinedOAS := OpenAPI{
		Openapi: "3.0.0",
		Info: Info{
			Title:   "Medusa API",
			Version: "1.0.0",
		},
		Servers:    []Server{},
		Paths:      make(map[string]Path),
		Tags:       []Tag{},
		Components: map[string]interface{}{},
	}

	for _, oas := range []OpenAPI{adminOAS, storeOAS} {
		for key, value := range oas.Paths {
			combinedOAS.Paths[key] = value
		}
		if oas.Tags != nil {
			combinedOAS.Tags = append(combinedOAS.Tags, oas.Tags...)
		}
		if oas.Components != nil {
			for key, value := range oas.Components {
				combinedOAS.Components[key] = value
			}
		}
		// Repeat the above if block for other components like Examples, Headers, etc.
	}

	return combinedOAS
}

func prepareOASForCombine(oas *OpenAPI, apiType string) *OpenAPI {
	fmt.Printf("🔵 Prefixing %s tags and operationId with %s\n", apiType, strings.ToUpper(apiType))
	for _, pathItem := range oas.Paths {
		for _, operation := range pathItem {
			if operation.Tags != nil {
				for i, tag := range operation.Tags {
					operation.Tags[i] = getPrefixedTagName(tag, apiType)
				}
			}
			if operation.Id != "" {
				operation.Id = getPrefixedOperationId(operation.Id, apiType)
			}
		}
	}
	if oas.Tags != nil {
		for _, tag := range oas.Tags {
			tag.Name = getPrefixedTagName(tag.Name, apiType)
		}
	}
	return oas
}

func getPrefixedTagName(tagName, apiType string) string {
	return fmt.Sprintf("%s %s", strings.ToUpper(apiType), tagName)
}

func getPrefixedOperationId(operationId, apiType string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(apiType), operationId)
}
