package services

import (
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/driver005/gateway/interfaces"
)

type CsvParserService struct {
	Schema    CsvSchema
	Delimiter string
}

type CsvSchema struct {
	Columns []CsvColumn
}

type CsvColumn struct {
	Name      string
	Required  bool
	Match     *regexp.Regexp
	MapTo     string
	Reducer   func(data interface{}, key string, value string) (interface{}, error)
	Transform func(value string) interface{}
	Validator interfaces.ICsvValidator
}

func NewCsvParserService(schema CsvSchema, delimiter string) *CsvParserService {
	return &CsvParserService{
		Schema:    schema,
		Delimiter: delimiter,
	}
}

func (p *CsvParserService) Parse(readableStream io.Reader) ([]interface{}, error) {
	parsedContent := []interface{}{}
	csvReader := NewCsvReader(readableStream, p.Delimiter)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		parsedLine, err := p.ParserFunc(line)
		if err != nil {
			return nil, err
		}
		parsedContent = append(parsedContent, parsedLine)
	}
	return parsedContent, nil
}

func (p *CsvParserService) ParserFunc(line []string) (interface{}, error) {
	outputTuple := make(map[string]interface{})
	columnMap := p.buildColumnMap(p.Schema.Columns)
	requiredColumnsMap := p.buildColumnMap(p.getRequiredColumns(p.Schema.Columns))
	processedColumns := make(map[string]bool)
	for i, tupleKey := range line {
		column := p.resolveColumn(tupleKey, columnMap)
		if column == nil {
			return nil, errors.New("Unable to treat column " + tupleKey + " from the csv file. No target column found in the provided schema")
		}
		processedColumns[column.Name] = true
		if tupleKey == "" && column.Required {
			return nil, errors.New("No value found for target column \"" + column.Name + "\" in line " + strconv.Itoa(i+1) + " of the given csv file")
		}
		context := interfaces.CsvParserContext{
			Line:       line,
			LineNumber: i + 1,
			Column:     column.Name,
			TupleKey:   tupleKey,
		}
		outputTuple = p.resolveTuple(outputTuple, *column, context)
	}
	missingColumns := p.getMissingColumns(requiredColumnsMap, processedColumns)
	if len(missingColumns) > 0 {
		return nil, errors.New("Missing column(s) " + formatMissingColumns(missingColumns) + " from the given csv file")
	}
	for i, column := range p.Schema.Columns {
		context := interfaces.CsvParserContext{
			Line:       line,
			LineNumber: i + 1,
			Column:     column.Name,
		}
		if column.Validator != nil {
			_, err := column.Validator.Validate(outputTuple, context)
			if err != nil {
				return nil, err
			}
		}
	}
	return outputTuple, nil
}

func (p *CsvParserService) buildColumnMap(columns []CsvColumn) map[string]CsvColumn {
	columnMap := make(map[string]CsvColumn)
	for _, column := range columns {
		if column.Name != "" {
			columnMap[column.Name] = column
		}
	}
	return columnMap
}

func (p *CsvParserService) resolveColumn(tupleKey string, columnMap map[string]CsvColumn) *CsvColumn {
	column, ok := columnMap[tupleKey]
	if ok && column.Match == nil {
		return &column
	}
	for _, c := range p.Schema.Columns {
		if c.Match != nil && c.Match.MatchString(tupleKey) {
			return &c
		}
	}
	return nil
}

func (p *CsvParserService) resolveTuple(tuple map[string]interface{}, column CsvColumn, context interfaces.CsvParserContext) map[string]interface{} {
	outputTuple := make(map[string]interface{})
	for k, v := range tuple {
		outputTuple[k] = v
	}
	resolvedKey := context.TupleKey
	if column.Match == nil && column.MapTo != "" {
		resolvedKey = column.MapTo
	}
	resolvedValue := context.TupleKey
	if column.Transform != nil {
		resolvedValue = column.Transform(context.TupleKey).(string)
	}
	outputTuple[resolvedKey] = resolvedValue
	return outputTuple
}

func (p *CsvParserService) getRequiredColumns(columns []CsvColumn) []CsvColumn {
	requiredColumns := []CsvColumn{}
	for _, column := range columns {
		if column.Required {
			requiredColumns = append(requiredColumns, column)
		}
	}
	return requiredColumns
}

func (p *CsvParserService) getMissingColumns(requiredColumnsMap map[string]CsvColumn, processedColumns map[string]bool) []string {
	missingColumns := []string{}
	for name, column := range requiredColumnsMap {
		if !processedColumns[name] {
			missingColumns = append(missingColumns, column.Name)
		}
	}
	return missingColumns
}

func formatMissingColumns(list []string) string {
	return strings.Join(list, ", ")
}

type CsvReader struct {
	Reader    io.Reader
	Delimiter string
}

func NewCsvReader(reader io.Reader, delimiter string) *CsvReader {
	return &CsvReader{
		Reader:    reader,
		Delimiter: delimiter,
	}
}

func (r *CsvReader) Read() ([]string, error) {
	buf := make([]byte, 1)
	line := []string{}
	field := ""
	quoted := false
	for {
		_, err := r.Reader.Read(buf)
		if err == io.EOF {
			if field != "" {
				line = append(line, field)
			}
			break
		}
		if err != nil {
			return nil, err
		}
		char := string(buf[0])
		if char == r.Delimiter && !quoted {
			line = append(line, field)
			field = ""
		} else if char == "\"" {
			quoted = !quoted
		} else {
			field += char
		}
	}
	return line, nil
}
