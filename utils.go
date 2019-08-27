package easy_parser

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

var (
	ErrorInvalidType = errors.New("invalid type or not support.")
	ConvertFunc      = map[string]Convert{
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
		"json":                            Str2Json,
	}
)

type Convert func(string) (interface{}, error)

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

func Str2Json(str string) (interface{}, error) {
	var inf interface{}
	err := json.Unmarshal([]byte(str), &inf)
	return inf, err
}

func String2Type(str, typename string) (interface{}, error) {
	fn, ok := ConvertFunc[typename]
	if ok {
		return fn(str)
	}
	return nil, ErrorInvalidType
}
