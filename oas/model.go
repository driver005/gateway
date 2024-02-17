package oas

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
)

var (
	regexpPath = regexp.MustCompile(`@oas:path\s\[(.*?)\]\s(.*?)\n`)
	// regexpPath    = regexp.MustCompile("@openapi:Path\n([^@]*)$")
	regexpSchema     = regexp.MustCompile(`@oas:schema:?(\w+)?:?(?:\[([\w,]+)\])?`)
	regexpExample    = regexp.MustCompile(`@oas:example [^\v\n]+`)
	regexpInfo       = regexp.MustCompile("@oas:info\n([^@]*)$")
	regexpImport     = regexp.MustCompile(`import\(([^\)]+)\)`)
	tab              = regexp.MustCompile(`\t`)
	regexStartOfLine = regexp.MustCompile(`(?m)^`)
)

type OpenAPI struct {
	Openapi    string                 `yaml:"openapi"`
	Info       Info                   `yaml:"info"`
	Servers    []Server               `yaml:"servers,omitempty"`
	Paths      map[string]Path        `yaml:"paths,omitempty"`
	Tags       []Tag                  `yaml:"tags,omitempty"`
	Components map[string]interface{} `yaml:"components,omitempty"`
	Security   []map[string][]string  `yaml:"security,omitempty"`
	XGroupTags []interface{}          `yaml:"x-tagGroups,omitempty"`

	registeredSchemas map[string]interface{}
}

type Server struct {
	URL         string                    `yaml:"url,omitempty"`
	Description string                    `yaml:"description,omitempty"`
	Variables   map[string]ServerVariable `yaml:"variables,omitempty"`
}

type ServerVariable struct {
	Default     string   `yaml:"default,omitempty"`
	Enum        []string `yaml:"enum,omitempty"`
	Description string   `yaml:"description,omitempty"`
}

func NewOpenAPI() OpenAPI {
	spec := OpenAPI{}
	spec.Openapi = "3.0.0"
	spec.Paths = make(map[string]Path)
	spec.Components = make(map[string]interface{})
	spec.registeredSchemas = map[string]interface{}{
		"AnyValue": map[string]string{
			"description": "Can be anything: string, number, array, object, etc., including `null`",
		},
	}
	return spec
}

func (spec *OpenAPI) parseYamlFile(Path string) error {
	data, err := parseYamlFile(spec, Path)
	if err != nil {
		return err
	}

	spec = data
	return nil
}

type SecuritySchemes struct {
	Type   string          `yaml:"type,omitempty"`
	Flows  map[string]Flow `yaml:"flows,omitempty"`
	Scheme string          `yaml:"scheme,omitempty"`
}

type Flow struct {
	AuthorizationURL string            `yaml:"authorizationUrl"`
	TokenURL         string            `yaml:"tokenUrl"`
	Scopes           map[string]string `yaml:"scopes"`
}

type Info struct {
	Version     string            `yaml:"version"`
	Title       string            `yaml:"title"`
	Description string            `yaml:"description"`
	XLogo       map[string]string `yaml:"x-logo,omitempty"`
	Contact     map[string]string `yaml:"contact,omitempty"`
	Licence     License           `yaml:"licence,omitempty"`
}

type License struct {
	Name string `yaml:"name,omitempty"`
	Url  string `yaml:"url,omitempty"`
}

type Tag struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
}

func NewEntity() Schema {
	e := Schema{}
	e.Properties = make(map[string]*Schema)
	e.Items = make(map[string]interface{})
	return e
}

type MetaSchema interface {
	RealName() string
	CustomName() string
	SetCustomName(string)
}

type Metadata struct {
	RealName   string `yaml:"-"`
	CustomName string `yaml:"-"`
}

type ComposedSchema struct {
	Metadata `yaml:"-"`
	AllOf    []*Schema `yaml:"allOf"`
}

type ExternalDoc struct {
	Description string `yaml:",omitempty"`
	Url         string `yaml:"url,omitempty"`
}

func (c *ComposedSchema) RealName() string {
	if c == nil {
		return ""
	}
	return c.Metadata.RealName
}

func (c *ComposedSchema) CustomName() string {
	if c == nil {
		return ""
	}
	return c.Metadata.CustomName
}

