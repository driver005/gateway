package interfaces

type CsvParserContext struct {
	Line       []string
	LineNumber int
	Column     string
	TupleKey   string
}

type LineContext struct {
	LineNumber int
	Line       interface{}
}

type ICsvValidator interface {
	Validate(value interface{}, context CsvParserContext) (bool, error)
}
