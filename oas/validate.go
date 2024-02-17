package oas

import "github.com/getkin/kin-openapi/openapi3"

func ValidateFile(path string) error {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return err
	}

	if err = doc.Validate(loader.Context); err != nil {
		return err
	}

	return nil
}