func (c *ComposedSchema) SetCustomName(customName string) {
	if c == nil {
		return
	}
	c.Metadata.CustomName = customName
}

type Schema struct {
	Metadata             `yaml:"-"`
	Description          string                 `yaml:"description,omitempty"`
	Nullable             *bool                  `yaml:"nullable,omitempty"`
	Required             []string               `yaml:"required,omitempty"`
	Type                 string                 `yaml:"type,omitempty"`
	Items                map[string]interface{} `yaml:"items,omitempty"`
	Format               string                 `yaml:"format,omitempty"`
	Ref                  string                 `yaml:"$ref,omitempty"`
	Enum                 []string               `yaml:"enum,omitempty"`
	Properties           map[string]*Schema     `yaml:"properties,omitempty"`
	AdditionalProperties *Schema                `yaml:"additionalProperties,omitempty"`
	OneOf                []Schema               `yaml:"oneOf,omitempty"`
	Example              interface{}            `yaml:"example,omitempty"`
}

func (s *Schema) RealName() string {
	if s == nil {
		return ""
	}
	return s.Metadata.RealName
}

func (s *Schema) CustomName() string {
	if s == nil {
		return ""
	}
	return s.Metadata.CustomName
}

func (s *Schema) SetCustomName(customName string) {
	if s == nil {
		return
	}
	s.Metadata.CustomName = customName
}

type BuildError struct {
	Err     error
	Content string
	Message string
}

func (e BuildError) Error() string {
	err, _ := logrus.
		WithError(e.Err).
		WithField("content", e.Content).
		WithField("message", e.Message).
		String()
	return err
}

type Path map[string]operation

type operation struct {
	Id            string                `yaml:"operationId,omitempty"`
	Summary       string                `yaml:"summary,omitempty"`
	Description   string                `yaml:"description"`
	Authenticated bool                  `yaml:"x-authenticated,omitempty"`
	Codegen       map[string]string     `yaml:"x-codegen,omitempty"`
	CodeSamples   []CodeSample          `yaml:"x-codeSamples,omitempty"`
	Security      []map[string][]string `yaml:"security,omitempty"`
	Tags          []string              `yaml:"tags,omitempty"`
	Responses     map[string]Response   `yaml:"responses"`
	Parameters    []interface{}         `yaml:"parameters,omitempty"`
	RequestBody   RequestBody           `yaml:"requestBody,omitempty"`
	Headers       map[string]Header     `yaml:"headers,omitempty"`
	Deprecated    bool                  `yaml:"deprecated,omitempty"`
	Servers       []Server              `yaml:"servers,omitempty"`
	ExternalDocs  ExternalDoc           `yaml:"externalDocs,omitempty"`
}

// type parameter struct {
// 	Example     string `yaml:"example,omitempty"`
// 	In          string
// 	Name        string
// 	Schema      Schema `yaml:",omitempty"`
// 	Required    bool
// 	Description string
// }

type RequestBody struct {
	Description string             `yaml:"description,omitempty"`
	Required    bool               `yaml:"required,omitempty"`
	Content     map[string]content `yaml:"content"`
}

type Response struct {
	Content     map[string]content `yaml:"content"`
	Description string             `yaml:"description,omitempty"`
	Headers     map[string]Header  `yaml:"headers,omitempty"`
}

type Header struct {
	Description string `yaml:",omitempty"`
	Schema      Schema `yaml:",omitempty"`
}

type content struct {
	Schema Schema
}

type CodeSample struct {
	Lang   string `yaml:"lang"`
	Label  string `yaml:"label"`
	Source string `yaml:"source,omitempty"`
}

