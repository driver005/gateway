package services

import (
	"io"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
)

func TestNewCsvParserService(t *testing.T) {
	type args struct {
		schema    CsvSchema
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want *CsvParserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCsvParserService(tt.args.schema, tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCsvParserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_Parse(t *testing.T) {
	type args struct {
		readableStream io.Reader
	}
	tests := []struct {
		name    string
		p       *CsvParserService
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Parse(tt.args.readableStream)
			if (err != nil) != tt.wantErr {
				t.Errorf("CsvParserService.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_ParserFunc(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name    string
		p       *CsvParserService
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.ParserFunc(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("CsvParserService.ParserFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.ParserFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_buildColumnMap(t *testing.T) {
	type args struct {
		columns []CsvColumn
	}
	tests := []struct {
		name string
		p    *CsvParserService
		args args
		want map[string]CsvColumn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.buildColumnMap(tt.args.columns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.buildColumnMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_resolveColumn(t *testing.T) {
	type args struct {
		tupleKey  string
		columnMap map[string]CsvColumn
	}
	tests := []struct {
		name string
		p    *CsvParserService
		args args
		want *CsvColumn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.resolveColumn(tt.args.tupleKey, tt.args.columnMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.resolveColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_resolveTuple(t *testing.T) {
	type args struct {
		tuple   map[string]interface{}
		column  CsvColumn
		context interfaces.CsvParserContext
	}
	tests := []struct {
		name string
		p    *CsvParserService
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.resolveTuple(tt.args.tuple, tt.args.column, tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.resolveTuple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_getRequiredColumns(t *testing.T) {
	type args struct {
		columns []CsvColumn
	}
	tests := []struct {
		name string
		p    *CsvParserService
		args args
		want []CsvColumn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.getRequiredColumns(tt.args.columns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.getRequiredColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvParserService_getMissingColumns(t *testing.T) {
	type args struct {
		requiredColumnsMap map[string]CsvColumn
		processedColumns   map[string]bool
	}
	tests := []struct {
		name string
		p    *CsvParserService
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.getMissingColumns(tt.args.requiredColumnsMap, tt.args.processedColumns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvParserService.getMissingColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatMissingColumns(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatMissingColumns(tt.args.list); got != tt.want {
				t.Errorf("formatMissingColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCsvReader(t *testing.T) {
	type args struct {
		reader    io.Reader
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want *CsvReader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCsvReader(tt.args.reader, tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCsvReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsvReader_Read(t *testing.T) {
	tests := []struct {
		name    string
		r       *CsvReader
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("CsvReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
