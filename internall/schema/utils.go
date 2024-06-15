package schema

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	TOUPPER = map[string]string{
		"_a": "A",
		"_b": "B",
		"_c": "C",
		"_d": "D",
		"_e": "E",
		"_f": "F",
		"_g": "G",
		"_h": "H",
		"_i": "I",
		"_j": "J",
		"_k": "K",
		"_l": "L",
		"_m": "M",
		"_n": "N",
		"_o": "O",
		"_p": "P",
		"_q": "Q",
		"_r": "R",
		"_s": "S",
		"_t": "T",
		"_u": "U",
		"_v": "V",
		"_w": "W",
		"_x": "X",
		"_y": "Y",
		"_z": "Z",
		"_0": "0",
		"_1": "1",
		"_2": "2",
		"_3": "3",
		"_4": "4",
		"_5": "5",
		"_6": "6",
		"_7": "7",
		"_8": "8",
		"_9": "9",
	}
)

func toUpper(data string) string {
	upper := cases.Title(language.English, cases.Compact).String(data)
	for i, v := range TOUPPER {
		upper = strings.ReplaceAll(upper, i, v)
	}

	return upper
}

// Map JSON types to Go/GORM types
func gormType(jsonType string) string {
	switch jsonType {
	case "text":
		return "string"
	case "boolean":
		return "bool"
	case "integer":
		return "int"
	case "decimal":
		return "float64"
	case "array":
		return "[]string"
	case "jsonb", "json":
		return "string"
	case "timestamptz", "datetime":
		return "*time.Time"
	default:
		return jsonType
	}
}