func validatePath(Path string, parseVendors []string) bool {
	// vendoring Path
	if strings.Contains(Path, "vendor") {
		found := false
		for _, vendorPath := range parseVendors {
			if strings.Contains(Path, vendorPath) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// not golang file
	if !strings.HasSuffix(Path, ".go") {
		return false
	}

	// dot file
	if strings.HasPrefix(Path, ".") {
		return false
	}

	return true
}

func (spec *OpenAPI) Parse(Path string, parseVendors []string, vendorsPath string, exitNonZeroOnError bool, apiType string) {
	// fset := token.NewFileSet() // positions are relative to fset
	if apiType != "" {
		if err := spec.parseYamlFile(fmt.Sprintf("./oas/data/%s.oas.base.yaml", apiType)); err != nil {
			os.Exit(1)
		}
	} else if apiType == "combined" {

	}

	walker := func(Path string, f os.FileInfo, err error) error {
		if validatePath(Path, parseVendors) {
			astFile, _ := parseFile(Path)
			infosErrors := spec.parseInfos(astFile)
			schemasErrors := spec.parseSchemas(astFile, apiType)
			pathErrors := spec.parsePaths(astFile, apiType)
			if exitNonZeroOnError &&
				(len(infosErrors) > 0 || len(schemasErrors) > 0 || len(pathErrors) > 0) {
				return errors.New("errors while generating OpenAPI Schema")
			}
		}
		return nil
	}

	err := filepath.Walk(Path, walker)
	if err != nil {
		os.Exit(1)
	}

	err = filepath.Walk(vendorsPath, walker)
	if err != nil {
		os.Exit(1)
	}

	spec.composeSpecSchemas()
}

func checkApiType(name string, apiType string) bool {
	if strings.HasPrefix(strings.ToLower(name), "admin") || strings.HasPrefix(strings.ToLower(name), "store") {
		if !strings.HasPrefix(strings.ToLower(name), apiType) {
			return true
		}
	}

	return false
}

func removeFirstLine(text string) string {
	lines := strings.SplitN(text, "\n", 2) // Split into lines, 2 means split only at the first newline
	if len(lines) > 1 {                    // Check if there's more than one line
		return lines[1] // Return the second line and onwards
	}
	return "" // If there's only one line or no line at all, return an empty string
}

func formatSource(text string, indexName string, start string) string {
	startIndex := strings.Index(text, "source: |") + len("source: | ")
	output := text[startIndex:]
	endIndex := strings.Index(output, "  - lang:")
	lastIndex := strings.LastIndex(output, indexName)
	if endIndex == -1 {
		data := text[:startIndex] + regexStartOfLine.ReplaceAllString(output[:lastIndex-1], "    ") + "\n"
		return start + data
	}

	data := text[:startIndex] + regexStartOfLine.ReplaceAllString(output[:endIndex-1], "    ") + "\n"
	return formatSource(output[endIndex:], indexName, start+data)
}

func (spec *OpenAPI) parsePaths(f *ast.File, apiType string) (errs []error) {
	for _, s := range f.Comments {
		t := s.Text()

		// Test if comments is a Path
		a := regexpPath.FindSubmatch([]byte(t))
		if len(a) == 0 {
			continue
		}

		method := string(a[1])
		url := string(a[2])
		data := operation{}

		if checkApiType(strings.SplitN(url, "/", 3)[1], apiType) {
			continue
		}

		yamlString := removeFirstLine(t)

		// Replacing tab with spaces
		content := tab.ReplaceAllString(yamlString, "  ")
		content = strings.Replace(content, "\n\n", "\n", -1)

		startIndex := strings.Index(content, "x-codeSamples:")
		tagIndex := strings.LastIndex(content, "tags:")
		endIndex := strings.LastIndex(content, "security:")
		indexName := "security:"
		if endIndex == -1 {
			endIndex = tagIndex
			indexName = "tags:"
		}
		if startIndex != -1 && endIndex != -1 {
			content = content[:startIndex] + formatSource(content[startIndex:endIndex+len(indexName)], indexName, "") + content[endIndex:]
		}

		// Unmarshal yaml
		err := yaml.Unmarshal([]byte(content), &data)
		if err != nil {
			logrus.
				WithError(err).
				WithField("content", content).
				Error("Unable to unmarshal Path")
			// fmt.Println(err)
			errs = append(errs, &BuildError{
				Err:     err,
				Content: content,
				Message: "unable to unmarshal Path",
			})
			continue
		}

		Path := make(map[string]operation)
		Path[method] = data
		// Add Path to spec
		// Path already exists in the spec
		if _, ok := spec.Paths[url]; ok {
			// Iterate over verbs
			for currentVerb, currentDesc := range Path {
				if _, operationAlreadyExists := spec.Paths[url][currentVerb]; operationAlreadyExists {
					logrus.
						WithField("url", url).
						WithField("verb", currentVerb).
						Error("Verb for this Path already exists")
					errs = append(errs, &BuildError{
						Err:     errors.New("verb for this Path already exists"),
						Content: fmt.Sprintf("url: %s, verb: %s", url, currentVerb),
					})
					continue
				}
				spec.Paths[url][currentVerb] = currentDesc
			}
		} else {
			spec.Paths[url] = Path
		}

		keys := []string{}
		for k := range Path {
			keys = append(keys, k)
		}

		logrus.
			WithField("url", url).
			WithField("verb", keys).
			Info("Parsing Path")

	}
	return
}

func (spec *OpenAPI) replaceSchemaNameToCustom(s *Schema) {
	if s == nil {
		return
	}

	for _, property := range s.Properties {
		spec.replaceSchemaNameToCustom(property)
	}
	spec.replaceSchemaNameToCustom(s.AdditionalProperties)

	refSplit := strings.Split(s.Ref, "/")
	if len(refSplit) != 4 {
		return
	}
	if replacementSchema, found := spec.registeredSchemas[refSplit[3]]; found {
		meta, ok := replacementSchema.(MetaSchema)
		if !ok {
			return
		}
		if meta.CustomName() != "" {
			refSplit[3] = meta.CustomName()
		}
	}
	s.Ref = strings.Join(refSplit, "/")
}

func (spec *OpenAPI) composeSpecSchemas() {
	for realName, registeredSchema := range spec.registeredSchemas {
		if realName == "AnyValue" {
			spec.Components[realName] = registeredSchema
			continue
		}

		meta, ok := registeredSchema.(MetaSchema)
		if !ok {
			continue
		}

		if composed, ok := registeredSchema.(*ComposedSchema); ok {
			for _, s := range composed.AllOf {
				spec.replaceSchemaNameToCustom(s)
			}
		} else if normal, ok := registeredSchema.(*Schema); ok {
			spec.replaceSchemaNameToCustom(normal)
		}

		name := realName
		if meta.CustomName() != "" {
			name = meta.CustomName()
		}
		spec.Components[name] = registeredSchema
	}
}

// func (spec *OpenAPI) parseMaps(f *ast.File, mp *ast.MapType) (*Schema, []error) {
// 	errors := make([]error, 0)

// 	p, err := parseNamedType(f, mp, nil)
// 	if err != nil {
// 		logrus.WithError(err).Error("Can't parse the type of field in map")
// 		errors = append(errors, BuildError{
// 			Err:     err,
// 			Message: "can't parse the type of field in map",
// 		})
// 	}

// 	return p, errors

// }

// func (spec *OpenAPI) parseStructs(f *ast.File, tpe *ast.StructType) (interface{}, []error) {
// 	errors := make([]error, 0)

// 	var cs *ComposedSchema
// 	e := NewEntity()
// 	e.Type = "object"

// 	for _, fld := range tpe.Fields.List {

// 		example, err := spec.parseExample(fld.Doc.Text(), fld.Type)
// 		if err != nil {
// 			errors = append(errors, err)
// 		}

// 		if len(fld.Names) > 0 && fld.Names[0] != nil && fld.Names[0].IsExported() {
// 			j, _ := parseJSONTag(fld)
// 			if j.ignore {
// 				continue
// 			}
// 			if j.required {
// 				e.Required = append(e.Required, j.name)
// 			}

// 			p, err := parseNamedType(f, fld.Type, nil)
// 			if err != nil {
// 				logrus.WithError(err).WithField("field", fld.Names[0]).Error("Can't parse the type of field in struct")
// 				errors = append(errors, BuildError{
// 					Err:     err,
// 					Content: fld.Names[0].String(),
// 					Message: "can't parse the type of field in struct",
// 				})
// 				continue
// 			}

// 			if example != nil {
// 				p.Example = example
// 			}

// 			if len(j.enum) > 0 {
// 				p.Enum = j.enum
// 			}

// 			if p != nil {
// 				e.Properties[j.name] = p
// 			}

// 		} else {
// 			// composition
// 			if cs == nil {
// 				cs = &ComposedSchema{
// 					AllOf: make([]*Schema, 0),
// 				}
// 			}

// 			p, err := parseNamedType(f, fld.Type, nil)
// 			if err != nil {
// 				logrus.WithError(err).WithField("field", fld.Type).Error("Can't parse the type of composed field in struct")
// 				errors = append(errors, BuildError{
// 					Err:     err,
// 					Message: "can't parse the type of composed field in struct",
// 				})
// 				continue
// 			}

// 			if example != nil {
// 				p.Example = example
// 			}

// 			cs.AllOf = append(cs.AllOf, p)
// 		}
// 	}

// 	if cs == nil {
// 		return &e, errors
// 	} else {
// 		cs.AllOf = append(cs.AllOf, &e)
// 		return cs, errors
// 	}
// }

func (spec *OpenAPI) parseExample(comment string, exampleType ast.Expr) (interface{}, error) {
	exampleLines := regexpExample.FindSubmatch([]byte(comment))
	if len(exampleLines) == 0 {
		return nil, nil
	}

	line := string(exampleLines[0])
	example := line[len("@openapi:example "):]

	return convertExample(example, exampleType)
}

func (spec *OpenAPI) parseSchemas(f *ast.File, apiType string) (errors []error) {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		t := gd.Doc.Text()

		// TODO: Rafacto with parseNamedType
		for _, spc := range gd.Specs {

			// If the node is a Type
			if ts, ok := spc.(*ast.TypeSpec); ok {
				realName := ts.Name.Name
				entityName := realName

				// Looking for openapi entity
				a := regexpSchema.FindSubmatch([]byte(t))
				example, err := spec.parseExample(t, ts.Type)
				if err != nil {
					errors = append(errors, err)
				}

				if len(a) == 0 {
					continue
				}

				if string(a[1]) != "" {
					entityName = string(a[1])
				}

				if checkApiType(entityName, apiType) {
					continue
				}

				yamlString := removeFirstLine(t)
				// Replacing tab with spaces
				content := tab.ReplaceAllString(yamlString, "  ")
				content = strings.Replace(content, "\n\n", "\n", -1)
				// fmt.Println(content)

				data := &Schema{}

				// Unmarshal yaml
				err = yaml.Unmarshal([]byte(content), data)
				if err != nil {
					logrus.
						WithError(err).
						WithField("content", content).
						Error("Unable to unmarshal Path")
					// fmt.Println(err)
					continue
				}

				data.SetCustomName(entityName)

				logrus.
					WithField("name", entityName).
					Info("Parsing Schema")

				if example != nil {
					data.Example = example
				}
				spec.registeredSchemas[realName] = data
			}
		}
	}
	return
}

// func (spec *OpenAPI) parseSchemas(f *ast.File) (errors []error) {
// 	for _, decl := range f.Decls {
// 		gd, ok := decl.(*ast.GenDecl)
// 		if !ok {
// 			continue
// 		}
// 		t := gd.Doc.Text()

// 		// TODO: Rafacto with parseNamedType
// 		for _, spc := range gd.Specs {

// 			// If the node is a Type
// 			if ts, ok := spc.(*ast.TypeSpec); ok {
// 				var entity interface{}
// 				realName := ts.Name.Name
// 				entityName := realName

// 				// Looking for openapi entity
// 				a := regexpSchema.FindSubmatch([]byte(t))
// 				example, err := spec.parseExample(t, ts.Type)
// 				if err != nil {
// 					errors = append(errors, err)
// 				}

// 				if len(a) == 0 {
// 					continue
// 				}

// 				if len(a) == 3 {
// 					if string(a[1]) != "" {
// 						entityName = string(a[1])
// 					}
// 				}

// 				switch n := ts.Type.(type) {
// 				case *ast.MapType:
// 					var errs []error
// 					entity, errs = spec.parseMaps(f, n)
// 					if len(errs) != 0 {
// 						errors = append(errors, errs...)
// 					}

// 					logrus.
// 						WithField("name", entityName).
// 						Info("Parsing Schema")

// 				case *ast.StructType:
// 					var errs []error
// 					entity, errs = spec.parseStructs(f, n)
// 					if len(errs) != 0 {
// 						errors = append(errors, errs...)
// 					}

// 					mtd, ok := entity.(MetaSchema)
// 					if ok {
// 						mtd.SetCustomName(entityName)
// 					}
// 					logrus.
// 						WithField("name", entityName).
// 						Info("Parsing Schema")

// 				case *ast.ArrayType:
// 					e := NewEntity()
// 					p, err := parseNamedType(f, n.Elt, nil)
// 					if err != nil {
// 						logrus.WithError(err).Error("Can't parse the type of field in struct")
// 						errors = append(errors, &BuildError{
// 							Err:     err,
// 							Message: "Can't parse the type of field in struct",
// 						})
// 						continue
// 					}

// 					e.Type = "array"
// 					if p.Ref != "" {
// 						e.Items["$ref"] = p.Ref
// 					} else {
// 						e.Items["type"] = p.Type
// 					}
// 					entity = &e

// 				default:
// 					p, err := parseNamedType(f, ts.Type, nil)
// 					if err != nil {
// 						logrus.WithError(err).Error("can't parse custom type")
// 						errors = append(errors, BuildError{
// 							Err:     err,
// 							Message: "can't parse custom type",
// 						})
// 						continue
// 					}
// 					p.SetCustomName(entityName)

// 					logrus.
// 						WithField("name", entityName).
// 						Info("Parsing Schema")
// 					entity = p
// 				}

// 				if entity != nil {
// 					s, ok := entity.(*Schema)
// 					if ok && example != nil {
// 						s.Example = example
// 					}
// 					spec.registeredSchemas[realName] = entity
// 				}
// 			}
// 		}
// 	}
// 	return
// }

func (spec *OpenAPI) AddOperation(Path, verb string, a operation) {
	if _, ok := spec.Paths[Path]; !ok {
		spec.Paths[Path] = make(map[string]operation)
	}
	spec.Paths[Path][verb] = a
}

func (spec *OpenAPI) parseInfos(f *ast.File) (errors []error) {
	for _, s := range f.Comments {
		t := s.Text()
		// Test if comment is an Info block
		a := regexpInfo.FindSubmatch([]byte(t))
		if len(a) == 0 {
			continue
		}

		// Replacing tab with spaces
		content := tab.ReplaceAllString(string(a[1]), "  ")

		// Unmarshal yaml
		infos := Info{}
		err := yaml.Unmarshal([]byte(content), &infos)
		if err != nil {
			logrus.
				WithError(err).
				WithField("content", content).
				Error("Unable to unmarshal infos")
			errors = append(errors, &BuildError{
				Err:     err,
				Content: content,
				Message: "Unable to unmarshal infos",
			})
			continue
		}

		version := infos.Version
		if spec.Info.Version != "" && spec.Info.Version != version {
			logrus.
				WithField("version", spec.Info.Version).
				WithField("version_scanned", version).
				Warn("Version already exists and is different!")
		} else {
			logrus.
				WithField("field", "version").
				WithField("value", version).
				Info("Parsing Info")
			spec.Info.Version = version
		}

		title := infos.Title
		if spec.Info.Title != "" && spec.Info.Title != title {
			logrus.
				WithField("title", spec.Info.Title).
				WithField("title_scanned", title).
				Warn("Title already exists and is different!")
		} else {
			logrus.
				WithField("field", "title").
				WithField("value", title).
				Info("Parsing Info")
			spec.Info.Title = title
		}

		description := infos.Description
		if spec.Info.Description != "" && spec.Info.Description != description {
			logrus.
				WithField("description", spec.Info.Description).
				WithField("description_scanned", description).
				Warn("Description already exists and is different!")
		} else {
			p, err := parseImportContentPath(description)
			// no need to import a file
			if err != nil {
				logrus.
					WithField("field", "description").
					WithField("value", description).
					Info("Parsing Info")
				spec.Info.Description = description
			} else {
				c, err := os.ReadFile(p)

				if err != nil {
					logrus.
						WithField("File", p).
						WithError(err).
						Error("Could not import file")
					return
				}

				logrus.
					WithField("field", "description").
					WithField("value", "content of file: "+p).
					Info("Parsing Info")
				spec.Info.Description = string(c)
			}
		}

		spec.Info.XLogo = infos.XLogo

	}
	return
}

func parseImportContentPath(str string) (string, error) {
	matches := regexpImport.FindStringSubmatch(str)
	if len(matches) == 2 {
		return matches[1], nil
	}
	return "", errors.New("Not an import")
}
