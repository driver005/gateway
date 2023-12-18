package interfaces

import (
	"io"
)

// IParser is the generic parsing interface. All different parsing implementations (csv, json, etc.) should implement this interface
type IParser interface {
	// Parse parses the readable stream with the given options
	Parse(readableStream io.Reader, options interface{}) ([]interface{}, error)
}
