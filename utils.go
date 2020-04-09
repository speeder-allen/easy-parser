package easy_parser

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"unicode"
)

var (
	ErrorInvalidType = errors.New("invalid type or not support.")
	ConvertBaseFunc  = map[string]ConvertBase{
		reflect.Int.String():              Str2Int,
		reflect.Uint.String():             Str2Uint,
		reflect.Int8.String():             Str2Int8,
		reflect.Uint8.String():            Str2Uint8,
		reflect.Int16.String():            Str2Int16,
		reflect.Uint16.String():           Str2Uint16,
		reflect.Int32.String():            Str2Int32,
		reflect.Uint32.String():           Str2Uint32,
		reflect.Int64.String():            Str2Int64,
		reflect.Uint64.String():           Str2Uint64,
		reflect.Float32.String():          Str2Float32,
		reflect.Float64.String():          Str2Float64,
		reflect.Bool.String():             Str2Bool,
		reflect.TypeOf([]byte{}).String(): Str2Bytes,
	}
	ConvertAdvancedFunc = map[string]ConvertAdvanced{
		"json": Str2Json,
		"xml":  Str2Xml,
	}
)

type ConvertBase func(string) (interface{}, error)
type ConvertAdvanced func(string, *reflect.Value) error

func Str2Int(str string) (interface{}, error) {
	return strconv.Atoi(str)
}

func Str2Uint(str string) (interface{}, error) {
	m, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(m), nil
}

func Str2Int8(str string) (interface{}, error) {
	m, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(m), nil
}

func Str2Uint8(str string) (interface{}, error) {
	m, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(m), nil
}

func Str2Int16(str string) (interface{}, error) {
	m, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(m), nil
}

func Str2Uint16(str string) (interface{}, error) {
	m, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(m), nil
}

func Str2Int32(str string) (interface{}, error) {
	m, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(m), nil
}

func Str2Uint32(str string) (interface{}, error) {
	m, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(m), nil
}

func Str2Int64(str string) (interface{}, error) {
	return strconv.ParseInt(str, 10, 64)
}

func Str2Uint64(str string) (interface{}, error) {
	return strconv.ParseUint(str, 10, 64)
}

func Str2Float32(str string) (interface{}, error) {
	m, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(m), nil
}

func Str2Float64(str string) (interface{}, error) {
	return strconv.ParseFloat(str, 64)
}

func Str2Bytes(str string) (interface{}, error) {
	return []byte(str), nil
}

func Str2Bool(str string) (interface{}, error) {
	return strconv.ParseBool(str)
}

func Str2Json(str string, value *reflect.Value) error {
	return json.Unmarshal([]byte(str), value.Interface())
}

func Str2Xml(str string, value *reflect.Value) error {
	return xml.Unmarshal([]byte(str), value.Interface())
}

func StringParser(str, decode string, inf *reflect.Value) error {
	if decode == reflect.String.String() {
		inf.Elem().SetString(str)
		return nil
	} else {
		fn, ok := ConvertBaseFunc[decode]
		if ok {
			m, err := fn(str)
			if err != nil {
				return err
			}
			inf.Elem().Set(reflect.ValueOf(m))
			return nil
		} else {
			fn, ok := ConvertAdvancedFunc[decode]
			if ok {
				return fn(str, inf)
			}
		}
	}
	return ErrorInvalidType
}

func isPointer(tp reflect.Type) bool {
	return tp.String()[0:1] == "*"
}

func isExportColumn(field reflect.StructField) bool {
	return unicode.IsUpper(rune(field.Name[0]))
}
