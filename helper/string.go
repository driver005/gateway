package helper

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/lib/pq"
)

type (
	RegisteredCases struct {
		cases  pq.StringArray
		actual string
	}
	errUnknownCase struct {
		*RegisteredCases
	}
	RegisteredPrefixes struct {
		prefixes pq.StringArray
		actual   string
	}
	errUnknownPrefix struct {
		*RegisteredPrefixes
	}
)

var (
	ErrUnknownCase   = errUnknownCase{}
	ErrUnknownPrefix = errUnknownPrefix{}
)

// Case

// ToLowerInitial converts a string's first character to lower case.
func ToLowerInitial(s string) string {
	if s == "" {
		return ""
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// ToUpperInitial converts a string's first character to upper case.
func ToUpperInitial(s string) string {
	if s == "" {
		return ""
	}
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

// Coalesce

// Coalesce returns the first non-empty string value
func Coalesce(str ...string) string {
	for _, s := range str {
		if s != "" {
			return s
		}
	}
	return ""
}

// Default

func DefaultIfEmpty(s string, defaultValue string) string {
	if len(s) == 0 {
		return defaultValue
	}
	return s
}

// Split

// Split is a special case of strings.Split
// which returns an empty slice if the string is empty
func Split(s, sep string) pq.StringArray {
	if s == "" {
		return pq.StringArray{}
	}

	return strings.Split(s, sep)
}

// SwitchCase

func SwitchExact(actual string) *RegisteredCases {
	return &RegisteredCases{
		actual: actual,
	}
}

func SwitchPrefix(actual string) *RegisteredPrefixes {
	return &RegisteredPrefixes{
		actual: actual,
	}
}

func (r *RegisteredCases) AddCase(c string) bool {
	r.cases = append(r.cases, c)
	return r.actual == c
}

func (r *RegisteredPrefixes) HasPrefix(prefix string) bool {
	r.prefixes = append(r.prefixes, prefix)
	return strings.HasPrefix(r.actual, prefix)
}

func (r *RegisteredCases) String() string {
	return "[" + strings.Join(r.cases, ", ") + "]"
}

func (r *RegisteredPrefixes) String() string {
	return "[" + strings.Join(r.prefixes, ", ") + "]"
}

func (r *RegisteredCases) ToUnknownCaseErr() error {
	return errUnknownCase{r}
}

func (r *RegisteredPrefixes) ToUnknownPrefixErr() error {
	return errUnknownPrefix{r}
}

func (e errUnknownCase) Error() string {
	return fmt.Sprintf("expected one of %s but got %s", e.String(), e.actual)
}

func (e errUnknownCase) Is(err error) bool {
	_, ok := err.(errUnknownCase)
	return ok
}

func (e errUnknownPrefix) Error() string {
	return fmt.Sprintf("expected %s to have one of the prefixes %s", e.actual, e.String())
}

func (e errUnknownPrefix) Is(err error) bool {
	_, ok := err.(errUnknownPrefix)
	return ok
}
